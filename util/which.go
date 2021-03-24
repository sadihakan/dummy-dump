package util

import "runtime"

// Which ...
func Which() (which string) {
	switch runtime.GOOS {
	case "darwin":
		which = "which"
	case "linux":
		which = "which"
	case "windows":
		which = "where"
	}

	return which
}
