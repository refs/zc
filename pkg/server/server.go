package server

import "net"

// Server is a UDP listener
type Server struct {
	Addr net.UDPAddr
}
