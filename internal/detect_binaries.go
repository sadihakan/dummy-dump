package internal

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/sadihakan/dummy-dump/config"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
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

	if runtime.GOOS == "windows" {
		_, file := filepath.Split(binaryPath)
		cmd = CreateVersionCommand(ctx, file, sourceType)
	} else {
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
		switch sourceType {
		case config.PostgreSQL:
			var path string
			for i := 0; i < len(predefinedPostgresPaths); i++ {
				path, err = findExport(ctx, predefinedPostgresPaths[i])
				if err == nil {
					break
				}
			}
			return path, nil
		case config.MySQL:
			var path string
			for i := 0; i < len(predefinedMySQLPaths); i++ {
				path, err = findExport(ctx, predefinedMySQLPaths[i])
				if err == nil {
					break
				}
			}
			return path, nil
		}
	}

	lines := strings.Split(out.String(), "\n")
	return strings.TrimSpace(lines[0]), nil
}

func findExport(ctx context.Context, path string) (string, error) {
	cmd := CreateCheckBinaryPathCommand(ctx, config.Config{BinaryPath: fmt.Sprintf("%s --version", path)})

	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return "", errors.New(fmt.Sprint(err) + ": " + stderr.String())
	}

	return path, nil
}
