package cmd

import (
	"fmt"

	"github.com/koki-develop/gotrash/internal/db"
	"github.com/koki-develop/gotrash/internal/util"
	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:          "clear",
	Short:        "Clear all trashed files or directories",
	Long:         "Clear all trashed files or directories.",
	Args:         cobra.MaximumNArgs(0),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := db.Open()
		if err != nil {
			return err
		}
		defer func() { _ = db.Close() }()

		if !flagClearForce {
			if !util.YesNo("clear all trashed files or directories?") {
				fmt.Println("canceled.")
				return nil
			}
		}

		if err := db.ClearAll(); err != nil {
			return err
		}

		return nil
	},
}
