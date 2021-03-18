package main

import (
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

	if len(os.Args) > 1 {

		var dump database.Dump
		dump = database.Postgres{}

		err := dump.Check()

		if err != nil {
			panic(err)
		}

		if os.Args[1] == "import" {

			user := os.Args[2]
			path := os.Args[3]

			if !util.PathExists(path) {
				panic("Path is not exist")
			}

			err := dump.Import(user, path)

			if err != nil {
				panic(err)
			}

		} else if os.Args[1] == "export" {

			user := os.Args[2]
			dbName := os.Args[3]

			err := dump.Export(user, dbName)

			if err != nil {
				panic(err)
			}
		}
	}
}
