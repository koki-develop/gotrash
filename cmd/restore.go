package cmd

import "github.com/spf13/cobra"

var restoreCmd = &cobra.Command{
	Use:     "restore",
	Aliases: []string{"rs"},
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
