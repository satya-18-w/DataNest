package tcp

import (
	"fmt"
	"testing"

	"github.com/satya-18-w/DataNest/internal/p2p"
	"github.com/stretchr/testify/assert"
)

func TestTcptransport(t *testing.T) {
	fmt.Println("This is  Tcp trasport  ")
	adder := "localhost:8000"
	tr := NewTCPtransPort(TCPtransportops{
		ListenAddr:    adder,
		HandshakeFunc: p2p.NophandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	})
	assert.Equal(t, tr.ListenAddr, "localhost:8000")

	// Server
	assert.Nil(t, tr.ListenAndAccept())
	select {}

}
