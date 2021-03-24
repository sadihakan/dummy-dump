package internal

import (
	"fmt"
	"github.com/sadihakan/DummyDump/model"
	"github.com/sadihakan/DummyDump/util"
	"os/exec"
	"runtime"
	"time"
)

const (
	pgDatabase     = "--dbname=postgres"
	pgFlagFileName = "-f"
	pgFlagCreate   = "--create"
	pgFlatFormat   = "--format=c"
	pgRestore      = "pg_restore"
	pgDump         = "pg_dump"
	mysqlImport    = "mysql"
	mysqlExport    = "mysqldump"
)

// CreateImportBinaryCommand ...
func CreateImportBinaryCommand(sourceType model.SOURCE_TYPE) *exec.Cmd {
	return exec.Command(util.Which(), getImportCommand(sourceType))
}

// CreateExportBinaryCommand ...
func CreateExportBinaryCommand(sourceType model.SOURCE_TYPE) *exec.Cmd {
	return exec.Command(util.Which(), getExportCommand(sourceType)...)
}

// CreateExportCommand ...
func CreateExportCommand(name string, sourceType model.SOURCE_TYPE, user string, database string) *exec.Cmd {
	return exec.Command(util.Name(name), getExportCommandArg(sourceType, user, database)...)
}

// CreateImportCommand ...
func CreateImportCommand(name string, sourceType model.SOURCE_TYPE, user string, database string) *exec.Cmd {
	return exec.Command(util.Name(name), getImportCommandArg(sourceType, user, database)...)
}

func getImportCommandArg(sourceType model.SOURCE_TYPE, user string, path string) (arg []string) {
	switch sourceType {
	case model.PostgreSQL:
		switch runtime.GOOS {
		case "darwin":
			arg = []string{user, pgDatabase, path, pgFlagCreate}
		case "linux":
			arg = []string{user, pgDatabase, path, pgFlagCreate}
		case "windows":
			arg = []string{"/C", pgRestore, user, pgDatabase, path, pgFlagCreate}
		}
	case model.MySQL:

	}

	return arg
}

func getExportCommandArg(sourceType model.SOURCE_TYPE, user string, database string) (arg []string) {
	today := time.Now().UTC().UnixNano()
	filename := fmt.Sprintf("%d.backup", today)
	switch sourceType {
	case model.PostgreSQL:
		switch runtime.GOOS {
		case "darwin":
			arg = []string{user, database, pgFlagFileName, filename, pgFlagCreate, pgFlatFormat}
		case "linux":
			arg = []string{user, database, pgFlagFileName, filename, pgFlagCreate, pgFlatFormat}
		case "windows":
			arg = []string{"/C", pgDump, user, database, pgFlagFileName, filename, pgFlagCreate, pgFlatFormat}
		}
	case model.MySQL:
	}

	return arg
}

func getImportCommand(sourceType model.SOURCE_TYPE) (command string) {
	switch sourceType {
	case model.PostgreSQL:
		command = "pg_restore"
	case model.MySQL:
		command = "mysql"
	}

	return command
}

func getExportCommand(sourceType model.SOURCE_TYPE) (command []string) {
	switch sourceType {
	case model.PostgreSQL:
		switch runtime.GOOS {
		case "darwin":
			command = []string{"pg_dump"}
		case "linux":
			command = []string{"pg_dump"}
		case "windows":
			command = []string{"/r", "C:\\Program Files\\Postgresql", "pg_dump"}
		}
	case model.MySQL:
		switch runtime.GOOS {
		case "darwin":
			command = []string{"mysqldump"}
		case "linux":
			command = []string{"mysqldump"}
		case "windows":
			command = []string{"/r", "C:\\Program Files\\MySQL", "mysqldump"}
		}
	}

	return command
}
