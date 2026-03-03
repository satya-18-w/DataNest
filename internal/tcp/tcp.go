package tcp

import (
	"fmt"
	"net"
	"sync"

	"github.com/satya-18-w/DataNest/internal/p2p"
)

// TCP Peer Represents the remote node over a TCP Connection Established

type TCPPeer struct {
	Conn     net.Conn // Connection of the Peer
	outbound bool      // If the User is dialing the it is a outbound call if the user is recieving the message it will be inbound
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		Conn:     conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	ListenAddress string
	Listener      net.Listener
	handshaker  p2p.HandshakerFunc
	// decoder 
	mu            sync.RWMutex
	Peers         map[net.Addr]p2p.Peer
}

func NewTCPtransPort(listenaddr string) *TCPTransport {
	return &TCPTransport{
		handshaker: func(p any) error{
			return  nil
		},
		ListenAddress: listenaddr,
	}

}

func (t *TCPTransport) ListenAndAccept() error {
	lr, err := net.Listen("tcp", t.ListenAddress)
	if err != nil {
		fmt.Println(" Error in ListenAndAccept ", err)
		return err

	}
	t.Listener = lr
	fmt.Printf(" Listening on port %s \n", t.ListenAddress)
	chan1:=make(chan bool)
	go t.StartAcceptLoop()
    <- chan1

	return nil

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

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer:= NewTCPPeer(conn,true)
	fmt.Printf(" New Incoming Connection : %+v\n", peer.Conn)
	if err:=t.handshaker(true) ; err !=nil{
		fmt.Println(" Error in handleConn ",err)
		return


	}
	for {
		buffer :=make([]byte,1024)
		_,err:= peer.Conn.Read(buffer)
		if err != nil{
			fmt.Println(" Error in handleConn ",err)
			break
			
		}
		fmt.Println(" Recieved Message : ",string(buffer)	)
	}
	

}
