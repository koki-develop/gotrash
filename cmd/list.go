package cmd

import (
	"fmt"

	"github.com/koki-develop/gotrash/internal/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Args:    cobra.MaximumNArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := db.Open()
		if err != nil {
			return err
		}
		defer db.Close()

		ts, err := db.List()
		if err != nil {
			return err
		}

		for i, t := range ts {
			fmt.Printf("%d: %s\n", i, t.Path)
		}

		return nil
	},
}
