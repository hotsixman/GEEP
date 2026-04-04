package cli

import (
	"fmt"
	"gpm/module/daemon"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Start the GPM daemon process",
	Run: func(cmd *cobra.Command, args []string) {
		if os.Getenv("GPM_DAEMON_PROCESS") != "1" {
			fmt.Println("Starting GPM daemon in background...")
		}
		daemon.Daemonize()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
