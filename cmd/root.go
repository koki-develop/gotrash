package cmd

import (
	"os"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var (
	version string
)

// flags
var (
	// list
	flagListAll bool

	// clear
	flagClearForce bool
)

var rootCmd = &cobra.Command{
	Use:  "gotrash",
	Long: "rm alternative written in Go.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	/*
	 * version
	 */

	if version == "" {
		if info, ok := debug.ReadBuildInfo(); ok {
			version = info.Main.Version
		}
	}

	rootCmd.Version = version

	/*
	 * commands
	 */

	rootCmd.AddCommand(
		putCmd,
		listCmd,
		restoreCmd,
		clearCmd,
	)

	/*
	 * flags
	 */

	// list
	listCmd.Flags().BoolVarP(&flagListAll, "all", "a", false, "show all trash")

	// clear
	clearCmd.Flags().BoolVarP(&flagClearForce, "force", "f", false, "skip confirmation before clear")
}
