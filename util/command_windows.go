//go:build windows
// +build windows

package util

import (
	"fmt"
	"path/filepath"
)

var windowsWhich = []string{
	filepath.Join(homeDrive(), "Windows", "System32", "where.exe"),
}

func Which() string {
	return checkCommands(windowsWhich)
}

func homeDrive() {
	home := os.Getenv("HOMEDRIVE")
	if home == "" {
		return fmt.Sprintf("C%s%s", filepath.ListSeparator, filepath.Separator)
	}
}
