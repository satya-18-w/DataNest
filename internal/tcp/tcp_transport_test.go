package tcp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTcptransport(t *testing.T) {
	fmt.Println("This is  Tcp trasport  ")
	adder := "localhost:8090"
	tr := NewTCPtransPort(adder)
	assert.Equal(t, tr.ListenAddress, adder)

	// Server

	assert.Nil(t, tr.ListenAndAccept())
	select {}

}
