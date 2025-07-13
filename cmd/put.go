package cmd

import (
	"github.com/koki-develop/gotrash/internal/db"
	"github.com/spf13/cobra"
)

var putCmd = &cobra.Command{
	Use:          "put [file]...",
	Short:        "Trash files or directories",
	Long:         "Trash files or directories.",
	Args:         cobra.MinimumNArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		database, err := db.Open()
		if err != nil {
			return err
		}
		defer func() { _ = database.Close() }()

		opts := db.PutOptions{
			RmMode:    flagPutRmMode,
			Recursive: flagPutRecursive,
			Force:     flagPutForce,
		}

		if err := database.Put(args, opts); err != nil {
			return err
		}

		return nil
	},
}
