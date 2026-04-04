package daemon

import (
	"gpm/module/logger"
	"gpm/module/uds"
	"os"
	"os/exec"
	"time"
)

const DAEMON_ENV = "GPM_DAEMON_PROCESS"

func Daemonize() {
	if os.Getenv(DAEMON_ENV) == "1" {
		return
	}

	cmd := exec.Command(os.Args[0], os.Args[1:]...)
	cmd.Env = append(os.Environ(), DAEMON_ENV+"=1")

	setupDaemon(cmd)
	if err := cmd.Start(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}

func DaemonInit() {
	if os.Getenv(DAEMON_ENV) != "1" {
		return
	}

	log, err := logger.GetMainLogger()
	if err != nil {
		os.Exit(1)
	}

	udsServer, err := uds.Listen()
	if err != nil {
		log.Log(err)
		os.Exit(1)
	}

	log.SetUDSServer(udsServer)

	go func() {
		for {
			log.Log(time.Now())
			time.Sleep(time.Second)
		}
	}()

	select {}
}
