package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPTransport struct {
	listenAddress string       //IP address nad the port on which the node is listening
	listner       net.Listener //The net.Listener to accept incoming connections

	mu    sync.RWMutex      //Is used for managing concurrent access to resources,allowing for safe reading and writing across multiple goroutines
	peers map[net.Addr]Peer // A map where each peer's network address (net.Addr) is mapped to a Peer object
}

func NewTcpTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{listenAddress: listenAddr}
}
func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listner, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}
	go t.startAcceptlistenLoop() //This method is responsible for accepting conncetion asynchronously

	return nil

}
func (t *TCPTransport) startAcceptlistenLoop() {
	for {
		conn, err := t.listner.Accept()
		if err != nil {
			fmt.Println("Tcp Accept error")
		}

		go t.handleConn(conn) // handels connection asynchronously
	}
}
func (t *TCPTransport) handleConn(conn net.Conn) {
	fmt.Printf(("%v\n"), conn)
}
