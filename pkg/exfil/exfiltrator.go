package exfil

import (
	"io"
)

type Exfiltrator interface {
	Exfiltrate(data io.Reader, to string) error
	Receive(output io.Writer) error
}
