package main
import (
	//"fmt"
	"github.com/libp2p/go-libp2p-core/crypto"
	"io/ioutil"
	"os"
	rand "crypto/rand"
)
// this is just for test..
// it would be replaced by getting address from statedb in ethereum
func main() {
	peer_num := os.Args[1]
	prvKey, pubKey, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, rand.Reader)
	if err != nil {
		panic(err)
	}
	// var mars_prv []byte
	// var mars_pub []byte
	byte_prv, _ := crypto.MarshalPrivateKey(prvKey)
	byte_pub, _ := crypto.MarshalPublicKey(pubKey)
	// fmt.Println(mars_prv)
	// fmt.Println(mars_pub)
	f1, _ := os.Create("../node/rivatekey"+peer_num+".txt")
	f2, _ := os.Create("../node/pubkey"+peer_num+".txt")
	f1.Close()
	f2.Close()
	_ = ioutil.WriteFile("../node/rivatekey"+peer_num+".txt", byte_prv, os.FileMode(744))
	_ = ioutil.WriteFile("../node/pubkey"+peer_num+".txt", byte_pub, os.FileMode(744))
	// byte_prv, _ = ioutil.ReadFile("../node/rivatekey.txt")
	// byte_pub, _ = ioutil.ReadFile("../node/pubkey.txt")
	// fmt.Println(byte_prv)
	// fmt.Println(byte_pub)
	
}