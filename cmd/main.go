package main

import (
	"fmt"
	"github.com/sadihakan/DummyDump/database"
	"os"
	"runtime"
)

func main() {

	goos := runtime.GOOS
	switch goos {
	case "windows":
		fmt.Println("Windows")
	case "darwin":
		fmt.Println("MAC operating system")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", goos)
	}

	if len(os.Args) > 1 {

		postgres := database.Postgres{}

		err := postgres.Check()

		if err != nil {
			panic(err)
		}

		if os.Args[1] == "import" {
			user := os.Args[2]
			path := os.Args[3]

			err := postgres.Import(user, path)

			if err != nil {
				panic(err)
			}

		} else if os.Args[1] == "export" {

			user := os.Args[2]
			dbName := os.Args[3]

			err := postgres.Export(user, dbName)
			if err != nil {
				panic(err)
			}
		}
	}
}
