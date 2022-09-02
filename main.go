package main

import (
	"log"
	"net"

	"github.com/txthinking/socks5"
)

func main() {
	s, err := socks5.NewClassicServer("0.0.0.0:1080", "0.0.0.0", "", "", 30, 30)
	if err != nil {
		panic(err)
	}

	h := Handler{
		&socks5.DefaultHandle{},
	}

	err = s.ListenAndServe(h)
	if err != nil {
		panic(err)
	}
}

type Handler struct {
	H socks5.Handler
}

func (h Handler) TCPHandle(s *socks5.Server, source *net.TCPConn, r *socks5.Request) error {
	log.Printf("TCP: L[%s], R[%s], ReqA[%s]\n", source.LocalAddr().String(), source.RemoteAddr().String(), r.Address())
	return h.H.TCPHandle(s, source, r)
}

func (h Handler) UDPHandle(s *socks5.Server, source *net.UDPAddr, d *socks5.Datagram) error {
	log.Printf("UDP: L[%s], DGA[%s]\n", source.String(), d.Address())
	return h.H.UDPHandle(s, source, d)
}
