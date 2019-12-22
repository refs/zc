package main

import (
	"log"
	"net"
	"os"
	"time"

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
		go serve(conn, addr, buf[:n])
	}

}

func serve(pc net.PacketConn, addr net.Addr, buf []byte) {
	// 0 - 1: ID
	// 2: QR(1): Opcode(4)
	buf[2] |= 0x80 // Set QR bit

	pc.WriteTo(buf, addr)
}

func newLogger() zlog.Logger {
	output := zlog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	logger := zlog.New(output).With().Timestamp().Logger()
	logger.Level(4)

	return logger
}
