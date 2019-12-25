package main

import (
	"encoding/json"
	"log"
	"net"
	"os"
	"time"

	"github.com/miekg/dns"
	zlog "github.com/rs/zerolog"
)

// Starts a new UDP server that listens on port "x"
/*
	TODO tracker
	- spin up UDP server [x]
	-
*/

func main() {
	logger := newLogger()

	conn, err := net.ListenPacket("udp", ":1053")
	logger.Info().Msg("listening on :1053...")
	if err != nil {
		log.Fatal(err)
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

func newLogger() zlog.Logger {
	output := zlog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	logger := zlog.New(output).With().Timestamp().Logger()
	logger.Level(4)

	return logger
}
