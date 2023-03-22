package cmd

import (
	"strconv"

	"github.com/koki-develop/gotrash/internal/db"
	"github.com/spf13/cobra"
)

var restoreCmd = &cobra.Command{
	Use:     "restore [index]...",
	Short:   "restore trashed files or directories",
	Aliases: []string{"rs"},
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := db.Open()
		if err != nil {
			return err
		}
		defer db.Close()

		var is []int
		for _, arg := range args {
			i, err := strconv.Atoi(arg)
			if err != nil {
				return err
			}
			is = append(is, i)
		}

		if err := db.Restore(is); err != nil {
			return err
		}

		return nil
	},
}
