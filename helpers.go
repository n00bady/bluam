package main

import (
	"errors"
	"log"
	"os"
)

func CheckPathExists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

func FindExePath() string {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	return ex
}
