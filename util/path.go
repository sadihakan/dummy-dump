package util

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !os.IsNotExist(err)
}


// HomeDir returns the base directory + backup-app
// On Unix, including macOS, it returns the $HOME environment variable.
// On Windows, it returns %USERPROFILE%/backup-app.
// On Plan 9, it returns the $home environment variable.
func HomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return fmt.Sprintf("%s%s", home, string(filepath.Separator))
	}
	return os.Getenv("HOME")
}
