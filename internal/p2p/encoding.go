package p2p

import (
	"encoding/gob"
	"fmt"

	// "fmt"
	"io"

	"github.com/satya-18-w/DataNest/internal/data"
	// "github.com/satya-18-w/DataNest/internal/data"
)

type Message struct {
}

type Decoder interface {
	Decode(io.Reader, any) error
}
type GOBDecoder struct {
}

func (d GOBDecoder) Decode(r io.Reader, v any) error {
	return gob.NewDecoder(r).Decode(v)
}

type DefaultDecoder struct{}

func (n DefaultDecoder) Decode(r io.Reader, v any) error {
	// msg, ok := v.(*data.Message)
	// if !ok {
	// 	return fmt.Errorf("invalid message type")
	// }
	msg := v.(*data.RPC)
	buf := make([]byte, 1024)
	m, err := r.Read(buf)
	if err != nil {
		return err
	}

	msg.Payload = buf[:m]
	fmt.Println(string(msg.Payload))

	return nil

}
