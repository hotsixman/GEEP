package pm

import (
	"bufio"
	"fmt"
	"gpm/module/logger"
	"gpm/module/types"
	"os/exec"
)

type PM struct {
	process map[string]*Process
	log     *logger.Logger
}

type Process struct {
	name   string
	cmd    *exec.Cmd
	args   []string
	stdin  *bufio.Writer
	stdout *bufio.Reader
	stderr *bufio.Reader
	logger *logger.Logger
}

func NewPM(log *logger.Logger) *PM {
	pm := &PM{
		make(map[string]*Process),
		log,
	}

	return pm
}

func (pm *PM) NewProcess(name string, udsServer types.UDSServerInterface, args ...string) error {
	cmd := exec.Command(args[0], args[1:]...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		pm.log.Errorln(err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		pm.log.Errorln(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		pm.log.Errorln(err)
	}
	logger, err := logger.CreateLogger(name, true, udsServer)
	if err != nil {
		pm.log.Errorln(err)
	}

	err = cmd.Start()
	if err != nil {
		pm.log.Errorln(err)
	}
	pm.log.Logln("Process started:", name)

	process := &Process{
		name:   name,
		cmd:    cmd,
		args:   args,
		stdin:  bufio.NewWriter(stdin),
		stdout: bufio.NewReader(stdout),
		stderr: bufio.NewReader(stderr),
		logger: logger,
	}
	pm.process[name] = process

	go func() {
		for {
			message, err := process.stdout.ReadString('\n')
			if err != nil {
				return
			}
			process.logger.Logln(message)
		}
	}()
	go func() {
		for {
			message, err := process.stderr.ReadString('\n')
			if err != nil {
				return
			}
			process.logger.Errorln(message)
		}
	}()

	go func() {
		err := cmd.Wait()
		process.logger.Logln(fmt.Sprintf("Process exited. Error: %v", err))

		delete(pm.process, name)
		stdin.Close()
		stdout.Close()
		stderr.Close()
	}()

	return nil
}

func (pm *PM) Input(name string, message string) {
	if pm.process[name] == nil {
		return
	}

	pm.process[name].stdin.Write([]byte(message))
}
