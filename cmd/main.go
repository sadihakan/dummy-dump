package main

import (
	"flag"
	"fmt"
	"github.com/sadihakan/DummyDump/internal"
	"github.com/sadihakan/DummyDump/internal/database"
	"github.com/sadihakan/DummyDump/model"
	"github.com/sadihakan/DummyDump/util"
	"log"
)

var (
	importArg bool
	exportArg bool
	sourceType string
	user string
	path string
	db string
	binaryPath string
)

func main() {
	flag.BoolVar(&importArg, "import", false, "Import process")
	flag.BoolVar(&exportArg, "export", false, "Export process")
	flag.StringVar(&sourceType, "source", "", "Source type is: mysql|postgres")
	flag.StringVar(&user, "user", "", "User name")
	flag.StringVar(&path, "path", "", "Import file path")
	flag.StringVar(&db, "db", "", "Database name")
	flag.StringVar(&binaryPath, "binaryPath", "", "Binary path")
	flag.Parse()

	fmt.Println(sourceType)

	if !model.SOURCE_TYPE(sourceType).IsValid()	{
			log.Println("invalid source type")
		return
	}

	if importArg && exportArg {
		log.Fatal("only one operation can be run")
	}

	var dump database.Dump

	switch sourceType {
	case "postgres":
		dump = database.Postgres{}
	case "mysql":
		//dump = database.MySQL{}
	}

	err := dump.Check()

	if err != nil {
		fmt.Println(err)
	}

	binaryPath = internal.CheckBinary(binaryPath, model.SOURCE_TYPE(sourceType), importArg, exportArg)

	if importArg {
		if !util.PathExists(path) {
			panic("Path is not exist")
		}

		err = dump.Import(binaryPath, user, path)

		if err != nil {
			panic(err)
		}
	}

	if exportArg {

		err = dump.Export(binaryPath, user, db)

		if err != nil {
			panic(err)
		}
	}
}
