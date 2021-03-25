package util

import (
	"os"
	"path/filepath"
)

//func GetDirectory() string {
//	dir, _ := os.Getwd()
//	return filepath.Join(strings.Split(dir, "dummy-dump")[0], "dummy-dump")
//}

func GetDirectory() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}
