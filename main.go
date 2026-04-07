package main

import (
	"geep/module/cli"
	"geep/module/daemon"
	"os"
)

func main() {
	if os.Getenv(daemon.DAEMON_ENV) == "1" {
		daemon.DaemonInit()
	} else {
		cli.Execute()
	}
}
