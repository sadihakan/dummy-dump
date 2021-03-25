package internal

import (
	"bytes"
	"github.com/sadihakan/dummy-dump/model"
	"log"
	"os"
	"os/exec"
	"strings"
)

// CheckBinary ...
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
	var out bytes.Buffer
	var cmd *exec.Cmd

	cmd = CreateImportBinaryCommand(sourceType)

	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSuffix(out.String(), "\n")
}

func checkExport(sourceType model.SOURCE_TYPE) string {
	var out bytes.Buffer
	var cmd *exec.Cmd

	cmd = CreateExportBinaryCommand(sourceType)

	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSuffix(out.String(), "\n")
}
