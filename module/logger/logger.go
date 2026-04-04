package logger

import (
	"fmt"
	"gpm/module/uds"
	"gpm/module/util"
	"log"
	"os"
	"path/filepath"
)

type Logger struct {
	dirPath        string
	name           string
	timeRecorded   bool
	errorSeperated bool
	udsServer      *uds.UDSServer
}

func GetMainLogger() (*Logger, error) {
	homeDir, err := util.GetHomeDirPath()
	if err != nil {
		return nil, err
	}

	dirPath := filepath.Join(homeDir, "log")
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			return nil, err
		}
	}

	return &Logger{
		dirPath,
		"",
		true,
		false,
		nil,
	}, nil
}

func GetLogger(name string, timeRecorded bool, errorSeperated bool, udsServer *uds.UDSServer) (*Logger, error) {
	homeDir, err := util.GetHomeDirPath()
	if err != nil {
		return nil, err
	}

	dirPath := filepath.Join(homeDir, "log-process", filepath.Clean(name))
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			return nil, err
		}
	}

	return &Logger{
		dirPath,
		name,
		timeRecorded,
		errorSeperated,
		udsServer,
	}, nil
}

func (this *Logger) SetUDSServer(udsServer *uds.UDSServer) {
	this.udsServer = udsServer
}

func (this *Logger) Log(v ...any) {
	message := fmt.Sprintln(v...)
	if this.udsServer != nil {
		this.udsServer.Broadcast(message)
	}
	log.Println(message)
}
