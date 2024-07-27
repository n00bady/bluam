package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// takes a string and checks for a number of prefixes and suffixes then removes them
// and returns the string with trimed spaces
func toPlainDomain(s string) string {
	// Check prefixes 1st then another pass for suffixes
	switch {
	case strings.HasPrefix(s, ":"):
		return ""
	case strings.HasPrefix(s, "["):
		return ""
	case strings.HasPrefix(s, "#"):
		return ""
	case strings.HasPrefix(s, "!"):
		return ""
	case strings.HasPrefix(s, "*"):
		s = strings.TrimSpace(s[1:])
	case strings.HasPrefix(s, "||"):
		s = strings.TrimSpace(s[2:])
	case strings.HasPrefix(s, "0.0.0.0"):
		s = strings.TrimSpace(s[len("0.0.0.0"):])
	case strings.HasPrefix(s, "127.0.0.1"):
		s = strings.TrimSpace(s[len("127.0.0.1"):])
	default:
		return strings.TrimSpace(s)
	}

	if strings.HasSuffix(s, "^") {
		s = strings.TrimSpace(s[0 : len(s)-1])
	}

	return s
}

func runCmd(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func blocklistsChanged() (bool, error) {
	cmd := exec.Command("git", "diff", "dns")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return false, err
	}

	if len(out) > 0 {
		return true, nil
	}

	return false, nil
}
