package multicast

// Multicast provides a service that listens on a multicast group
// One should be able to start `n` servers / listeners, send a message to the UDP Multicast address and have all servers responding.
import (
	"encoding/json"
	"net"
	"os"
	"time"

	"github.com/miekg/dns"
	"github.com/rs/zerolog"
)

// resources:
// - https://askubuntu.com/questions/247625/what-is-the-loopback-device-and-how-do-i-use-it

func main() {
	log := NewLogger(1)
	// get a network interface and print it
	// in preparation for listening multicast:
	// func ListenMulticastUDP(network string, ifi *Interface, gaddr *UDPAddr) (*UDPConn, error)
	iface, err := net.InterfaceByIndex(1)
	if err != nil {
		log.Error().Err(err).Msg("oops")
		os.Exit(1)
	}

	gaddr, err := net.ResolveUDPAddr("udp", "224.0.0.1:5353")
	if err != nil {
		log.Error().Err(err).Msg("oops")
		os.Exit(1)
	}

	// start listening multicast...
	conn, err := net.ListenMulticastUDP("udp", iface, gaddr)
	if err != nil {
		log.Error().Err(err).Msg("oops")
		os.Exit(1)
	}

	defer conn.Close()

	for {
		buf := make([]byte, 1024) // TODO make buffer size configurable, or unlimited. Perhaps there's a convention here
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			break
		}
		go serve(conn, addr, buf[:n])
	}
}

func serve(pc net.PacketConn, addr net.Addr, datagram []byte) {
	log := NewLogger(1)
	// Test response with an A RR
	aRr := dns.A{
		Hdr: dns.RR_Header{
			Name: "refs.com",
		},
		A: net.IPv4(0, 4, 2, 0),
	}

	data, err := json.Marshal(aRr)
	if err != nil {
		log.Error().Err(err).Msg("marshaling A RR")
		pc.Close()
		os.Exit(1)
	}

	n, err := pc.WriteTo(data, addr)
	if err != nil {
		log.Error().Err(err).Msg("error sending the package to origin")
	}

	log.Info().Msgf("%v bytes written to %v", n, addr.String())
}

// NewLogger is a utility function
func NewLogger(level zerolog.Level) zerolog.Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	return zerolog.New(output).With().Timestamp().Logger().Level(level)
}
