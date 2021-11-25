//go:build linux
// +build linux

package util

import "testing"

func TestWhichLinux(t *testing.T) {
	if "/usr/bin/which" != Which() {
		t.Error("which not defined")
	}
}
