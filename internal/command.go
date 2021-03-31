package internal

import (
	"fmt"
	"github.com/sadihakan/dummy-dump/config"
	"github.com/sadihakan/dummy-dump/util"
	"os/exec"
	"runtime"
	"strings"
)

const (
	//pgDatabase     = "--dbname=postgres"
	pgFlagDatabase       = "-d"
	pgFlagFileName       = "-f"
	pgFlagCreateDatabase = "-C"
	pgFlagCreate         = "--create"
	pgFlagFormat         = "--format=c"

	//pgRestore="pg_restore"
	//pgDump="pg_dump"
	mysqlDatabase     = "deneme"
	mysqlFlagUser     = "-u"
	mysqlFlagPassword = "-p"
	mysqlFlagExecute  = "-e"
	//mysqlImport="mysql"
	//mysqlDump="mysqldump"
)

// CreateCheckBinaryCommand ...
func CreateCheckBinaryCommand(sourceType config.SourceType) *exec.Cmd {
	return exec.Command(util.Which(), getCheckCommand(sourceType)...)
}

// CreateImportBinaryCommand ...
func CreateImportBinaryCommand(sourceType config.SourceType) *exec.Cmd {
	return exec.Command(util.Which(), getImportCommand(sourceType)...)
}

// CreateExportBinaryCommand ...
func CreateExportBinaryCommand(sourceType config.SourceType) *exec.Cmd {
	return exec.Command(util.Which(), getExportCommand(sourceType)...)
}

// CreateExportCommand ...
func CreateExportCommand(cfg config.Config) *exec.Cmd {
	return exec.Command(cfg.BinaryPath, getExportCommandArg(cfg)...)
}

// CreateImportCommand ...
func CreateImportCommand(cfg config.Config) *exec.Cmd {
	return exec.Command(cfg.BinaryPath, getImportCommandArg(cfg)...)
}

// getImportCommandArg ...
func getImportCommandArg(cfg config.Config) (arg []string) {
	switch cfg.Source {
	case config.PostgreSQL:
		dns := fmt.Sprintf(`user=%s password=%s dbname=%s`, cfg.User, cfg.Password, cfg.DB)
		arg = []string{dns, pgFlagCreateDatabase, pgFlagCreate, cfg.Path}
	case config.MySQL:
		user := fmt.Sprintf("%s=%s", mysqlFlagUser, cfg.User)
		password := fmt.Sprintf("%s=%s", mysqlFlagPassword, cfg.Password)
		arg = []string{user, password, cfg.DB, mysqlFlagExecute, "source " + cfg.Path}

	}
	return arg
}

// getExportCommandArg ...
func getExportCommandArg(cfg config.Config) (arg []string) {
	filename := fmt.Sprintf("%s/%s.backup", cfg.Path, cfg.BackupName)
	switch cfg.Source {
	case config.PostgreSQL:
		dns := fmt.Sprintf(`user=%s password=%s dbname=%s`, cfg.User, cfg.Password, cfg.DB)
		arg = []string{dns, pgFlagFileName, filename, pgFlagCreate, pgFlagFormat}
	case config.MySQL:
		user := fmt.Sprintf("%s%s", mysqlFlagUser, cfg.User)
		password := fmt.Sprintf("%s%s", mysqlFlagPassword, cfg.Password)
		arg = []string{user,password, cfg.DB}
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
			p := strings.TrimSpace(mysqlBinaryDirectory())
			command = []string{"/r", p, "mysql"}
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
			p := strings.TrimSpace(mysqlBinaryDirectory())
			command = []string{"/r", p, "mysql"}
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
			p := strings.TrimSpace(mysqlBinaryDirectory())
			command = []string{"/r", p, "mysqldump"}

		}
	}
	return command
}
