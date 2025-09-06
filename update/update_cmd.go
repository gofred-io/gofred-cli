package update

import (
	"github.com/spf13/cobra"
)

var (
	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update gofred",
		Run: func(cmd *cobra.Command, args []string) {
			CheckForUpdates(true)
		},
	}
)

func Init(rootCmd *cobra.Command) {
	rootCmd.AddCommand(updateCmd)
}
