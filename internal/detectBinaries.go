package internal

import (
	"bytes"
	"github.com/sadihakan/DummyDump/model"
	"os"
	"os/exec"
	"strings"
)

func CheckBinary(opsys string, binaryPath string, sourceType model.SOURCE_TYPE, importArg bool, exportArg bool) string {
	if binaryPath == "" {

		if importArg {
			binaryPath = checkImport(opsys, sourceType)
		}

		if exportArg {
			binaryPath = checkExport(opsys, sourceType)
		}
	}
	return binaryPath
}

func checkImport(opsys string, sourceType model.SOURCE_TYPE) string {
	var out bytes.Buffer
	var cmd *exec.Cmd
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = &out
	switch sourceType {
	case model.PostgreSQL:
		if opsys == "windows" {
			cmd = exec.Command("where", "/r", "C:\\Program Files\\Postgresql", "pg_restore")
		} else {
			cmd = exec.Command("which", "pg_restore")
		}
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
		return strings.TrimSuffix(out.String(), "\n")
	case model.MySQL:
		if opsys == "windows" {
			cmd = exec.Command("where", "/r", "C:\\Program Files\\MySQL", "mysql")
		} else {
			cmd = exec.Command("which", "mysql")
		}
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
		return strings.TrimSuffix(out.String(), "\n")
	}
	return ""
}

func checkExport(opsys string, sourceType model.SOURCE_TYPE) string {
	var out bytes.Buffer
	var cmd *exec.Cmd
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = &out
	switch sourceType {
	case model.PostgreSQL:
		if opsys == "windows" {
			cmd = exec.Command("where", "/r", "C:\\Program Files\\Postgresql", "pg_dump")
		} else {
			cmd = exec.Command("which", "pg_dump")
		}
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
		return strings.TrimSuffix(out.String(), "\n")
	case model.MySQL:
		if opsys == "windows" {
			cmd = exec.Command("where", "/r", "C:\\Program Files\\MySQL", "mysqldump")
		} else{
			cmd = exec.Command("which", "mysqldump")
		}
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
		return strings.TrimSuffix(out.String(), "\n")
	}
	return ""
}
