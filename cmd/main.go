package main

import (
	"fmt"
	"log"

	"github.com/satya-18-w/DataNest/internal/tcp"
)

func main() {
	tr := tcp.NewTCPtransPort("localhost:8002")
	log.Fatal(tr.ListenAndAccept())
	fmt.Println("Hello This is the Styarting of teh DataNest")
	select {}
	// t := tcp.NewTCPtransPort("localhost:8002")

}
