//go:build !windows

package daemon

import (
	"os"
	"os/exec"
	"syscall"
)

func setupDaemon(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setsid: true, // 새로운 세션 시작 (제어 터미널 분리)
	}
}

func isAlive(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}

	// Windows: FindProcess가 항상 프로세스 객체를 반환하므로 추가 확인 필요
	err = process.Signal(syscall.Signal(0))
	if err == nil {
		return true
	}

	return false
}
