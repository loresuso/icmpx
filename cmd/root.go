package cmd

import (
	"github.com/loresuso/icmpx/cmd/exfil"
	"github.com/loresuso/icmpx/cmd/shell"
	"github.com/spf13/cobra"
)

func NewRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "icmpx",
		Short: "icmpx is a tool for ICMP-based data exfiltration and reverse shells",
		Long:  "icmpx provides ICMP-based tools for data exfiltration and reverse shell connections",
	}

	// Add subcommands
	cmd.AddCommand(exfil.NewExfilCommand())
	cmd.AddCommand(shell.NewShellCommand())

	return cmd
}
