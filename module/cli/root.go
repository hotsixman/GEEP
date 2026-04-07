package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "geep",
	Short: "GEEP is a process manager for Go applications",
	Long:  `GEEP (Go + Keep) allows you to manage background processes with ease.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
