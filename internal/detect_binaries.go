package internal

import (
	"bytes"
	"context"
	"github.com/sadihakan/dummy-dump/config"
	"os"
	"os/exec"
	"strings"
)

// CheckBinary ...
func CheckBinary(ctx context.Context, binaryPath string, sourceType config.SourceType, importArg bool, exportArg bool) (string, error) {
	var err error
	if binaryPath == "" {
		if importArg {
			binaryPath, err = checkImport(ctx, sourceType)
			if err != nil {
				return "", err
			}
		}

		if exportArg {
			binaryPath, err = checkExport(ctx, sourceType)
			if err != nil {
				return "", err
			}
		}

	}
	return binaryPath, nil
}

// CheckVersion ...
func CheckVersion(ctx context.Context, binaryPath string, sourceType config.SourceType) (string, error) {
	version, err := checkVersion(ctx, binaryPath, sourceType)

	if err != nil {
		return "", err
	}

	return version, nil
}

func checkVersion(ctx context.Context, binaryPath string, sourceType config.SourceType) (string, error) {
	var out bytes.Buffer
	var cmd *exec.Cmd

	if sourceType == config.Oracle {

	}
	switch sourceType {
	case config.Oracle:
		//cmd = CreateVersionCommand(filepath.Join(oraclelBinaryDirectories()[0], "sqlplus.exe"), sourceType)
	default:
		cmd = CreateVersionCommand(ctx, binaryPath, sourceType)
	}

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

func checkImport(ctx context.Context, sourceType config.SourceType) (string, error) {
	var out bytes.Buffer
	var cmd *exec.Cmd

	cmd = CreateImportBinaryCommand(ctx, sourceType)

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

func checkExport(ctx context.Context, sourceType config.SourceType) (string, error) {
	var out bytes.Buffer
	var cmd *exec.Cmd

	cmd = CreateExportBinaryCommand(ctx, sourceType)

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
