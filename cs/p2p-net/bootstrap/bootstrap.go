package main

import (
	"context"
	"flag"
	"io/ioutil"
	"fmt"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"

	dht "github.com/libp2p/go-libp2p-kad-dht"
	//rand "crypto/rand"
	multiaddr "github.com/multiformats/go-multiaddr"
	logging "github.com/whyrusleeping/go-logging"

	"github.com/ipfs/go-log"
)

var logger = log.Logger("rendezvous")

func main() {
	log.SetAllLoggers(logging.WARNING)
	log.SetLogLevel("rendezvous", "info")
	help := flag.Bool("h", false, "Display Help")
	config, err := ParseFlags()
	if err != nil {
		panic(err)
	}

	if *help {
		fmt.Println("This program demonstrates a simple p2p chat application using libp2p")
		fmt.Println()
		fmt.Println("Usage: Run './bootstrap ")
		flag.PrintDefaults()
		return
	}

	ctx := context.Background()

	byte_prv, err := ioutil.ReadFile("../../nodekey/rivatekey"+config.PeerNum+".txt")
	if err!=nil{
		panic(err)
	}
	prvKey, _ := crypto.UnmarshalPrivateKey(byte_prv)
	fmt.Println(prvKey)
	
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


	// Start a DHT, for use in peer discovery. We can't just make a new DHT
	// client because we want each peer to maintain its own local copy of the
	// DHT, so that the bootstrapping node of the DHT can go down without
	// inhibiting future peer discovery.
	_, err = dht.New(ctx, host)
	if err != nil {
		panic(err)
	}

	select {}
	
}
