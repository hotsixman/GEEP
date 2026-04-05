package pm

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gpm/module/logger"
	"gpm/module/types"
	"io"
	"os"
	"os/exec"
	"sync"
)

type logLine struct {
	Source string `json:"source"`
	Line   string `json:"line"`
}

func (pm *PM) NewProcess(startMessage types.StartMessage) error {
	process := exec.Command(startMessage.Run, startMessage.Args...)
	process.Dir = startMessage.Cwd
	env := os.Environ()
	for k, v := range startMessage.Env {
		env = append(env, fmt.Sprintf("%s=%s", k, v))
	}
	process.Env = env

	out, outWriter := io.Pipe()
	outMutex := &sync.Mutex{}
	stdin, err := process.StdinPipe()
	if err != nil {
		pm.mainLogger.Errorln(err)
		return err
	}
	stdout, err := process.StdoutPipe()
	if err != nil {
		pm.mainLogger.Errorln(err)
		return err
	}
	stderr, err := process.StderrPipe()
	if err != nil {
		pm.mainLogger.Errorln(err)
		return err
	}
	logger, err := logger.CreateLogger(startMessage.Name, true, pm.server)
	if err != nil {
		pm.mainLogger.Errorln(err)
		return err
	}
	go transform(pm.mainLogger, stdout, outWriter, outMutex, "log")
	go transform(pm.mainLogger, stderr, outWriter, outMutex, "error")

	err = process.Start()
	if err != nil {
		pm.mainLogger.Errorln(err)
		return err
	}
	pm.mainLogger.Logln("Process started:", startMessage.Name)

	pmProcess := &PMProcess{
		name:         startMessage.Name,
		process:      process,
		stdin:        bufio.NewWriter(stdin),
		out:          out,
		logger:       logger,
		startMessage: startMessage,
	}
	pm.process[startMessage.Name] = pmProcess

	go func() {
		scanner := bufio.NewScanner(pmProcess.out)
		for scanner.Scan() {
			var logline logLine
			err := json.Unmarshal([]byte(scanner.Text()), &logline)
			if err == nil {
				if logline.Source == "log" {
					pmProcess.logger.Logln(logline.Line)
					pm.mainLogger.Logln(logline.Line)
				} else if logline.Source == "error" {
					pmProcess.logger.Errorln(logline.Line)
					pm.mainLogger.Errorln(logline.Line)
				}

			} else {
				pm.mainLogger.Errorln(err)
			}
		}
	}()

	go func() {
		err := pmProcess.process.Wait()
		if err == nil {
			pmProcess.logger.Logln(fmt.Sprintf("Process exited. Error: %v", err))
		} else {
			pmProcess.logger.Errorln(fmt.Sprintf("Process exited. Error: %v", err))
		}

		delete(pm.process, pmProcess.startMessage.Name)
		stdin.Close()
		stdout.Close()
		stderr.Close()
	}()

	return nil
}

func transform(mainLogger *logger.Logger, reader io.Reader, writer io.Writer, mutex *sync.Mutex, source string) {
	scanner := bufio.NewScanner(reader)
	encoder := json.NewEncoder(writer)
	for scanner.Scan() {
		mutex.Lock()
		logline := logLine{
			Source: source,
			Line:   scanner.Text(),
		}
		err := encoder.Encode(logline)
		mutex.Unlock()
		if err != nil {
			mainLogger.Errorln(err)
			return
		}
	}
}
