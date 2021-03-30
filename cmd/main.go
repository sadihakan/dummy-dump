package main

import (
	"flag"
	"fmt"
	"github.com/sadihakan/dummy-dump/config"
	"github.com/sadihakan/dummy-dump/internal"
	"github.com/sadihakan/dummy-dump/util"
	"log"
)

var (
	importArg  bool
	exportArg  bool
	sourceType string
	user       string
	path       string
	db         string
	binaryPath string
)

func main() {

	flag.BoolVar(&importArg, "import", false, "Import process")
	flag.BoolVar(&exportArg, "export", false, "Export process")
	flag.StringVar(&sourceType, "source", "", "Source type is: mysql|postgres|mssql")
	flag.StringVar(&user, "user", "", "User name")
	flag.StringVar(&path, "path", "", "Import file path")
	flag.StringVar(&db, "db", "", "Database name")
	flag.StringVar(&binaryPath, "binaryPath", "", "Binary path")
	flag.Parse()

	var err error
	binaryPath, err = internal.CheckBinary(binaryPath, config.SourceType(sourceType), importArg, exportArg)
	if err != nil {
		panic(err)
	}

	dumpConfig := config.Config{
		Source:     config.SourceType(sourceType),
		Import:     importArg,
		Export:     exportArg,
		User:       user,
		Path:       path,
		DB:         db,
		BinaryPath: binaryPath,
		Password:   "",
	}

	var dump internal.Dump

	switch sourceType {
	case "postgres":
		if err := dumpConfig.CheckConfigPostgreSQL(); err != nil {
			panic(err)
		}

		dump = internal.Postgres{}

	case "mysql":
		password, err := util.GetPassword()

		if err != nil {
			log.Fatalln(err)
		}

		dumpConfig.Password = password

		if err := dumpConfig.CheckConfigMySQL(); err != nil {
			panic(err)
		}

		dump = internal.MySQL{}

	case "mssql":
		password, err := util.GetPassword()

		if err != nil {
			log.Fatalln(err)
		}

		dumpConfig.Password = password

		if err := dumpConfig.CheckConfigMsSQL(); err != nil {
			panic(err)
		}

		dump = internal.MSSQL{}

	default:
		fmt.Println("\nYou did not chose a valid database source. Currently supported databases are:\n  - MySQL\n  - PostgreSQL\n  - MSSQL\n")
		return
	}

	if err := dump.Check(); err != nil {
		panic(err)
	}

	if importArg {
		if !util.PathExists(path) {
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
