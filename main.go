package main

import (
	"encoding/json"
	"net"
	"os"
	"time"

	"github.com/miekg/dns"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog"
)

func main() {
	logger := NewLogger(zlog.Level(1))
	conn, err := net.ListenPacket("udp", ":1053")
	logger.Info().Msg("listening on :1053...")
	if err != nil {
		os.Exit(1)
	}
	defer conn.Close()
	// conn.SetReadDeadline(time.Now().Add(30 * time.Second)) // uncomment for a 30s deadline...

	for {
		buf := make([]byte, 1024)
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			logger.Error().Err(err).Msg("closing server")
			break
		}
		logger.Info().Msg("connection received")
		go serve(conn, addr, buf[:n], logger)
	}

}

func serve(pc net.PacketConn, addr net.Addr, buf []byte, logger zlog.Logger) {
	// Test response with an A RR
	aRr := dns.A{
		Hdr: dns.RR_Header{
			Name: "refs.com",
		},
		A: net.IPv4(0, 0, 0, 0),
	}

	data, err := json.Marshal(aRr)
	if err != nil {
		logger.Error().Err(err).Msg("marshaling A RR")
		pc.Close()
		os.Exit(1)
	}

	pc.WriteTo(data, addr)
}

// NewLogger is a utility function
func NewLogger(level zerolog.Level) zerolog.Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	return zerolog.New(output).With().Timestamp().Logger().Level(level)
}
