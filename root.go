package main

import (
	"github.com/gofred-io/gofred-cli/app"
	"github.com/gofred-io/gofred-cli/flags"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gofred",
		Short: "Gofred CLI",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

func init() {
	rootCmd.Flags().BoolVarP(flags.OfflineRef(), "offline", "o", false, "Run the application in offline mode")
	checkForUpdates()
	app.Init(rootCmd)
}
