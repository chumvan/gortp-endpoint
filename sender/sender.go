package sender

import (
	"net"
	"os"

	"github.com/chumvan/gortp-endpoint/endpoint"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

// Sender is the interface that wraps the Send method.
type Sender interface {
	Dial(string, string) error
	Send([]byte) error
	Close() error
}

// UDPConnection is a struct that implements the Sender interface.
type UDPConnection struct {
	conn *net.UDPConn
}

// NewUDPConnection returns a new UDPConnection.
func NewUDPConnection() *UDPConnection {
	return &UDPConnection{}
}

// Dial connects to the host and port specified.
func (c *UDPConnection) Dial(host string, port endpoint.Port) error {
	udpAddr, err := net.ResolveUDPAddr("udp", host+":"+port)
	log.Info("Resolving ", udpAddr)
	if err != nil {
		return err
	}
	c.conn, err = net.DialUDP("udp", nil, udpAddr)
	log.Info("Dialing ", c.conn)
	if err != nil {
		return err
	}
	return nil
}

// Send sends the data to the host and port specified.
func (c *UDPConnection) Send(data []byte) error {
	_, err := c.conn.Write(data)
	if err != nil {
		return err
	}
	return nil
}

// Close closes the connection.
func (c *UDPConnection) Close() error {
	err := c.conn.Close()
	log.Info("Closing ", c.conn)
	if err != nil {
		return err
	}
	return nil
}
