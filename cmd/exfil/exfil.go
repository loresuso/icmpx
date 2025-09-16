package exfil

import (
	"github.com/loresuso/icmpx/cmd/exfil/recv"
	"github.com/loresuso/icmpx/cmd/exfil/send"
	"github.com/spf13/cobra"
)

func NewExfilCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exfil",
		Short: "Data exfiltration via ICMP",
		Long:  "Exfiltrate data using ICMP packets for covert data transfer",
	}

	// Add subcommands
	cmd.AddCommand(send.New())
	cmd.AddCommand(recv.New())

	return cmd
}
