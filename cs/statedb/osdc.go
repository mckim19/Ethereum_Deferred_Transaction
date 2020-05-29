package statedb

import (
	"io/ioutil"
	"bufio"
	"fmt"
	"sync"
	"context"
	"time"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"

	"github.com/libp2p/go-libp2p-discovery"


	dht "github.com/libp2p/go-libp2p-kad-dht"
	multiaddr "github.com/multiformats/go-multiaddr"

	logging "github.com/whyrusleeping/go-logging"
	"github.com/ipfs/go-log"
)
// 1. Logging
var logger = log.Logger("rendezvous")
func SetLog(){
	log.SetAllLoggers(logging.WARNING)
	log.SetLogLevel("rendezvous", "info")
}

// 2. p2p networking related variables and functions
var Och chan []byte
var Ich [](chan []byte)

func handleStream(stream network.Stream) {
	logger.Info("Got a new stream!")

	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

	newIch := make(chan []byte)
	Ich = append(Ich, newIch)
	go writeData(rw, newIch)
	go readData(rw, newIch)

	// 'stream' will stay open until you close it (or the other side closes it).

	totalPeerNum += 1
	if totalPeerNum >= config.TotalPeerNum - 1 {
		done <- true
	}
}
func readData(rw *bufio.ReadWriter, ch chan []byte) {
	for {
		data := make([]byte, 32)
		_, err := rw.Read(data)
		if err != nil {
			fmt.Println("Error reading from buffer")
			fmt.Println(data)
			//panic(err)
			Och <- data
			return
		}
		fmt.Println(data)
		
		// Green console colour: 	\x1b[32m
		// Reset console colour: 	\x1b[0m
		fmt.Printf("\x1b[32m%s\x1b[0m \n", string(data))
		Och <- data
		
	}
}
func writeData(rw *bufio.ReadWriter, ch chan []byte) {
	for {
		sendData := <- ch
		fmt.Println("sendData:", sendData)
		_, err := rw.Write(sendData)
		if err != nil {
			fmt.Println("Error writing to buffer")
			//panic(err)
			return
		}
		err = rw.Flush()

		if err != nil {
			fmt.Println("Error flushing buffer")
			panic(err)
		}
	}
}

// 3. running node related variables
var totalPeerNum = 0
var done chan bool

// 4. Get Private Key and use it as a p2p node id
func GetPrvKey() (crypto.PrivKey, error){
	byte_prv, err := ioutil.ReadFile("../nodekey/rivatekey"+config.PeerNum+".txt")
	if(err!=nil){
		panic(err)
	}
	prvKey, err := crypto.UnmarshalPrivateKey(byte_prv)
	return prvKey, err
}

func RunNode(config Config) {
	Och = make(chan []byte, 10000)
	done = make(chan bool, 1)

	SetLog()

	ctx := context.Background()
	prvKey , _:= GetPrvKey()
	
	// libp2p.New constructs a new libp2p Host. Other options can be added
	// here.
	host, err := libp2p.New(ctx,
		libp2p.ListenAddrs([]multiaddr.Multiaddr(config.ListenAddresses)...),
		libp2p.Identity(prvKey),
	)
	if err != nil {
		panic(err)
	}
	logger.Info("Host created. We are:", host.ID())
	logger.Info(host.Addrs())

	// Set a function as stream handler. This function is called when a peer
	// initiates a connection and starts a stream with this peer.
	host.SetStreamHandler(protocol.ID(config.ProtocolID), handleStream)

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
	logger.Debug("Bootstrapping the DHT")
	if err = kademliaDHT.Bootstrap(ctx); err != nil {
		panic(err)
	}

	// Let's connect to the bootstrap nodes first. They will tell us about the
	// other nodes in the network.
	fmt.Println(config.BootstrapPeers[0])
	var wg sync.WaitGroup
	for _, peerAddr := range config.BootstrapPeers {
		peerinfo, _ := peer.AddrInfoFromP2pAddr(peerAddr)
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := host.Connect(ctx, *peerinfo); err != nil {
				logger.Warning(err)
			} else {
				logger.Info("Connection established with bootstrap node:", *peerinfo)
			}
		}()
	}
	wg.Wait()


	// We use a rendezvous point "meet me here" to announce our location.
	// This is like telling your friends to meet you at the Eiffel Tower.
	logger.Info("Announcing ourselves...")
	routingDiscovery := discovery.NewRoutingDiscovery(kademliaDHT)
	discovery.Advertise(ctx, routingDiscovery, config.RendezvousString)
	logger.Info("Successfully announced!")

	// Now, look for others who have announced
	// This is like your friend telling you the location to meet you.
	logger.Info("Searching for other peers...")
	peerChan, err := routingDiscovery.FindPeers(ctx, config.RendezvousString)
	if err != nil {
		panic(err)
	}

	for peer := range peerChan {
		if peer.ID == host.ID() {
			continue
		}

		logger.Info("Found peer.. Connecting to:", peer)
		stream, err := host.NewStream(ctx, peer.ID, protocol.ID(config.ProtocolID))

		if err != nil {
			logger.Warning("Connection failed:", err)
			continue
		} else {
			go handleStream(stream)
		}
		logger.Info("Connected to:", peer)
	}
	<-done
}
func SleepNode(config Config) {
	time.Sleep(time.Second)
}