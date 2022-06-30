package cmd

import (
	"github.com/spf13/cobra"
)

type CLIOptions struct {
	Verbose bool
	Address string
	Port    int32
}

var (
	rootCmd = &cobra.Command{
		Use:   "timestamp-service",
		Short: "Timestamp Service",
	}
	CLI CLIOptions
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().BoolVarP(&CLI.Verbose, "verbose", "v", false, "Run timestamp service verbosely")
	rootCmd.Flags().StringVarP(&CLI.Address, "address", "a", "127.0.0.1", "Address timestamp service would listen")
	rootCmd.Flags().Int32VarP(&CLI.Port, "port", "p", 8080, "Port timestamp service would listen")
}
