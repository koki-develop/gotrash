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
	flagListCurrentDir bool

	// restore
	flagRestoreForce bool

	// clear
	flagClearForce bool
)

var rootCmd = &cobra.Command{
	Use:          "gotrash",
	Long:         "rm alternative written in Go.",
	SilenceUsage: true,
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
	listCmd.Flags().BoolVarP(&flagListCurrentDir, "current-dir", "c", false, "show only the trash in the current directory")

	// restore
	restoreCmd.Flags().BoolVarP(&flagRestoreForce, "force", "f", false, "overwrite a file or directory if it already exists")

	// clear
	clearCmd.Flags().BoolVarP(&flagClearForce, "force", "f", false, "skip confirmation before clear")
}
