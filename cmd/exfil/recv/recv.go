package recv

import (
	"log"
	"os"

	"github.com/loresuso/icmpx/pkg/exfil"
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

	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	defer file.Close()

	err = exfil.NewICMP().Receive(file)
	if err != nil {
		log.Fatalf("Error receiving data: %v", err)
	}
}
