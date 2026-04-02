package data

import "net"

// RPC  represent the message from the peer send over each transport between each nodes of the network
type RPC struct {
	Payload []byte

	From net.Addr
}

