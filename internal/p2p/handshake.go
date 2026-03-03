package p2p



type HandshakerFunc func( p any) error



type Handshaker interface{
	Handshake(p Peer) error
}

type DefaultHandshaker struct{

}

func (h *DefaultHandshaker) Handshake(p Peer) error{
	return nil
}	