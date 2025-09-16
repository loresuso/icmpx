package main

import "github.com/loresuso/icmpx/cmd"

func main() {
	rootCmd := cmd.NewRoot()
	rootCmd.Execute()
}
