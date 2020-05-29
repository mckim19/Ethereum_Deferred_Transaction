package statedb

import (
	"flag"
	"strings"
	"fmt"
	
	dht "github.com/libp2p/go-libp2p-kad-dht"
	maddr "github.com/multiformats/go-multiaddr"
)

var config Config
type addrList []maddr.Multiaddr

func (al *addrList) String() string {
	strs := make([]string, len(*al))
	for i, addr := range *al {
		strs[i] = addr.String()
	}
	return strings.Join(strs, ",")
}

func (al *addrList) Set(value string) error {
	addr, err := maddr.NewMultiaddr(value)
	if err != nil {
		return err
	}
	*al = append(*al, addr)
	return nil
}

func StringsToAddrs(addrStrings []string) (maddrs []maddr.Multiaddr, err error) {
	for _, addrString := range addrStrings {
		addr, err := maddr.NewMultiaddr(addrString)
		if err != nil {
			return maddrs, err
		}
		maddrs = append(maddrs, addr)
	}
	return
}
// A new type we need for writing a custom flag parser
type Config struct {
	Help			 bool
	RendezvousString string
	Role             string
	BootstrapPeers   addrList
	ListenAddresses  addrList
	ProtocolID       string
	PeerNum          string
	TotalPeerNum     int
}

func ParseFlags() (Config, bool) {
	flag.BoolVar(&config.Help, "h", false, "Display Help")
	flag.StringVar(&config.Role, "role", "mapper", "add role")
	flag.StringVar(&config.RendezvousString, "rendezvous", "yoom network",
		"Unique string to identify group of nodes. Share this with your friends to let them connect with you")
	flag.Var(&config.BootstrapPeers, "peer", "Adds a peer multiaddress to the bootstrap list")
	flag.Var(&config.ListenAddresses, "listen", "Adds a multiaddress to the listen list")
	flag.StringVar(&config.ProtocolID, "pid", "/yoomee1313/1.0.1", "Sets a protocol id for stream headers")
	flag.StringVar(&config.PeerNum, "peernum", "nothing", "add private key path")
	flag.IntVar(&config.TotalPeerNum, "totalpeernum", 0, "add total peer num of p2p network")
	flag.Parse()

	if config.Help {
		fmt.Println("This program demonstrates a simple p2p chat application using libp2p")
		fmt.Println()
		fmt.Println("Usage: Run './chat in two different terminals. Let them connect to the bootstrap nodes, announce themselves and connect to the peers")
		flag.PrintDefaults()
		return config, false
	}
	if len(config.BootstrapPeers) == 0 {
		config.BootstrapPeers = dht.DefaultBootstrapPeers
	}
	if config.PeerNum == "nothing" {
		fmt.Println("wrong number")
		return config, false
	}
	return config, true
}
