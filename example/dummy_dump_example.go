package main

import (
	dummydump "github.com/sadihakan/dummy-dump"
	"github.com/sadihakan/dummy-dump/config"
)

func main() {
	dd, err := dummydump.New(&config.Config{
		Source:     "postgres",
		Import:     true,
		Export:     false,
		User:       "hakankosanoglu",
		Password:   "",
		Path:       "/Users/hakankosanoglu/Desktop/backup.backup",
		DB:         "test",
		BinaryPath: "/usr/local/opt/postgresql@12/bin/pg_restore",
		BackupName: "",
		Host:       "localhost",
		Port:       5432,
	})

	if err != nil {
		panic(err)
	}

	_, err = dd.Import().Run()

	if err != nil {
		panic(err)
	}

}
