package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gpm",
	Short: "GPM is a process manager for Go applications",
	Long:  `Go Process Manager (GPM) allows you to manage background processes with ease.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
