package shell

import (
	"github.com/loresuso/icmpx/cmd/shell/listen"
	"github.com/loresuso/icmpx/cmd/shell/start"
	"github.com/spf13/cobra"
)

func NewShellCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "shell",
		Short: "ICMP-based reverse shell",
		Long:  "Establish reverse shell connections using ICMP packets for covert communication",
	}

	// Add subcommands
	cmd.AddCommand(start.New())
	cmd.AddCommand(listen.New())

	return cmd
}
