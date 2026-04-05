package pm

import (
	"bufio"
	"gpm/module/logger"
	"gpm/module/types"
	"io"
	"os/exec"
)

type PM struct {
	process    map[string]*PMProcess
	mainLogger *logger.Logger
	server     types.ServerInterface
}

type PMProcess struct {
	name         string
	process      *exec.Cmd
	stdin        *bufio.Writer
	out          *io.PipeReader
	logger       *logger.Logger
	startMessage types.StartMessage
}

func NewPM(mainLogger *logger.Logger) *PM {
	pm := &PM{
		make(map[string]*PMProcess),
		mainLogger,
		nil,
	}

	return pm
}

func (pm *PM) SetServer(server types.ServerInterface) {
	pm.server = server
}

func (pm *PM) Input(name string, message string) {
	if pm.process[name] == nil {
		return
	}

	pm.process[name].stdin.Write([]byte(message))
}
