package server

import (
	"net"
	"time"
)

// Server is a UDP listener
type Server struct {
	Addr    net.UDPAddr
	Timeout time.Time
}

// Start implements the UDPServer interface
func (s *Server) Start() {}

// Stop implements the UDPServer interface
func (s *Server) Stop() {}

// UDPServer - please find a more fitting name.
type UDPServer interface {
	Start() // Starts the server
	Stop()  // Gracefully stops the server
}
