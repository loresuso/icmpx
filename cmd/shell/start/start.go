package start

import (
	"fmt"

	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start [target]",
		Short: "Start a reverse shell connection",
		Long:  "Start a reverse shell connection to the target host using ICMP",
		Args:  cobra.ExactArgs(1),
		Run:   runStart,
	}

	// Add flags
	cmd.Flags().StringP("interface", "i", "", "Network interface to use")
	cmd.Flags().IntP("interval", "n", 1000, "Ping interval in milliseconds")
	cmd.Flags().StringP("shell", "s", "/bin/sh", "Shell to execute")
	cmd.Flags().BoolP("verbose", "v", false, "Verbose output")
	cmd.Flags().IntP("timeout", "t", 30, "Connection timeout in seconds")

	return cmd
}

func runStart(cmd *cobra.Command, args []string) {
	target := args[0]

	// Get flags
	iface, _ := cmd.Flags().GetString("interface")
	interval, _ := cmd.Flags().GetInt("interval")
	shell, _ := cmd.Flags().GetString("shell")
	verbose, _ := cmd.Flags().GetBool("verbose")
	timeout, _ := cmd.Flags().GetInt("timeout")

	if verbose {
		fmt.Printf("Starting reverse shell to target: %s\n", target)
		fmt.Printf("Interface: %s, Interval: %dms, Shell: %s, Timeout: %ds\n", iface, interval, shell, timeout)
	}

	// TODO: Implement actual reverse shell logic
	fmt.Printf("Starting reverse shell connection to '%s'\n", target)
}
