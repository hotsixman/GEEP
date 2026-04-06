package cli

import (
	"gpm/module/daemon"
	"gpm/module/logger"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Start the GPM daemon process",
	Run: func(cmd *cobra.Command, args []string) {
		if os.Getenv("GPM_DAEMON_PROCESS") == "1" {
			os.Exit(1)
		}
		logger.Logln("Starting GPM daemon in background...")
		status, err := daemon.SpawnDaemon()
		switch status {
		case -1:
			{
				logger.Errorln("???")
				os.Exit(1)
			}
		case 0:
			{
				logger.Logln("GPM daemon started successfully!")
				os.Exit(0)
			}
		case 1:
			{
				logger.Errorln("Cannot start GPM daemon")
				logger.Errorln(err)
				os.Exit(1)
			}
		case 2:
			{
				logger.Errorln("GPM daemon is already running")
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
