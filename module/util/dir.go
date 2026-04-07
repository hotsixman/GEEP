package util

import (
	"os"
	"path/filepath"
)

func GetHomeDirPath() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, ".geep"), nil
}

func GetUDSPath() string {
	home, _ := os.UserHomeDir()
	dir := filepath.Join(home, ".geep")
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}
	return filepath.Join(dir, "geep.sock")
}
