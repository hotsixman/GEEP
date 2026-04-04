package cli

import (
	"bufio"
	"gpm/module/logger"
	"gpm/module/uds"
	"os"

	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to process",
	Run: func(cmd *cobra.Command, args []string) {
		closeChan := make(chan bool)

		client, err := uds.Connect(args[0], closeChan)
		if err != nil {
			logger.Errorln(err)
			os.Exit(1)
		}

		scanner := bufio.NewScanner(os.Stdin)
		go func() {
			for scanner.Scan() {
				command := scanner.Text()
				client.Command(command)
			}
		}()

		<-closeChan
	},
}

func init() {
	connectCmd.Flags().StringP("name", "", "", "Name of process")
	rootCmd.AddCommand(connectCmd)
}
