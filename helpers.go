package main

import (
	"errors"
	"log"
	"os"
	"strings"
)

// takes a string and checks for a number of prefixes and suffixes then removes them
// and returns the string with trimed spaces
func toPlainDomain(s string) string {
	prefixes := []string{"#", "!", "*", "||", "0.0.0.0", "127.0.0.1"}
	suffixes := []string{"^"}

	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			s = strings.TrimSpace(s[len(prefix):])
		}
	}

	for _, suffix := range suffixes {
		if strings.HasSuffix(s, suffix) {
			s = strings.TrimSpace(s[0 : len(s)-len(suffix)])
		}
	}

	return strings.TrimSpace(s)
}

// take a directory path and returns false if doesn't exist, true if it does
func CheckPathExists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

// Finds the executalbe path
func FindExePath() string {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	return ex
}
