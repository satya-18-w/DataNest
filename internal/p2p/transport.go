package p2p

import "github.com/satya-18-w/DataNest/internal/data"

// Peer is an abstraction of a remote node in the network
type Peer interface {
	Close() error // So that we can close the connection of the perr

}

// Transport is anything that handle the communication between the Nodes in the NetWork
// Tcp , Udp and Websocket

type Transport interface {
	ListenAndAccept() error
	Consume() <-chan data.RPC
}


type Options struct{
	
}