package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/koki-develop/gotrash/internal/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:          "list",
	Short:        "List trashed flies or directories",
	Long:         "List trashed flies or directories.",
	Aliases:      []string{"ls"},
	Args:         cobra.MaximumNArgs(0),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := db.Open()
		if err != nil {
			return err
		}
		defer db.Close()

		ts, err := db.List(true)
		if err != nil {
			return err
		}

		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		digits := len(strconv.Itoa(len(ts)))
		f := fmt.Sprintf("%%%dd: (%%s), %%s\n", digits)
		for i, t := range ts {
			if flagListCurrentDir {
				if !strings.HasPrefix(t.Path, cwd) {
					continue
				}
			}

			fmt.Printf(f, i, t.TrashedAt.Format(time.DateTime), t.Path)
		}

		return nil
	},
}
