package uds

import (
	"os"
	"path/filepath"
)

func GetSocketPath() string {
	home, _ := os.UserHomeDir()
	dir := filepath.Join(home, ".gpm")
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}
	return filepath.Join(dir, "gpm.sock")
}
