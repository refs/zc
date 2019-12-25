package server

import (
	"net"
	"time"
)

// UDPServer - please find a more fitting name.
type UDPServer interface {
	Serve() // Starts the server
	Stop()  // Gracefully stops the server
}

// Server is a UDP listener
type Server struct {
	Addr    net.UDPAddr
	Timeout time.Time
}

// Serve implements the UDPServer interface
func (s *Server) Serve() {}

// Stop implements the UDPServer interface
func (s *Server) Stop() {}
