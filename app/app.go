package app

import (
	"github.com/gofred-io/gofred-cli/app/create"
	"github.com/gofred-io/gofred-cli/app/run"
	"github.com/spf13/cobra"
)

var (
	appCmd = &cobra.Command{
		Use:   "app",
		Short: "Application related commands",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

func Init(rootCmd *cobra.Command) {
	create.Init(appCmd)
	run.Init(appCmd)
	rootCmd.AddCommand(appCmd)
}
