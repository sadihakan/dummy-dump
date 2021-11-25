//go:build linux
// +build linux

package util

var linuxWhich = []string{
	"/usr/bin/which",
}

func Which() string {
	return checkCommands(linuxWhich)
}
