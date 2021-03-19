package main

import (
	"flag"
	"fmt"
	"github.com/sadihakan/DummyDump/database"
	"github.com/sadihakan/DummyDump/util"
	"os"
	"runtime"
)

func main() {

	switch runtime.GOOS {
	case "windows":
		fmt.Println("Windows")
	case "darwin":
		fmt.Println("MAC operating system")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", runtime.GOOS)
	}

	flag.Parse()

	if len(os.Args) > 1 {
		var dump database.Dump
		if os.Args[1]=="postgres"{
			dump = database.Postgres{}
		} else if os.Args[1]=="mysql"{
			dump = database.MySQL{}
		}

		err := dump.Check()

		if err != nil {
			panic(err)
		}

		if os.Args[2] == "import" {

			user := os.Args[3]
			path := os.Args[4]
			if !util.PathExists(path) {
				panic("Path is not exist")
			}

			err := dump.Import(user, path)

			if err != nil {
				panic(err)
			}

		} else if os.Args[2] == "export" {

			user := os.Args[3]
			dbName := os.Args[4]

			err := dump.Export(user, dbName)

			if err != nil {
				panic(err)
			}
		}
	}
}
