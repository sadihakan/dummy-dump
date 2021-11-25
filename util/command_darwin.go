//go:build darwin
// +build darwin

package util

var darwinWhich = []string{
	"/usr/bin/which",
}

func Which() string {
	return checkCommands(darwinWhich)
}
