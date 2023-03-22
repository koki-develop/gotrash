package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// flags
var (
	flagListCurrentDir bool
)

var rootCmd = &cobra.Command{
	Use:   "gotrash",
	Short: "rm alternative written in Go.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// commands
	rootCmd.AddCommand(
		putCmd,
		listCmd,
		restoreCmd,
		clearCmd,
	)

	// flags
	listCmd.Flags().BoolVarP(&flagListCurrentDir, "current-dir", "c", false, "show only the trash in the current directory")
}
