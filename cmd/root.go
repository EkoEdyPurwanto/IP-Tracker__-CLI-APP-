package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "Trace",
	Short: "IP-Tracker",
	Long:  `IP-Tracker__-CLI-APP-`,
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
