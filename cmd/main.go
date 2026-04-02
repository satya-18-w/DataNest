package main

import (
	"fmt"
	"log"

	"github.com/satya-18-w/DataNest/internal/p2p"
	"github.com/satya-18-w/DataNest/internal/tcp"
)

func OnPeer(peer p2p.Peer) error {
	fmt.Println("Doing something With the Peer Some Kind of action")
	// peer.Close()
	return nil
}
func main() {
	tcpopts := tcp.TCPtransportops{
		ListenAddr:    "localhost:8002",
		HandshakeFunc: p2p.NophandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		OnPeer:        OnPeer,
	}
	tr := tcp.NewTCPtransPort(tcpopts)

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("Recieved Message: %+v\n", msg)
		}
	}()
	log.Fatal(tr.ListenAndAccept())
	fmt.Println("Hello This is the Styarting of teh DataNest")
	select {}
	// t := tcp.NewTCPtransPort("localhost:8002")

}


