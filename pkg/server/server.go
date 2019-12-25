package server

import (
	"net"
	"os"
	"time"

	"github.com/rs/zerolog"
)

// UDPServer - please find a more fitting name.
type UDPServer interface {
	Start() // Starts the server
	Serve() // Serves a UDP request
	Stop()  // Gracefully stops the server
}

// Server is a UDP listener
type Server struct {
	Addr    net.UDPAddr
	Timeout time.Time
	Log     zerolog.Logger // TODO use this own package's logger
}

// Serve implements the UDPServer interface
func (s *Server) Serve() {
	conn, err := net.ListenPacket("udp", ":1053")
	s.Log.Info().Msg("listening on :1053...")
	if err != nil {
		os.Exit(1)
	}
	defer conn.Close()
	// conn.SetReadDeadline(time.Now().Add(30 * time.Second)) // uncomment for a 30s deadline...

	for {
		buf := make([]byte, 1024)
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			s.Log.Error().Err(err).Msg("closing server")
			break
		}
		s.Log.Info().Msg("connection received")
	}
}

// Stop implements the UDPServer interface
func (s *Server) Stop() {}
