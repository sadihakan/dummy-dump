package internal

import (
	"bytes"
	"github.com/sadihakan/dummy-dump/config"
	"log"
	"os"
	"os/exec"
	"strings"
)

// CheckBinary ...
func CheckBinary(binaryPath string, sourceType config.SourceType, importArg bool, exportArg bool) string {
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

func checkImport(sourceType config.SourceType) string {
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
	lines := strings.Split(out.String(), "\n")
	return strings.TrimSpace(lines[0])
}

func checkExport(sourceType config.SourceType) string {
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

	lines := strings.Split(out.String(), "\n")
	return strings.TrimSpace(lines[0])
}
