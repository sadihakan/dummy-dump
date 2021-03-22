package main

import (
	"flag"
	"fmt"
	"github.com/sadihakan/DummyDump/database"
	"github.com/sadihakan/DummyDump/util"
	"log"
)

var sourceTypes = []string{
	"mysql", "postgres",
}

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

	if e, _ := util.InArray(sourceType, sourceTypes); !e {
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

	if importArg {
		if !util.PathExists(path) {
			panic("Path is not exist")
		}

		err := dump.Import(binaryPath, user, path)

		if err != nil {
			panic(err)
		}
	}

	if exportArg {
		err := dump.Export(binaryPath, user, db)

		if err != nil {
			panic(err)
		}
	}
}
