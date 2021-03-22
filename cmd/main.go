package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/sadihakan/DummyDump/database"
	"github.com/sadihakan/DummyDump/util"
	"log"
	"os/exec"
	"strings"
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
	pgDump string
	pgRestore string
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

		if binaryPath == "" {
			var out, _ bytes.Buffer
			cmd := exec.Command("which", "pg_restore")
			cmd.Stdout = &out
			err := cmd.Run()

			if err != nil {
				panic(err)
			}

			pgRestore = out.String()

			binaryPath = strings.TrimSuffix(pgRestore, "\n")
		}

		err = dump.Import(binaryPath, user, path)

		if err != nil {
			panic(err)
		}
	}

	if exportArg {

		if binaryPath == "" {
			binaryPath = strings.TrimSuffix(pgDump, "\n")

			var out, _ bytes.Buffer
			cmd := exec.Command("which", "pg_dump")
			cmd.Stdout = &out
			err := cmd.Run()

			if err != nil {
				panic(err)
			}

			pgDump = out.String()
		}

		err = dump.Export(binaryPath, user, db)

		if err != nil {
			panic(err)
		}
	}
}
