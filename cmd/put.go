package cmd

import "github.com/spf13/cobra"

var putCmd = &cobra.Command{
	Use: "put",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
