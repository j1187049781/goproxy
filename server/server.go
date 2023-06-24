package server

import (
	"log"
	"net"
	"goproxy/util"
	"golang.org/x/net/proxy"
)

type Server struct {
	ListenAddr string
	DstAddr    string
	ProxyDialer *proxy.Dialer
}

func NewServer(listenAddr string, dstAddr string, proxyDialer *proxy.Dialer) *Server {
	return &Server{
		ListenAddr: listenAddr,
		DstAddr:    dstAddr,
		ProxyDialer: proxyDialer,
	}
}

func (s *Server) Run() {
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	log.Printf("listen on %s, --> %s\n", ln.Addr(), s.DstAddr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConnection(conn, s.ProxyDialer, s.DstAddr)
	}
}

func handleConnection(conn net.Conn, proxyDialer *proxy.Dialer, targetAddr string) {
	defer conn.Close()

	remoteConn, err := (*proxyDialer).Dial("tcp", targetAddr)
	if err != nil {
		log.Println(err)
		return
	}
	defer remoteConn.Close()

	util.Relay(remoteConn, conn)
	log.Printf("close connection %s <-> %s\n", conn.LocalAddr(), conn.RemoteAddr())

}