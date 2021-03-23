package internal

import (
	"bytes"
	"github.com/sadihakan/DummyDump/model"
	"os/exec"
	"strings"
)

func CheckBinary(binaryPath string, sourceType model.SOURCE_TYPE, importArg bool, exportArg bool) string {
	if binaryPath == "" {

		if importArg {
			binaryPath = checkImport(sourceType)
		}

		if exportArg {
			binaryPath = checkExport(sourceType)
		}
	}

	return binaryPath
}

func checkImport(sourceType model.SOURCE_TYPE) string {
	switch sourceType {
	case model.PostgreSQL:
		var out, _ bytes.Buffer
		cmd := exec.Command("which", "pg_restore")
		cmd.Stdout = &out
		err := cmd.Run()

		if err != nil {
			panic(err)
		}

		restore := out.String()

		return strings.TrimSuffix(restore, "\n")
	case model.MySQL:
		var out, _ bytes.Buffer
		cmd := exec.Command("which", "*")
		cmd.Stdout = &out
		err := cmd.Run()

		if err != nil {
			panic(err)
		}

		restore := out.String()

		return strings.TrimSuffix(restore, "\n")
	}

	return ""
}

func checkExport(sourceType model.SOURCE_TYPE) string {
	switch sourceType {
	case "postgres":
		var out, _ bytes.Buffer
		cmd := exec.Command("which", "pg_dump")
		cmd.Stdout = &out
		err := cmd.Run()

		if err != nil {
			panic(err)
		}

		restore := out.String()

		return strings.TrimSuffix(restore, "\n")
	case "mysql":
		var out, _ bytes.Buffer
		cmd := exec.Command("which", "*")
		cmd.Stdout = &out
		err := cmd.Run()

		if err != nil {
			panic(err)
		}

		restore := out.String()

		return strings.TrimSuffix(restore, "\n")
	}

	return ""
}
