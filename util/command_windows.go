//go:build windows
// +build windows

package util

import (
	"fmt"
	"os"
	"path/filepath"
)

var windowsWhich = []string{
	filepath.Join(homeDrive(), "Windows", "System32", "where.exe"),
}

func Which() string {
	return checkCommands(windowsWhich)
}

func homeDrive() string {
	home := os.Getenv("HOMEDRIVE")
	if home == "" {
		home = "C:"
	}
	return fmt.Sprintf("%s%s", home, string(filepath.Separator))
}
