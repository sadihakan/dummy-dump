package internal

import (
	"fmt"
	"github.com/sadihakan/dummy-dump/model"
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

//CreateCheckBinaryCommand
func CreateCheckBinaryCommand(sourceType model.SOURCE_TYPE)*exec.Cmd  {
	return exec.Command(util.Which(),getCheckCommand(sourceType)...)
}

// CreateImportBinaryCommand ...
func CreateImportBinaryCommand(sourceType model.SOURCE_TYPE) *exec.Cmd {
	return exec.Command(util.Which(), getImportCommand(sourceType)...)
}

// CreateExportBinaryCommand ...
func CreateExportBinaryCommand(sourceType model.SOURCE_TYPE) *exec.Cmd {
	return exec.Command(util.Which(), getExportCommand(sourceType)...)
}

// CreateExportCommand ...
func CreateExportCommand(binaryPath string, sourceType model.SOURCE_TYPE, user string, database string) *exec.Cmd {
	return exec.Command(binaryPath, getExportCommandArg(sourceType, user, database)...)
}

// CreateImportCommand ...
func CreateImportCommand(binaryPath string, sourceType model.SOURCE_TYPE, user string, database string, path string) *exec.Cmd {
	return exec.Command(binaryPath, getImportCommandArg(sourceType, user, database, path)...)
}

func getImportCommandArg(sourceType model.SOURCE_TYPE, user string, database, path string) (arg []string) {
	switch sourceType {
	case model.PostgreSQL:
		arg = []string{user, pgFlagDatabase, database, pgFlagCreate}
	case model.MySQL:
		arg = []string{mysqlFlagUser, user, mysqlFlagPassword, database, mysqlFlagExecute, "source " + path}
	case model.MSSQL:
		importQuery := fmt.Sprintf(`RESTORE DATABASE [%s] FROM DISK = '%s'`,
			database,
			path)
		arg = []string{mssqlFlagUser, user, mssqlFlagQuery, importQuery}
	}
	return arg
}

func getExportCommandArg(sourceType model.SOURCE_TYPE, user string, database string) (arg []string) {
	today := time.Now().UTC().UnixNano()
	filename := fmt.Sprintf("%d.backup", today)
	switch sourceType {
	case model.PostgreSQL:
		arg = []string{user, database, pgFlagFileName, filename, pgFlagCreate, pgFlatFormat}
	case model.MySQL:
		arg = []string{mysqlFlagUser, user, mysqlFlagPassword, database}
	case model.MSSQL:
		exportQuery := fmt.Sprintf(`BACKUP DATABASE [%s] TO DISK = '%s' WITH STATS = 10`,
			database,
			util.GetMSSQLBackupDirectory()+`\`+fmt.Sprintf("%d.bak", today))

		arg = []string{mssqlFlagUser, user, mssqlFlagQuery, exportQuery}
	}
	return arg
}

func getCheckCommand(sourceType model.SOURCE_TYPE) (command[]string){
	switch sourceType {
	case model.PostgreSQL:
		switch runtime.GOOS {
		case "darwin":
			command = []string{"psql"}
		case "linux":
			command = []string{"psql"}
		case "windows":
			command = []string{"/r", "C:\\Program Files\\Postgresql", "psql"}
		}
	case model.MySQL:
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

func getImportCommand(sourceType model.SOURCE_TYPE) (command []string) {
	switch sourceType {
	case model.PostgreSQL:
		switch runtime.GOOS {
		case "darwin":
			command = []string{"pg_restore"}
		case "linux":
			command = []string{"pg_restore"}
		case "windows":
			command = []string{"/r", "C:\\Program Files\\Postgresql", "pg_restore"}
		}
	case model.MySQL:
		switch runtime.GOOS {
		case "darwin":
			command = []string{"mysql"}
		case "linux":
			command = []string{"mysql"}
		case "windows":

			command = []string{"/r", "C:\\Program Files\\MySQL", "mysql"}

		}
	case model.MSSQL:
		switch runtime.GOOS {
		case "darwin":
		case "linux":
		case "windows":
			command = []string{"sqlcmd"}
		}
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
	case model.MSSQL:
		switch runtime.GOOS {
		case "darwin":
		case "linux":
		case "windows":
			command = []string{"sqlcmd"}
		}

	}

	return command
}
