package version

import (
	"fmt"

	"github.com/gofred-io/gofred-cli/flags"
	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Show the version of gofred",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("gofred version: %s\n", flags.Version)
		},
	}
)

func Init(rootCmd *cobra.Command) {
	rootCmd.AddCommand(versionCmd)
}
