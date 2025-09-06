package run

import (
	"github.com/spf13/cobra"
)

var (
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run the application",
		Run:   runApp,
	}
)

func Init(parentCmd *cobra.Command) {
	parentCmd.AddCommand(runCmd)
}

func runApp(cmd *cobra.Command, args []string) {
	runServer()
}
