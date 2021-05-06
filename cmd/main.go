package main

import (
	"flag"
	"fmt"
	"github.com/sadihakan/dummy-dump/config"
	"github.com/sadihakan/dummy-dump/errors"
	"github.com/sadihakan/dummy-dump/internal"
	"github.com/sadihakan/dummy-dump/util"
	"path/filepath"
)

var (
	source         string
	importArg      bool
	exportArg      bool
	user           string
	password       string
	db             string
	host           string
	port           int
	backupFilePath string
	backupName     string
	binaryPath     string
	scanStyle      bool
)

func main() {
	flag.StringVar(&source, "source", "", "Source type is: mysql|postgres|mssql")
	flag.BoolVar(&importArg, "import", false, "Import process")
	flag.BoolVar(&exportArg, "export", false, "Export process")
	flag.StringVar(&user, "user", "", "User name")
	flag.StringVar(&backupFilePath, "path", "", "Import file path")
	flag.StringVar(&db, "db", "", "Database name")
	flag.StringVar(&host, "host", "", "Host name")
	flag.IntVar(&port, "port", 0, "Port number")
	flag.StringVar(&binaryPath, "binaryPath", "", "Binary path")
	flag.StringVar(&backupName, "backupName", "", "Backup name")
	flag.StringVar(&backupFilePath, "backupFilePath", "", "Backup file path")
	flag.BoolVar(&scanStyle, "scan", false, "")
	flag.Parse()

	if scanStyle {
		fmt.Println("Source type is: <mysql|postgres|mssql>")
		if _, err := fmt.Scanln(&source); err != nil {
			panic(errors.New("Source cannot be nil"))
		}

		fmt.Println("Import: <true|false>")
		if _, err := fmt.Scanln(&importArg); err != nil {
			panic(errors.New("Import cannot be nil"))
		}

		fmt.Println("Export: <true|false>")
		if _, err := fmt.Scanln(&exportArg); err != nil {
			panic(errors.New("Export cannot be nil"))
		}

		fmt.Println("User name is: ")
		if _, err := fmt.Scanln(&user); err != nil {
			panic(errors.New("User name cannot be nil"))
		}

		fmt.Println("Password is: ")
		password, _ = util.GetPassword()

		fmt.Println("Database is: ")
		if _, err := fmt.Scanln(&db); err != nil {
			panic(errors.New("Database cannot be nil"))
		}

		fmt.Println("Host is: ")
		if _, err := fmt.Scanln(&host); err != nil {
			panic(errors.New("Host cannot be nil"))
		}

		fmt.Println("Port is: ")
		if _, err := fmt.Scanln(&port); err != nil {
			panic(errors.New("Port cannot be nil"))
		}

		fmt.Println("Backup name is: ")
		if _, err := fmt.Scanln(&backupName); err != nil {
			panic(err)
		}

		fmt.Println("Backup file path is: ")
		if _, err := fmt.Scanln(&backupFilePath); err != nil {
			panic(err)
		}

		fmt.Println("Binary Path is (Press Enter if you dont know path): ")
		if _, err := fmt.Scanln(&binaryPath); err != nil {
			if err.Error() != "unexpected newline" {
				fmt.Println(err)
			}
		}
	} else {
		password, _ = util.GetPassword()
	}

	dumpConfig := config.Config{
		Source:         config.SourceType(source),
		Import:         importArg,
		Export:         exportArg,
		User:           user,
		Password:       password,
		DB:             db,
		Host:           host,
		Port:           port,
		BackupFilePath: backupFilePath,
		BackupName:     backupName,
		BinaryPath:     binaryPath,
	}

	var dump internal.Dump

	switch source {
	case "postgres":
		binaryPath, err := internal.CheckBinary(binaryPath, config.SourceType(source), importArg, exportArg)
		if err != nil {
			panic(err)
		}
		dumpConfig.BinaryPath = binaryPath
		if err = dumpConfig.CheckConfigPostgreSQL(); err != nil {
			panic(err)
		}
		dump = internal.Postgres{}

	case "mysql":
		binaryPath, err := internal.CheckBinary(binaryPath, config.SourceType(source), importArg, exportArg)
		if err != nil {
			panic(err)
		}
		dumpConfig.BinaryPath = binaryPath
		if err := dumpConfig.CheckConfigMySQL(); err != nil {
			panic(err)
		}
		dump = internal.MySQL{}

	case "mssql":
		if err := dumpConfig.CheckConfigMsSQL(); err != nil {
			panic(err)
		}
		dump = internal.MSSQL{}

	default:
		fmt.Println("\nYou did not chose a valid database source. " +
			"Currently supported databases are:\n  - MySQL\n  - PostgreSQL\n  - MSSQL\n")
		return
	}

	if err := dump.Check(); err != nil {
		panic(err)
	}

	if importArg {
		if !util.PathExists(filepath.Join(dumpConfig.BackupFilePath, dumpConfig.BackupName)) {
			panic("Path is not exist")
		}
		if err := dump.Import(dumpConfig); err != nil {
			panic(err)
		}
	}

	if exportArg {
		if err := dump.Export(dumpConfig); err != nil {
			panic(err)
		}
	}
}
