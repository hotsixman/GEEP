package db

import (
	"database/sql"
	"gpm/module/util"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func OpenDB() (*sql.DB, error) {
	dbPath, err := GetDBPath()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetDBPath() (string, error) {
	homeDir, err := util.GetHomeDirPath()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, "main.db"), nil
}
