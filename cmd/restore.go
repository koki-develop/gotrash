package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/koki-develop/go-fzf"
	"github.com/koki-develop/gotrash/internal/db"
	"github.com/koki-develop/gotrash/internal/trash"
	"github.com/spf13/cobra"
)

const (
	mainColor = "#00ADD8"
)

var restoreCmd = &cobra.Command{
	Use:          "restore [index]...",
	Short:        "Restore trashed files or directories",
	Long:         "Restore trashed files or directories.",
	Aliases:      []string{"rs"},
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := db.Open()
		if err != nil {
			return err
		}
		defer db.Close()

		if len(args) == 0 {
			ts, err := db.List(false)
			if err != nil {
				return err
			}

			f, err := fzf.New(
				fzf.WithNoLimit(true),
				fzf.WithStyles(
					fzf.WithStyleCursor(fzf.Style{ForegroundColor: mainColor}),
					fzf.WithStyleCursorLine(fzf.Style{Bold: true}),
					fzf.WithStyleMatches(fzf.Style{ForegroundColor: mainColor}),
					fzf.WithStyleSelectedPrefix(fzf.Style{ForegroundColor: mainColor}),
					fzf.WithStyleUnselectedPrefix(fzf.Style{Faint: true}),
				),
			)
			if err != nil {
				return err
			}
			idxs, err := f.Find(
				ts,
				func(i int) string { return ts[i].Path },
				fzf.WithItemPrefix(func(i int) string { return fmt.Sprintf("(%s) ", ts[i].TrashedAt.Format(time.DateTime)) }),
			)
			if err != nil {
				return err
			}

			var choices trash.TrashList
			for _, i := range idxs {
				choices = append(choices, ts[i])
			}

			if err := db.Restore(choices, flagRestoreForce); err != nil {
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

			if err := db.RestoreByIndexes(is, flagRestoreForce); err != nil {
				return err
			}
		}

		return nil
	},
}
