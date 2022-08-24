package internal

import (
	"bytes"
	"context"
	"errors"
	"fmt"
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

	cmd = CheckVersionCommand(ctx, binaryPath, sourceType)

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

func checkImport(_ context.Context, sourceType config.SourceType) (string, error) {
	var path string

	switch sourceType {
	case config.PostgreSQL:
		path, _ = exec.LookPath("pg_restore")
	case config.MySQL:
		path, _ = exec.LookPath("mysql")
	case config.Oracle:
		path, _ = exec.LookPath("impdp")
	}

	return path, nil
}

func checkExport(ctx context.Context, sourceType config.SourceType) (string, error) {
	var err error
	var path string

	switch sourceType {
	case config.PostgreSQL:
		path, _ = exec.LookPath("pg_dump")
	case config.MySQL:
		path, _ = exec.LookPath("mysqldump")
	case config.Oracle:
		path, _ = exec.LookPath("expdp")
	}

	if path != "" {
		path, err = findAndValidateExport(ctx, path)
		if err != nil {
			return path, nil
		}
	} else {
		switch sourceType {
		case config.PostgreSQL:
			for i := 0; i < len(predefinedPostgresPaths); i++ {
				path, err = findAndValidateExport(ctx, predefinedPostgresPaths[i])
				if err == nil {
					break
				}
			}
			return path, nil
		case config.MySQL:
			for i := 0; i < len(predefinedMySQLPaths); i++ {
				path, err = findAndValidateExport(ctx, predefinedMySQLPaths[i])
				if err == nil {
					break
				}
			}
			return path, nil
		}
	}

	return path, nil
}

func findAndValidateExport(ctx context.Context, path string) (string, error) {
	cmd := CheckBinaryPathCommand(ctx, config.Config{BinaryPath: fmt.Sprintf("%s --version", path)})

	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return "", errors.New(fmt.Sprint(err) + ": " + stderr.String())
	}

	return path, nil
}
