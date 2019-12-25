package server

import (
	"net"
	"time"

	"github.com/rs/zerolog"
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
	Log     zerolog.Logger // TODO use this own package's logger
}

// Serve implements the UDPServer interface
func (s *Server) Serve() {}

// Stop implements the UDPServer interface
func (s *Server) Stop() {}
