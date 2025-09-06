package main

import (
	"github.com/gofred-io/gofred-cli/app"
	"github.com/gofred-io/gofred-cli/flags"
	"github.com/gofred-io/gofred-cli/update"
	"github.com/gofred-io/gofred-cli/version"
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

	app.Init(rootCmd)
	update.Init(rootCmd)
	version.Init(rootCmd)
}
