package cli

import (
	"gpm/module/daemon"
	"gpm/module/logger"
	"os"

	"github.com/spf13/cobra"
)

var terminateCmd = &cobra.Command{
	Use:   "terminate",
	Short: "Terminate GPM daemon process",
	Run: func(cmd *cobra.Command, args []string) {
		status, err := daemon.KillDaemon()
		switch status {
		case -1:
			{
				logger.Errorln("Cannot find GPM daemon.")
				os.Exit(1)
			}
		case 0:
			{
				logger.Logln("Successfully killed GPM daemon.")
				os.Exit(0)
			}
		case 1:
			{
				logger.Errorln("Cannot kill GPM daemon.")
				logger.Errorln(err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(terminateCmd)
}
