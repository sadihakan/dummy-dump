package internal

import (
	"context"
	"fmt"
	"github.com/sadihakan/dummy-dump/config"
	"github.com/sadihakan/dummy-dump/util"
	"os/exec"
	"path/filepath"
	"runtime"
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

	mysqlDatabase       = "deneme"
	mysqlFlagUser       = "--user"
	mysqlFlagPassword   = "--password"
	mysqlFlagExecute    = "-e"
	mysqlFlagResultFile = "--result-file"

	host    = "--host="
	port    = "--port="
	noOwner = "-x"
)

// CreateCheckBinaryCommand ...
func CreateCheckBinaryCommand(ctx context.Context, sourceType config.SourceType) *exec.Cmd {
	return homeDirCommand(ctx, util.Which(), getCheckCommand(sourceType))
}

// CreateCheckBinaryPathCommand ...
func CreateCheckBinaryPathCommand(ctx context.Context, cfg config.Config) *exec.Cmd {
	return homeDirCommand(ctx, "", getCheckBinaryPathCommand(cfg))
}

// CreateImportBinaryCommand ...
func CreateImportBinaryCommand(ctx context.Context, sourceType config.SourceType) *exec.Cmd {
	return homeDirCommand(ctx, util.Which(), getImportCommand(sourceType))
}

// CreateVersionCommand ...
func CreateVersionCommand(ctx context.Context, binaryPath string, sourceType config.SourceType) *exec.Cmd {
	return homeDirCommand(ctx, binaryPath, getVersionCommandArg(sourceType))
}

// CreateExportBinaryCommand ...
func CreateExportBinaryCommand(ctx context.Context, sourceType config.SourceType) *exec.Cmd {
	return homeDirCommand(ctx, util.Which(), getExportCommand(sourceType))
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
		command = strings.Join(params, " ")
		fmt.Println(command)
		cmd = exec.CommandContext(ctx, "sh", "-c", command)
	case "linux":
		cmd = exec.CommandContext(ctx, command, arg...)
	case "windows":
		if command != "" {
			params = append(params, command)
		}
		params = append(params, arg...)
		command = strings.Join(params, " ")
		cmd = exec.CommandContext(ctx, "powershell.exe", strings.TrimSpace(command))
	}

	cmd.Dir = util.HomeDir()

	return cmd
}

func getCheckBinaryPathCommand(cfg config.Config) (command []string) {
	switch runtime.GOOS {
	case "darwin":
		command = []string{cfg.BinaryPath}
	case "linux":
		command = []string{cfg.BinaryPath}
	case "windows":
		command = []string{cfg.BinaryPath}
	}
	return command
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

// getCheckCommand ...
func getCheckCommand(sourceType config.SourceType) (command []string) {
	switch sourceType {
	case config.PostgreSQL:
		switch runtime.GOOS {
		case "darwin":
			command = []string{"psql"}
		case "linux":
			command = []string{"psql"}
		case "windows":
			p := strings.TrimSpace(postgresqlBinaryDirectories()[0])
			command = []string{"/r", p, "psql"}
		}
	case config.MySQL:
		switch runtime.GOOS {
		case "darwin":
			command = []string{"mysql"}
		case "linux":
			command = []string{"mysql"}
		case "windows":
			p := strings.TrimSpace(mysqlBinaryDirectory()[0])
			command = []string{"/r", p, "mysqldump.exe"}
		}
	case config.Oracle:
		switch runtime.GOOS {
		case "darwin":
			command = []string{"oracle"}
		case "linux":
			command = []string{"oracle"}
		case "windows":
			p := strings.TrimSpace(oraclelBinaryDirectories()[0])
			command = []string{"/r", p, "oracle"}
		}
	}
	return command
}

// getImportCommand ...
func getImportCommand(sourceType config.SourceType) (command []string) {
	switch sourceType {
	case config.PostgreSQL:
		switch runtime.GOOS {
		case "darwin":
			command = []string{"pg_restore"}
		case "linux":
			command = []string{"pg_restore"}
		case "windows":
			p := strings.TrimSpace(postgresqlBinaryDirectories()[0])
			command = []string{"/r", p, "pg_restore"}
		}
	case config.MySQL:
		switch runtime.GOOS {
		case "darwin":
			command = []string{"mysql"}
		case "linux":
			command = []string{"mysql"}
		case "windows":
			p := strings.TrimSpace(mysqlBinaryDirectory()[0])
			command = []string{"/r", p, "mysql"}
		}
	case config.Oracle:
		switch runtime.GOOS {
		case "darwin":
			command = []string{"impdp"}
		case "linux":
			command = []string{"impdp"}
		case "windows":
			p := strings.TrimSpace(oraclelBinaryDirectories()[0])
			command = []string{"/r", p, "impdp"}
		}
	}
	return command
}

// getExportCommand ...
func getExportCommand(sourceType config.SourceType) (command []string) {
	switch sourceType {
	case config.PostgreSQL:
		switch runtime.GOOS {
		case "darwin":
			command = []string{"pg_dump"}
		case "linux":
			command = []string{"pg_dump"}
		case "windows":
			p := strings.TrimSpace(postgresqlBinaryDirectories()[0])
			command = []string{"/r", p, "pg_dump"}
		}
	case config.MySQL:
		switch runtime.GOOS {
		case "darwin":
			command = []string{"mysqldump"}
		case "linux":
			command = []string{"mysqldump"}
		case "windows":
			command = []string{"mysqldump.exe"}
		}
	case config.Oracle:
		switch runtime.GOOS {
		case "darwin":
			command = []string{"expdp"}
		case "linux":
			command = []string{"expdp"}
		case "windows":
			p := strings.TrimSpace(oraclelBinaryDirectories()[0])
			command = []string{"/r", p, "expdp"}
		}
	}
	return command
}
