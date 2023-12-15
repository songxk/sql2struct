package main

import (
	"os"
	"path/filepath"
)

func GetExecPath() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	absPath, err := filepath.Abs(exePath)
	if err != nil {
		return "", err
	}
	return filepath.Dir(absPath), nil
}
