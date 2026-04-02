package p2p

import "errors"

type HandshakerFunc func(p any) error

// Handshake is a process of verifying the identity of the remote node
var ErrInValidHadnshake = errors.New("invalid handshake")

type Handshaker interface {
	Handshake(p Peer) error
}

type DefaultHandshaker struct {
}

func (h *DefaultHandshaker) Handshake(p Peer) error {
	return nil
}

func NophandshakeFunc(p any) error {
	return nil
}
	