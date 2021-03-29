package main

import (
	"flag"
	"github.com/sadihakan/dummy-dump/config"
	"github.com/sadihakan/dummy-dump/internal"
	"github.com/sadihakan/dummy-dump/util"
)

var (
	importArg  bool
	exportArg  bool
	sourceType string
	user       string
	password   string
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

	dumpConfig := config.Config{
		Source:     config.SourceType(sourceType),
		Import:     importArg,
		Export:     exportArg,
		User:       user,
		Path:       path,
		DB:         db,
		BinaryPath: binaryPath,
	}

	password, err := util.GetPassword()
	if err != nil {
		log.Fatalln(err)
	}
	config.Config{
		Source:     "",
		Import:     false,
		Export:     false,
		User:       "",
		Path:       "",
		DB:         "",
		BinaryPath: "",
	}

	var dump internal.Dump

	switch sourceType {
	case "postgres":
		if err := dumpConfig.CheckConfigPostgreSQL(); err != nil {
			panic(err)
		}
		dump = internal.Postgres{}

	case "mysql":
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
		panic("")
	}

	if err = dump.Check(); err != nil {
		panic(err)
	}

	binaryPath = internal.CheckBinary(binaryPath, config.SourceType(sourceType), importArg, exportArg)

	if importArg {
		if !util.PathExists(path) {
			panic("Path is not exist")
		}

		if err := dump.Import(binaryPath, user, db, path); err != nil {
			panic(err)
		}
	}

	if exportArg {

		if err := dump.Export(binaryPath, user, db); err != nil {
			panic(err)
		}
	}
}
