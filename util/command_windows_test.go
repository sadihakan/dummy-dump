//go:build windows
// +build windows

package util

import (
	"testing"
)

func TestWhichWindows(t *testing.T) {
	if "C:\\Windows\\System32\\where.exe" != Which() {
		t.Error("which not defined")
	}
}
