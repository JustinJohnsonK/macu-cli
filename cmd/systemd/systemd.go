package systemd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewFile() *cobra.Command {
	return &cobra.Command{
		Use:   "newF",
		Short: "Create new file and move inside",
		Long:  "CLI commands to create new file and move inside that file",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "Created %s and moved inside", args[0])
		},
	}
}
