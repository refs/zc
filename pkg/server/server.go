package server

import (
	"encoding/json"
	"net"
	"os"
	"time"

	"github.com/miekg/dns"
	"github.com/rs/zerolog"
)

// UDPServer - please find a more fitting name.
type UDPServer interface {
	Start()                                 // Starts the server
	Stop()                                  // Gracefully stops the server
	Serve(net.PacketConn, net.Addr, []byte) // Sends back a UDP packet to the origin address
}

// Server is a UDP listener
type Server struct {
	Addr    net.UDPAddr
	Timeout time.Time
	Log     zerolog.Logger // TODO use this own package's logger
	end     chan struct{}  // zero allocation channel
}

// Start starts the main UDP server flow
func (s *Server) Start() {
	conn, err := net.ListenPacket("udp", ":1053")
	s.Log.Info().Msg("listening on :1053...")
	if err != nil {
		os.Exit(1)
	}
	defer conn.Close()
	// conn.SetReadDeadline(time.Now().Add(30 * time.Second)) // uncomment for a 30 seconds timeout

	for {
		buf := make([]byte, 1024) // TODO make buffer size configurable, or unlimited. Perhaps there's a convention here
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			s.Log.Error().Err(err).Msg("closing server")
			break
		}
		s.Log.Info().Msg("connection received")
		go s.Serve(conn, addr, buf[:n])
	}
}

// Serve sends back a UDP packet to the origin address
func (s *Server) Serve(pc net.PacketConn, addr net.Addr, datagram []byte) {
	// Test response with an A RR
	aRr := dns.A{
		Hdr: dns.RR_Header{
			Name: "refs.com",
		},
		A: net.IPv4(0, 0, 0, 0),
	}

	data, err := json.Marshal(aRr)
	if err != nil {
		s.Log.Error().Err(err).Msg("marshaling A RR")
		pc.Close()
		os.Exit(1)
	}

	pc.WriteTo(data, addr)
}

// Stop implements the UDPServer interface
func (s *Server) Stop() {}
