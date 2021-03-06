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
// TODO: add some functional options for the funs of it, and a constructor.
type Server struct {
	Addr    net.UDPAddr
	Timeout time.Time
	Log     zerolog.Logger // TODO use this own package's logger
	stop    chan struct{}  // zero allocation channel
}

// Start starts the main UDP server flow
func (s *Server) Start() error {
	conn, err := net.ListenPacket("udp", s.Addr.String())
	if err != nil {
		os.Exit(1)
	}
	s.Log.Info().Msgf("listening on %v...", s.Addr.String())
	defer conn.Close()
	// conn.SetReadDeadline(s.Timeout)

	for {
		buf := make([]byte, 1024) // TODO make buffer size configurable, or unlimited. Perhaps there's a convention here
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			s.Log.Error().Err(err).Msg("closing server")
			break
		}
		s.Log.Info().Msgf("package from %v", addr.String())
		go s.Serve(conn, addr, buf[:n])
	}

	return nil
}

// Serve sends back a UDP packet to the origin address
func (s *Server) Serve(pc net.PacketConn, addr net.Addr, datagram []byte) {
	// Test response with an A RR
	aRr := dns.A{
		Hdr: dns.RR_Header{
			Name: "refs.com",
		},
		A: net.IPv4(0, 4, 2, 0),
	}

	data, err := json.Marshal(aRr)
	if err != nil {
		s.Log.Error().Err(err).Msg("marshaling A RR")
		pc.Close()
		os.Exit(1)
	}

	n, err := pc.WriteTo(data, addr)
	if err != nil {
		s.Log.Error().Err(err).Msg("error sending the package to origin")
	}

	s.Log.Info().Msgf("%v bytes written to %v", n, addr.String())
}

// Stop implements the UDPServer interface
func (s *Server) Stop() {
	// TODO write to `s.stop` channel
}
