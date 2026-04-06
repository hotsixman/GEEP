//go:build windows

package daemon

import (
	"os/exec"
	"syscall"
	"unsafe"
)

var (
	modkernel32            = syscall.NewLazyDLL("kernel32.dll")
	procGetExitCodeProcess = modkernel32.NewProc("GetExitCodeProcess")
	procOpenProcess        = modkernel32.NewProc("OpenProcess")
)

const (
	PROCESS_QUERY_LIMITED_INFORMATION = 0x1000
	STILL_ACTIVE                      = 259
)

func setupDaemon(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,                             // 콘솔 창 숨기기
		CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP, // 독립된 프로세스 그룹 생성
	}
}

func isAlive(pid int) bool {
	handle, _, _ := procOpenProcess.Call(
		uintptr(PROCESS_QUERY_LIMITED_INFORMATION),
		0,
		uintptr(pid),
	)
	if handle == 0 {
		return false // 프로세스가 없거나 접근 권한 없음
	}
	defer syscall.CloseHandle(syscall.Handle(handle))

	// 2. 프로세스의 종료 코드를 확인합니다.
	var exitCode uint32
	ret, _, _ := procGetExitCodeProcess.Call(
		handle,
		uintptr(unsafe.Pointer(&exitCode)),
	)

	if ret == 0 {
		return false
	}

	// 종료 코드가 STILL_ACTIVE(259)이면 살아있는 것입니다.
	return exitCode == STILL_ACTIVE
}
