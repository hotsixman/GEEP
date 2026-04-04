package main

import (
	"gpm/module/cli"
	"gpm/module/daemon"
	"os"
)

func main() {
	if os.Getenv(daemon.DAEMON_ENV) == "1" {
		daemon.DaemonInit()
	} else {
		cli.Execute()
	}
}
