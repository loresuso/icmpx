package recv

import (
	"fmt"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "recv [output-file]",
		Short: "Receive file via ICMP exfiltration",
		Long:  "Listen for incoming ICMP packets and reconstruct the exfiltrated file",
		Args:  cobra.ExactArgs(1),
		Run:   run,
	}

	// Add flags here

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	outputFile := args[0]

	// TODO: Implement actual recv logic
	fmt.Printf("Listening for exfiltrated data, saving to '%s'\n", outputFile)
}
