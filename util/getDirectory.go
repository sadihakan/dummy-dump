package util

import (
	"os"
	"path/filepath"
	"strings"
)

func GetDirectory() string {
	dir, _ := os.Getwd()
	return filepath.Join(strings.Split(dir, "DummyDump")[0], "DummyDump")
}
