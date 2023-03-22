package cmd

import (
	"github.com/koki-develop/gotrash/internal/db"
	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "clear all trashed files or directories",
	Args:  cobra.MaximumNArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := db.Open()
		if err != nil {
			return err
		}
		defer db.Close()

		if err := db.ClearAll(); err != nil {
			return err
		}

		return nil
	},
}
