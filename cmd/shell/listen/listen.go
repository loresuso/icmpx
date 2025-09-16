package listen

import (
	"fmt"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "listen",
		Short: "Listen for incoming shell connections",
		Long:  "Listen for incoming reverse shell connections via ICMP",
		Args:  cobra.NoArgs,
		Run:   runListen,
	}

	// Add flags
	cmd.Flags().StringP("interface", "i", "", "Network interface to listen on")
	cmd.Flags().IntP("port", "p", 0, "Port to bind to (for additional protocols)")
	cmd.Flags().BoolP("verbose", "v", false, "Verbose output")
	cmd.Flags().StringP("filter", "f", "", "Additional packet filter")
	cmd.Flags().IntP("timeout", "t", 0, "Timeout in seconds (0 for no timeout)")

	return cmd
}

func runListen(cmd *cobra.Command, args []string) {
	// Get flags
	iface, _ := cmd.Flags().GetString("interface")
	port, _ := cmd.Flags().GetInt("port")
	verbose, _ := cmd.Flags().GetBool("verbose")
	filter, _ := cmd.Flags().GetString("filter")
	timeout, _ := cmd.Flags().GetInt("timeout")

	if verbose {
		fmt.Printf("Listening for incoming shell connections\n")
		fmt.Printf("Interface: %s, Port: %d, Filter: %s, Timeout: %ds\n", iface, port, filter, timeout)
	}

	// TODO: Implement actual listen logic
	fmt.Println("Listening for incoming reverse shell connections via ICMP...")
}
