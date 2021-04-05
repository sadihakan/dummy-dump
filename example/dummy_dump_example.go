package main

import (
	dummydump "github.com/sadihakan/dummy-dump"
	"github.com/sadihakan/dummy-dump/config"
)

func main() {
	//dd, err := dummydump.New(&config.Config{
	//	Source:     "postgres",
	//	Import:     false,
	//	Export:     true,
	//	User:       "atma",
	//	Password:   "rMASZk8KHrNJfzU4JzgS",
	//	Path:       "",
	//	DB:         "atma",
	//	BinaryPath: "/usr/local/opt/postgresql@12/bin/pg_dump",
	//	BackupName: "/Users/hakankosanoglu/Desktop/backup.backup",
	//	Host:       "atmaproductiondb.cfenfmd1itz9.eu-central-1.rds.amazonaws.com",
	//	Port:       5432,
	//})

	dd, err := dummydump.New(&config.Config{
		Source:     "postgres",
		Import:     true,
		Export:     false,
		User:       "postgres",
		Password:   "",
		Path:       "/Users/hakankosanoglu/Desktop/backup.backup",
		DB:         "postgres",
		BinaryPath: "/usr/local/opt/postgresql@12/bin/pg_restore",
		BackupName: "backup.backup",
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
