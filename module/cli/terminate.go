package cli

import (
	"geep/module/client"
	"geep/module/daemon"
	"geep/module/logger"
	"os"

	"github.com/spf13/cobra"
)

var terminateCmd = &cobra.Command{
	Use:   "terminate",
	Short: "Terminate GEEP daemon process",
	Run: func(cmd *cobra.Command, args []string) {
		conn, reader, err := client.MakeUDSConn()
		if err != nil {
			logger.Errorln(err)
			os.Exit(1)
		}

		resultMessage, err := client.KillAll(conn, reader)
		if err != nil || !resultMessage.Success {
			logger.Errorln(err)
			os.Exit(1)
		}

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
