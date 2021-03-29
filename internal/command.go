package internal

import (
	"fmt"
	"github.com/sadihakan/dummy-dump/config"
	"github.com/sadihakan/dummy-dump/util"
	"os/exec"
	"runtime"
	"time"
)

const (
	//pgDatabase     = "--dbname=postgres"
	pgFlagDatabase = "-d"
	pgFlagFileName = "-f"
	pgFlagCreate   = "--create"
	pgFlatFormat   = "--format=c"

	//pgRestore="pg_restore"
	//pgDump="pg_dump"
	mysqlDatabase     = "deneme"
	mysqlFlagUser     = "-u"
	mysqlFlagPassword = "-p"
	mysqlFlagExecute  = "-e"
	//mysqlImport="mysql"
	//mysqlDump="mysqldump"
	mssqlFlagUser  = "-U"
	mssqlFlagQuery = "-Q"
)

// CreateCheckBinaryCommand ...
func CreateCheckBinaryCommand(sourceType config.SourceType)*exec.Cmd  {
	return exec.Command(util.Which(),getCheckCommand(sourceType)...)
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
func CreateExportCommand(binaryPath string, sourceType config.SourceType, user string, database string) *exec.Cmd {
	return exec.Command(binaryPath, getExportCommandArg(sourceType, user, database)...)
}

// CreateImportCommand ...
func CreateImportCommand(binaryPath string, sourceType config.SourceType, user string, database string, path string) *exec.Cmd {
	return exec.Command(binaryPath, getImportCommandArg(sourceType, user, database, path)...)
}

// getImportCommandArg ...
func getImportCommandArg(sourceType config.SourceType, user string, database, path string) (arg []string) {
	switch sourceType {
	case config.PostgreSQL:
		arg = []string{user, pgFlagDatabase, database, pgFlagCreate}
	case config.MySQL:
		arg = []string{mysqlFlagUser, user, mysqlFlagPassword, database, mysqlFlagExecute, "source " + path}
	case config.MSSQL:
		importQuery := fmt.Sprintf(`RESTORE DATABASE [%s] FROM DISK = '%s'`,
			database,
			path)
		arg = []string{mssqlFlagUser, user, mssqlFlagQuery, importQuery}
	}
	return arg
}

// getExportCommandArg ...
func getExportCommandArg(sourceType config.SourceType, user string, database string) (arg []string) {
	today := time.Now().UTC().UnixNano()
	filename := fmt.Sprintf("%d.backup", today)
	switch sourceType {
	case config.PostgreSQL:
		arg = []string{user, database, pgFlagFileName, filename, pgFlagCreate, pgFlatFormat}
	case config.MySQL:
		arg = []string{mysqlFlagUser, user, mysqlFlagPassword, database}
	case config.MSSQL:
		exportQuery := fmt.Sprintf(`BACKUP DATABASE [%s] TO DISK = '%s' WITH STATS = 10`,
			database,
			util.GetMSSQLBackupDirectory()+`\`+fmt.Sprintf("%d.bak", today))

		arg = []string{mssqlFlagUser, user, mssqlFlagQuery, exportQuery}
	}
	return arg
}

// getCheckCommand ...
func getCheckCommand(sourceType config.SourceType) (command[]string){
	switch sourceType {
	case config.PostgreSQL:
		switch runtime.GOOS {
		case "darwin":
			command = []string{"psql"}
		case "linux":
			command = []string{"psql"}
		case "windows":
			command = []string{"/r", "C:\\Program Files\\Postgresql", "psql"}
		}
	case config.MySQL:
		switch runtime.GOOS {
		case "darwin":
			command = []string{"mysql"}
		case "linux":
			command = []string{"mysql"}
		case "windows":
			command = []string{"/r", "C:\\Program Files\\MySQL", "mysql"}
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
			command = []string{"/r", "C:\\Program Files\\Postgresql", "pg_restore"}
		}
	case config.MySQL:
		switch runtime.GOOS {
		case "darwin":
			command = []string{"mysql"}
		case "linux":
			command = []string{"mysql"}
		case "windows":

			command = []string{"/r", "C:\\Program Files\\MySQL", "mysql"}

		}
	case config.MSSQL:
		switch runtime.GOOS {
		case "darwin":
		case "linux":
		case "windows":
			command = []string{"sqlcmd"}
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
			command = []string{"/r", "C:\\Program Files\\Postgresql", "pg_dump"}
		}
	case config.MySQL:
		switch runtime.GOOS {
		case "darwin":
			command = []string{"mysqldump"}
		case "linux":
			command = []string{"mysqldump"}
		case "windows":
			command = []string{"/r", "C:\\Program Files\\MySQL", "mysqldump"}
		}
	case config.MSSQL:
		switch runtime.GOOS {
		case "darwin":
		case "linux":
		case "windows":
			command = []string{"sqlcmd"}
		}

	}

	return command
}
