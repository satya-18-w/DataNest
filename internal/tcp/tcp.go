package tcp

import (
	"fmt"
	"net"
	"reflect"
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/satya-18-w/DataNest/internal/data"
	"github.com/satya-18-w/DataNest/internal/p2p"
)

// TCP Peer Represents the remote node over a TCP Connection Established

type TCPPeer struct {
	Conn     net.Conn // Connection of the Peer
	outbound bool     // If the User is dialing the it is a outbound call if the user is recieving the message it will be inbound
}

func (tp *TCPPeer) Close() error {
	return tp.Conn.Close()
}

type TCPtransportops struct {
	ListenAddr    string
	HandshakeFunc p2p.HandshakerFunc
	Decoder       p2p.Decoder
	OnPeer        func(p2p.Peer) error
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		Conn:     conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	TCPtransportops
	Listener net.Listener
	mu       sync.RWMutex
	Peers    map[net.Addr]p2p.Peer
	RpcChan  chan data.RPC
}

func NewTCPtransPort(ops TCPtransportops) *TCPTransport {
	return &TCPTransport{
		TCPtransportops: ops,
		RpcChan:         make(chan data.RPC),
		Peers:           make(map[net.Addr]p2p.Peer),
	}

}

func (t *TCPTransport) ListenAndAccept() error {
	lr, err := net.Listen("tcp", t.ListenAddr)
	if err != nil {
		fmt.Println(" Error in ListenAndAccept ", err)
		return err

	}
	t.Listener = lr
	fmt.Printf(" Listening on port %s \n", t.ListenAddr)
	chan1 := make(chan bool)
	go t.StartAcceptLoop()
	<-chan1

	return nil

}

// Consume implment the Trnasport interface which will return read only Channel as we can see in the p2p package
// For reading the incoming messages recieved from another peer in the network

func (t *TCPTransport) Consume() <-chan data.RPC {
	return t.RpcChan
}

func (t *TCPTransport) StartAcceptLoop() {
	for {
		conn, err := t.Listener.Accept()
		if err != nil {
			fmt.Println(" Error in StartAcceptLoop ", err)
		}

		go t.handleConn(conn)

	}

}



type Temp struct{}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)
	defer peer.Close()

	fmt.Printf(" New Incoming Connection : %+v\n", peer.Conn)
	if err := t.HandshakeFunc(peer); err != nil {
		conn.Close()
		log.Err(err).Msg("Tcp Handshake Error ")

		return
	}
	if t.OnPeer != nil {
		err := t.OnPeer(peer)
		if err != nil {
			fmt.Println("Error in OnPeer: %s", err)
			return
		}

	}
	//  Read Loop
	msg := data.RPC{}
	// buf := make([]byte, 1024)
	for {
		// n,err :=conn.Read(buf)
		err := t.Decoder.Decode(conn, &msg)
		fmt.Println(reflect.TypeOf(err))
		// panic(err)
		// if errors.Is(err, net.ErrClosed) {
		// 	return

		// }
		if err != nil {
			fmt.Printf("TCP Error : %s\n", err)
			return

		}
		msg.From = conn.RemoteAddr()
		fmt.Printf("recieved Message:  %v", msg)
		t.RpcChan <- msg

	}

}
