package main

import (
	"flag"
	"github.com/sadihakan/dummy-dump/internal"
	"github.com/sadihakan/dummy-dump/model"
	"github.com/sadihakan/dummy-dump/util"
	"log"
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

	if !model.SOURCE_TYPE(sourceType).IsValid() {
		log.Println("invalid source type")
		return
	}

	if importArg && exportArg {
		log.Fatal("only one operation can be run")
	}

	password,err:=util.GetPassword()
	if err != nil {
		log.Fatalln(err)
	}

	var dump internal.Dump
	var source model.SOURCE_TYPE

	switch sourceType {
	case "postgres":
		dump = internal.Postgres{}
		source=model.PostgreSQL
	case "mysql":
		dump = internal.MySQL{}
		source=model.MySQL
	case "mssql":
		dump = internal.MSSQL{}
		source=model.MSSQL
	default:
		panic("")

	}

	if err = dump.Check(); err != nil {
		panic(err)
	}

	binaryPath = internal.CheckBinary(binaryPath, model.SOURCE_TYPE(sourceType), importArg, exportArg)

	config := model.Config{
		Source:     source,
		Import:     importArg,
		Export:     exportArg,
		User:       user,
		Password:   password,
		Path:       password,
		DB:         db,
		BinaryPath: binaryPath,
	}
	if importArg {
		if !util.PathExists(config.Path) {
			panic("Path is not exist")
		}

		if err := dump.Import(config); err != nil {
			panic(err)
		}
	}

	if exportArg {

		if err := dump.Export(config); err != nil {
			panic(err)
		}
	}
}
