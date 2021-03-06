package main

import (
	"net"
	"os"
	"time"

	"github.com/refs/mdns/pkg/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":1054") // TODO: why does :1054 work but `localhost:1054` or `127.0.0.1:1054` don't?
	if err != nil {
		log.Err(err)
		os.Exit(1)
	}

	s := server.Server{
		Addr: *addr,
		Log:  NewLogger(1),
	}

	if err := s.Start(); err != nil {
		log.Err(err)
		os.Exit(1)
	}
}

// NewLogger is a utility function
func NewLogger(level zerolog.Level) zerolog.Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	return zerolog.New(output).With().Timestamp().Logger().Level(level)
}
