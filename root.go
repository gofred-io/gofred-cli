package main

import (
	"github.com/gofred-io/gofred-cli/app"
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
	app.Init(rootCmd)
}
