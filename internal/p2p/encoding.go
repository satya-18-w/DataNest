package p2p

import "io"

type Message struct {
}

type Decoder interface {
	Decode(io.Reader, any) error
}
