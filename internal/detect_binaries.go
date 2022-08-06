package internal

import (
	"bytes"
	"errors"
	"github.com/sadihakan/dummy-dump/config"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// CheckBinary ...
func CheckBinary(binaryPath string, sourceType config.SourceType, importArg bool, exportArg bool) (string, error) {
	var err error
	if binaryPath == "" {
		if importArg {
			binaryPath, err = checkImport(sourceType)
			if err != nil {
				return "", err
			}
		}

		if exportArg {
			binaryPath, err = checkExport(sourceType)
			if err != nil {
				return "", err
			}
		}

	}
	return binaryPath, nil
}

// CheckVersion ...
func CheckVersion(binaryPath string, sourceType config.SourceType) (string, error) {
	version, err := checkVersion(binaryPath, sourceType)

	if err != nil {
		return "", err
	}

	return version, nil
}

func checkVersion(binaryPath string, sourceType config.SourceType) (string, error) {
	var out bytes.Buffer
	var cmd *exec.Cmd
	defer func() {
		out.Reset()
	}()

	switch sourceType {
	case config.Oracle:
		cmd = CreateVersionCommand(filepath.Join(oraclelBinaryDirectories()[0], "sqlplus.exe"), sourceType)
	default:
		cmd = CreateVersionCommand(binaryPath, sourceType)
	}
	setCmdDirectory(cmd)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		return "", err
	}
	lines := strings.Split(out.String(), "\n")
	return strings.TrimSpace(lines[0]), nil
}

func checkImport(sourceType config.SourceType) (string, error) {
	var out bytes.Buffer
	var cmd *exec.Cmd
	defer func() {
		out.Reset()
	}()

	cmd = CreateImportBinaryCommand(sourceType)
	setCmdDirectory(cmd)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		return "", err
	}
	lines := strings.Split(out.String(), "\n")
	return strings.TrimSpace(lines[0]), nil
}

func checkExport(sourceType config.SourceType) (string, error) {
	var out bytes.Buffer
	var errOut bytes.Buffer
	var cmd *exec.Cmd
	defer func() {
		out.Reset()
		errOut.Reset()
	}()

	cmd = CreateExportBinaryCommand(sourceType)
	setCmdDirectory(cmd)
	cmd.Stderr = &errOut
	cmd.Stdin = os.Stdin
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	if errOut.Len() > 0 {
		return "", errors.New(errOut.String())
	}

	lines := strings.Split(out.String(), "\n")
	return strings.TrimSpace(lines[0]), nil
}
