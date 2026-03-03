package p2p

// Peer is an abstraction of a remote node in the network
type Peer interface {
}

// Transport is anything that handle the communication between the Nodes in the NetWork
// Tcp , Udp and Websocket

type Transport interface {
	ListenAndAccept() error
}
