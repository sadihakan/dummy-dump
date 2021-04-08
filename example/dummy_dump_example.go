package main

import (
	dummydump "github.com/sadihakan/dummy-dump"
	"github.com/sadihakan/dummy-dump/config"
)

func main() {
	dd, err := dummydump.New(&config.Config{
		Source:         "postgres",
		Import:         false,
		Export:         true,
		User:           "hakankosanoglu",
		Password:       "",
		DB:             "test",
		Host:           "localhost",
		Port:           5432,
		BackupFilePath: "",
		BackupName:     "test.backup",
		BinaryPath:     "/usr/local/opt/postgresql@12/bin/pg_dump",
	})

	if err != nil {
		panic(err)
	}

	_, err = dd.Export().Run()

	if err != nil {
		panic(err)
	}

}
