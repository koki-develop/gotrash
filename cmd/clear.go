package cmd

import (
	"github.com/koki-develop/gotrash/internal/db"
	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use: "clear",
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
