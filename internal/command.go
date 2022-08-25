package internal

import (
	"context"
	"fmt"
	"github.com/sadihakan/dummy-dump/config"
	"github.com/sadihakan/dummy-dump/util"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

const (
	pgDatabase           = "--dbname=postgres"
	pgFlagDatabase       = "-d"
	pgFlagFileName       = "-f"
	pgFlagCreateDatabase = "-C"
	pgFlagCreate         = "--create"
	pgFlagFormat         = "--format=c"
	pgVersion            = "--version"
	pgFlagPassword       = "-W"
	mysqlVersion         = "--version"
	oracleVersion        = "-V"

	mysqlFlagUser       = "--user"
	mysqlFlagPassword   = "--password"
	mysqlFlagExecute    = "-e"
	mysqlFlagResultFile = "--result-file"

	host    = "--host="
	port    = "--port="
	noOwner = "-x"
)

// CheckBinaryPathCommand ...
func CheckBinaryPathCommand(ctx context.Context, cfg config.Config) *exec.Cmd {
	return homeDirCommand(ctx, cfg.BinaryPath, getVersionCommandArg(cfg.Source))
}

// CheckVersionCommand ...
func CheckVersionCommand(ctx context.Context, binaryPath string, sourceType config.SourceType) *exec.Cmd {
	return homeDirCommand(ctx, binaryPath, getVersionCommandArg(sourceType))
}

// CreateExportCommand ...
func CreateExportCommand(ctx context.Context, cfg config.Config) *exec.Cmd {
	return homeDirCommand(ctx, cfg.BinaryPath, getExportCommandArg(cfg))
}

// CreateImportCommand ...
func CreateImportCommand(ctx context.Context, cfg config.Config) *exec.Cmd {
	return homeDirCommand(ctx, cfg.BinaryPath, getImportCommandArg(cfg))
}

func homeDirCommand(ctx context.Context, command string, arg []string) *exec.Cmd {
	params := make([]string, 0)
	cmd := new(exec.Cmd)
	switch runtime.GOOS {
	case "darwin":
		if command != "" {
			params = append(params, command)
		}
		params = append(params, arg...)
		cmd = exec.CommandContext(ctx, "sh", "-c", strings.Join(params, " "))
	case "linux":
		cmd = exec.CommandContext(ctx, command, arg...)
	case "windows":
		if command != "" {
			params = append(params, strconv.Quote(command))
		}
		params = append(params, arg...)
		cmd = exec.CommandContext(ctx, "powershell.exe", "&", strings.TrimSpace(strings.Join(params, " ")))
	}

	cmd.Dir = util.HomeDir()

	return cmd
}

// getVersionCommandArg ...
func getVersionCommandArg(sourceType config.SourceType) (arg []string) {
	switch sourceType {
	case config.PostgreSQL:
		arg = []string{pgVersion}
	case config.MySQL:
		arg = []string{mysqlVersion}
	case config.Oracle:
		arg = []string{oracleVersion}
	}
	return arg
}

// getImportCommandArg ...
func getImportCommandArg(cfg config.Config) (arg []string) {
	host := fmt.Sprintf("%s%s", host, cfg.Host)
	port := fmt.Sprintf("%s%d", port, cfg.Port)
	switch cfg.Source {
	case config.PostgreSQL:
		arg = []string{"-d", fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB), filepath.Join(cfg.BackupFilePath, cfg.BackupName)}
	case config.MySQL:
		user := fmt.Sprintf("%s=%s", mysqlFlagUser, cfg.User)
		password := fmt.Sprintf("%s=%s", mysqlFlagPassword, cfg.Password)
		arg = []string{user, password, host, port, cfg.DB, mysqlFlagExecute, "source " + filepath.Join(cfg.BackupFilePath, cfg.BackupName)}
	case config.Oracle:
		//impdp user/password@service schemas=db1,db2.. directory=directory dumpfile=filename
		connStr := fmt.Sprintf("%s/%s@%s", cfg.User, cfg.Password, cfg.Service)
		directory := fmt.Sprintf("directory=%s", cfg.BackupFilePath)
		filename := fmt.Sprintf("dumpfile=%s", cfg.BackupName)
		schemas := fmt.Sprintf("schemas=%s", cfg.DB)
		arg = []string{connStr, schemas, directory, filename}
	}
	return arg
}

// getExportCommandArg ...
func getExportCommandArg(cfg config.Config) (arg []string) {
	host := fmt.Sprintf("%s%s", host, cfg.Host)
	port := fmt.Sprintf("%s%d", port, cfg.Port)
	switch cfg.Source {
	case config.PostgreSQL:
		filename := filepath.Join(cfg.BackupFilePath, cfg.BackupName)
		dns := fmt.Sprintf(`--dbname=postgresql://%s:%s@%s:%d/%s`, cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
		arg = []string{dns, pgFlagFormat, noOwner, pgFlagFileName, filename}
	case config.MySQL:
		filename := filepath.Join(cfg.BackupFilePath, cfg.BackupName)
		user := fmt.Sprintf(`%s=%s`, mysqlFlagUser, cfg.User)
		password := fmt.Sprintf(`%s=%s`, mysqlFlagPassword, cfg.Password)
		resultFile := fmt.Sprintf(`%s=%s`, mysqlFlagResultFile, filename)
		arg = []string{user, password, host, port, cfg.DB, resultFile}
	case config.Oracle:
		//expdp user/password@db schemas=DB1,DB2... directory=directory dumpfile=filename
		connStr := fmt.Sprintf("%s/%s@%s", cfg.User, cfg.Password, cfg.Service)
		directory := fmt.Sprintf("directory=%s", cfg.BackupFilePath)
		filename := fmt.Sprintf("dumpfile=%s", cfg.BackupName)
		schemas := fmt.Sprintf("schemas=%s", cfg.DB)
		arg = []string{connStr, schemas, directory, filename}
	}
	return arg
}
