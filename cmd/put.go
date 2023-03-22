package cmd

import (
	"github.com/koki-develop/gotrash/internal/db"
	"github.com/spf13/cobra"
)

var putCmd = &cobra.Command{
	Use:   "put [file]...",
	Short: "Trash files or directories",
	Long:  "Trash files or directories.",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := db.Open()
		if err != nil {
			return err
		}
		defer db.Close()

		if err := db.Put(args); err != nil {
			return err
		}

		return nil
	},
}
