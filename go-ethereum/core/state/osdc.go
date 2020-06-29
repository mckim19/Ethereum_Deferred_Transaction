// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package state

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"time"	
	"encoding/binary"

	"bufio"
	"io/ioutil"
	"strconv"
	"sync"
	"context"
	"os"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	maddr "github.com/multiformats/go-multiaddr"
)
/*
	OSDC deferred transaction parallel execution project.
*/
// 1. Data structure and Global Variables
const BUFSIZE int = 32
type Msg struct {
	Type			string
	Data			[]byte
	ContractAddress	common.Address
}
type Config struct {
	RendezvousString string
	BootstrapPeers   addrList
	ListenAddresses  addrList
	ProtocolID       string
	PeerNum          string
	TotalPeerNum     int
}

type addrList []maddr.Multiaddr
var Och chan []byte
var Ich [](chan []byte)

var totalPeerNum = 0
var done chan bool
var Cfg Config

// 2. Epc opcode handler functions
/*
	EpcInit: InitNode(port int, peerNum int, totalPeerNum int)
	explanation: EpcInit handler function, called by opEpcInit
*/
func (self *StateDB) InitNode(port int, peerNum int, totalPeerNum int) {
	m, _ := maddr.NewMultiaddr("/ip4/166.104.144.103/tcp/4001/p2p/QmSxL8RFC5jWy81nNY32PLNANDTJisjAmT9KJUN1qntLvb")
	l, _ := maddr.NewMultiaddr("/ip4/127.0.0.1/tcp/"+strconv.Itoa(port))
	Cfg.BootstrapPeers = append(Cfg.BootstrapPeers, m)
	Cfg.ListenAddresses = append(Cfg.ListenAddresses, l)

	Cfg.RendezvousString = "osdc p2p network"
	Cfg.ProtocolID = "/osdc1313/1.0.1"
	Cfg.PeerNum = strconv.Itoa(peerNum)
	Cfg.TotalPeerNum = totalPeerNum

	Och = make(chan []byte, 10000)
	done = make(chan bool, 1)

	ctx := context.Background()
	prvKey , _:= GetPrvKey()
	
	// libp2p.New constructs a new libp2p Host. Other options can be added
	// here.
	host, err := libp2p.New(ctx,
		libp2p.ListenAddrs([]maddr.Multiaddr(Cfg.ListenAddresses)...),
		libp2p.Identity(prvKey),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("[rendevous] Host created. We are:", host.ID())
	fmt.Println("[rendevous]", host.Addrs())

	// Set a function as stream handler. This function is called when a peer
	// initiates a connection and starts a stream with this peer.
	host.SetStreamHandler(protocol.ID(Cfg.ProtocolID), handleStream)

	// Start a DHT, for use in peer discovery. We can't just make a new DHT
	// client because we want each peer to maintain its own local copy of the
	// DHT, so that the bootstrapping node of the DHT can go down without
	// inhibiting future peer discovery.
	kademliaDHT, err := dht.New(ctx, host)
	if err != nil {
		panic(err)
	}

	// Bootstrap the DHT. In the default configuration, this spawns a Background
	// thread that will refresh the peer table every five minutes.
	fmt.Println("[rendevous] Bootstrapping the DHT")
	if err = kademliaDHT.Bootstrap(ctx); err != nil {
		panic(err)
	}

	// Let's connect to the bootstrap nodes first. They will tell us about the
	// other nodes in the network.
	fmt.Println(Cfg.BootstrapPeers[0])
	var wg sync.WaitGroup
	for _, peerAddr := range Cfg.BootstrapPeers {
		peerinfo, _ := peer.AddrInfoFromP2pAddr(peerAddr)
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := host.Connect(ctx, *peerinfo); err != nil {
				fmt.Println("[rendevous]",err)
			} else {
				fmt.Println("[rendevous] Connection established with bootstrap node:", *peerinfo)
			}
		}()
	}
	wg.Wait()


	// We use a rendezvous point "meet me here" to announce our location.
	// This is like telling your friends to meet you at the Eiffel Tower.
	fmt.Println("[rendevous] Announcing ourselves...")
	routingDiscovery := discovery.NewRoutingDiscovery(kademliaDHT)
	discovery.Advertise(ctx, routingDiscovery, Cfg.RendezvousString)
	fmt.Println("[rendevous] Successfully announced!")

	// Now, look for others who have announced
	// This is like your friend telling you the location to meet you.
	fmt.Println("[rendevous] Searching for other peers...")
	peerChan, err := routingDiscovery.FindPeers(ctx, Cfg.RendezvousString)
	if err != nil {
		panic(err)
	}

	for peer := range peerChan {
		if peer.ID == host.ID() {
			continue
		}

		fmt.Println("[rendevous] Found peer.. Connecting to:", peer)
		stream, err := host.NewStream(ctx, peer.ID, protocol.ID(Cfg.ProtocolID))

		if err != nil {
			fmt.Println("[rendevous] Connection failed:", err)
			continue
		} else {
			go handleStream(stream)
		}
		fmt.Println("[rendevous] Connected to:", peer)
	}
	<-done
}
/*
	EpcExit: ExitNode()
	explanation: Exit handler function, called by opEpcExit
*/
func (self *StateDB) ExitNode(){
	time.Sleep(time.Second)
}
/*
	EpcSend: sendData(rw *bufio.ReadWriter, ch chan []byte)
	explanation: EpcSend handler function, called by opEpcSend
*/
func (self *StateDB) SendMsg(data []byte){
	for idx := range Ich {
		Ich[idx] <- data
	}
}
/*
	EpcRecv: recvData(rw *bufio.ReadWriter, ch chan []byte)
	explanation: EpcSend handler function, called by opEpcRecv
	handleStream
*/
func (self *StateDB) RecvMsg() (res []byte){
	var tmp []byte

	size := int64(binary.BigEndian.Uint64((<-Och)[24:]))
	for i:=0; i<int(size)/BUFSIZE; i++ {
		res = append(res, (<-Och)...)
	}
	if int(size)%BUFSIZE!=0 {
		tmp = <- Och
		for i:=0; i<BUFSIZE; i++ {
			if tmp[i] == 0 {
				tmp = tmp[0:i]
				break
			}
		}
		res = append(res, tmp...)
	}
	return
}
// 3. internal functions
func sendData(rw *bufio.ReadWriter, ch chan []byte) {
	for {
		data := make([]byte, BUFSIZE)
		_, err := rw.Read(data)
		if err != nil {
			fmt.Println("Error reading from buffer")
			Och <- data
			return
		}
		fmt.Println("readData:", data)
		// Green console colour: 	\x1b[32m
		// Reset console colour: 	\x1b[0m
		fmt.Printf("\x1b[32m%s\x1b[0m \n", string(data))
		Och <- data
	}
}
func recvData(rw *bufio.ReadWriter, ch chan []byte) {
	for {
		data := <- ch
		fmt.Println("sendData:", data)
		_, err := rw.Write(data)
		if err != nil {
			fmt.Println("Error writing to buffer")
			return
		}
		err = rw.Flush()
		if err != nil {
			fmt.Println("Error flushing buffer")
			panic(err)
		}
	}
}
func handleStream(stream network.Stream) {
	fmt.Println("[rendevous] Got a new stream!")

	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

	newIch := make(chan []byte)
	Ich = append(Ich, newIch)
	go sendData(rw, newIch)
	go recvData(rw, newIch)

	// 'stream' will stay open until you close it (or the other side closes it).

	totalPeerNum += 1
	if totalPeerNum >= Cfg.TotalPeerNum - 1 {
		done <- true
	}
}
func GetPrvKey() (crypto.PrivKey, error){ 
	byte_prv, err := ioutil.ReadFile(os.Getenv("HOME")+"/Ethereum_Deferred_Transaction/samples/nodekey/rivatekey"+Cfg.PeerNum+".txt")
	if(err!=nil){
		panic(err)
	}
	prvKey, err := crypto.UnmarshalPrivateKey(byte_prv)
	return prvKey, err
}
/*
func (statedb *StateDB) StartMutexThread(type_t int) {
		statedb.SetChannel(make(chan ChanMessage, 10), true)
		go statedb.MutexThread(statedb.GetChannel(true), true)
}
func (statedb *StateDB) TerminateMutexThread (type_t int){
	var nil_address common.Address
	msg:=ChanMessage{
		ContractAddress: nil_address,
		Data: nil,
		Type:"TERMINATION",
		Channel: nil,
	}
	statedb.GetChannel(true)<- msg
	return
}
func (self *StateDB) MutexThread(com_channel chan ChanMessage, isDoCall bool){
    for{
		msg := <- com_channel
		switch msg.Type {
		case "READ":
			msg.Data = <- Och
			msg.Channel<-msg
			break
		case "WRITE":
			for idx := range Ich {
				Ich[idx] <- msg.Data
			}
			msg.Channel<-msg
			break
		case "TERMINATE":
			time.Sleep(time.Second)
			return
		}
		
		if(msg.LockType =="LOCK") {
			recording_info[key] = append(recording_info[key], msg.TxHash)
			msg.LockType="OK"
			msg.IsLockBusy = true
			if(CurrentLockArray[key].IsLockBusy == false){	//nobody holds this lock
				CurrentLockArray[key] = msg
				msg.Channel <- msg
			} else {								//somebody holds this lock	
				LockRequestArray[key] = append(LockRequestArray[key], msg)
			}
		}else if (msg.LockType=="UNLOCK"){
			// Networking
			if(CurrentLockArray[key].TxHash == msg.TxHash){ //it must be a same transaction who have been locked
				msg.LockType="OK"
				msg.Channel <- msg
				if(len(LockRequestArray[key]) != 0){ //somebody is waiting
					LockRequestArray[key][0].Channel <- msg
					CurrentLockArray[key] = LockRequestArray[key][0]	//change current lock tx
					LockRequestArray[key] = LockRequestArray[key][1:]	//get rid of the current lock tx from lock request array
				} else {
					msg.IsLockBusy = false
					CurrentLockArray[key] = msg
				}
			} else { //somebody tries to unlock fakely
				msg.LockType="NOT_OK"
				msg.Channel <- msg
			}
		}else if(msg.LockType=="TERMINATION") {
			if(isDoCall!=true) { //when mining
				resChannel<-recording_info
			}
			return
		}	
		
    }
}
func (self *StateDB)GetChannel(isDoCall bool)(chan ChanMessage){
	if isDoCall == true {
		return self.ch_com2
	}
	return self.ch_com
}
func (self *StateDB)GetTHash()(thash common.Hash){
	return self.thash
}
func (self *StateDB)SetChannel(new_ch chan ChanMessage, isDoCall bool){
	if isDoCall == true {
		self.ch_com2 = new_ch
	}
	self.ch_com = new_ch
}*/