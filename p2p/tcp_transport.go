package p2p

import (
	"net"
	"sync"
)

type TCPTransport struct {
	listenAddress string       //Address of a lisstening node in go
	listner       net.Listener //The net.Listener to accept incoming connections

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTcpTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{listenAddress: listenAddr}
}
