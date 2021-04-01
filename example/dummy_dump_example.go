package main

import (
	"fmt"
	dummydump "github.com/sadihakan/dummy-dump"
	"github.com/sadihakan/dummy-dump/config"
)

func main() {
	dd , err:= dummydump.New(&config.Config{
		Source:     "postgres",
		Import:     false,
		Export:     true,
		User:       "hakankosanoglu",
		Password:   "",
		Path:       "",
		DB:         "test",
		BinaryPath: "/usr/local/opt/postgresql@12/bin/pg_dump",
		BackupName: "test",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(dd.GetBinary())


	if _, err = dd.Check().Export().Run(); err != nil {
		panic(err)
	}

}


