//go:build windows
// +build windows

package util

var windowsWhich = []string{
	"where.exe",
}

func Which() string {
	return "where.exe"
}
