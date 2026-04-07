package cli

import (
	"geep/module/daemon"
	"geep/module/logger"
	"os"

	"github.com/spf13/cobra"
)

var terminateCmd = &cobra.Command{
	Use:   "terminate",
	Short: "Terminate GEEP daemon process",
	Run: func(cmd *cobra.Command, args []string) {
		status, err := daemon.KillDaemon()
		switch status {
		case -1:
			{
				logger.Errorln("Cannot find GEEP daemon.")
				os.Exit(1)
			}
		case 0:
			{
				logger.Logln("Successfully killed GEEP daemon.")
				os.Exit(0)
			}
		case 1:
			{
				logger.Errorln("Cannot kill GEEP daemon.")
				logger.Errorln(err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(terminateCmd)
}
