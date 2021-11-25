//go:build windows
// +build windows

package util

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

//The function below is for Windows systems only
//GetMSSQLBackupDirectory ...
func GetBackupDirectory() string {
	cmd := exec.Command("reg", "query", "HKEY_LOCAL_MACHINE\\SOFTWARE\\Microsoft\\MICROSOFT SQL SERVER", "/s", "/v", "/f", "BackupDirectory", "/k")
	var outb bytes.Buffer
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = &outb
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
	b, err := ioutil.ReadAll(&outb)
	if err != nil {
		log.Fatalln(err)
	}
	l := strings.Split(strings.Split(string(b), "\n")[2], "    ")

	return strings.TrimSuffix(l[3], "\r")
}
