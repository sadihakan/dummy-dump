//go:build darwin
// +build darwin

package util

import "testing"

func TestWhichDarwin(t *testing.T) {
	if "/usr/bin/which" != Which() {
		t.Error("which not defined")
	}
}
