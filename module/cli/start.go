package cli

import (
	"gpm/module/uds"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a new process",
	Run: func(cmd *cobra.Command, args []string) {
		uds.Start(args[0], args[1:]...)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
