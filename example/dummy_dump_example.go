package main

import (
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
		BackupName: "",
	})

	if err != nil {
		panic(err)
	}


	if _, err = dd.Check().Export().Run(); err != nil {
		panic(err)
	}

}


