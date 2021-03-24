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

func Name(name string) string{
if runtime.GOOS=="windows"{
	return "cmd"
}
return name
}