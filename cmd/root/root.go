package root

import (
	"fmt"

	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "himacu",
		Short: "Macbook Cli for personal usecase",
		Long:  "CLI commands to make life easier using personal macbooks",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprint(cmd.OutOrStdout(), "My Macu is a good girl")
		},
	}
}
