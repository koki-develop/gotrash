package cmd

import "github.com/spf13/cobra"

var clearCmd = &cobra.Command{
	Use: "clear",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
