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

	return filepath.Join(dir, ".gpm"), nil
}
