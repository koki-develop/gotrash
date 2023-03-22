package cmd

import (
	"strconv"

	"github.com/koki-develop/gotrash/internal/db"
	"github.com/koki-develop/gotrash/internal/restoreui"
	"github.com/spf13/cobra"
)

var restoreCmd = &cobra.Command{
	Use:     "restore [index]...",
	Short:   "Restore trashed files or directories",
	Long:    "Restore trashed files or directories.",
	Aliases: []string{"rs"},
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := db.Open()
		if err != nil {
			return err
		}
		defer db.Close()

		if len(args) == 0 {
			ts, err := db.List()
			if err != nil {
				return err
			}

			m := restoreui.New(ts)
			if err := restoreui.Run(m); err != nil {
				return err
			}
		} else {
			var is []int
			for _, arg := range args {
				i, err := strconv.Atoi(arg)
				if err != nil {
					return err
				}
				is = append(is, i)
			}

			if err := db.Restore(is, flagRestoreForce); err != nil {
				return err
			}
		}

		return nil
	},
}
