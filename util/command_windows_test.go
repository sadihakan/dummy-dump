//go:build windows
// +build windows

package util

func TestWhichDarwin(t *testing.T) {
	if "/usr/bin/which" != Which() {
		t.Error("which not defined")
	}
}
