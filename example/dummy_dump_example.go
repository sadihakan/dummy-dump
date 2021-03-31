package main

import (
	dummydump "github.com/sadihakan/dummy-dump"
	"github.com/sadihakan/dummy-dump/config"
)

func main() {
	dd , err:= dummydump.New(&config.Config{
		Source:     "mysql",
		Import:     false,
		Export:     true,
		User:       "testuser",
		Password:   "123456",
		Path:       "",
		DB:         "deneme",
		BinaryPath: "/usr/bin/mysqldump",
		BackupName: "deneme",
	})

	if err != nil {
		panic(err)
	}

	if _, err = dd.Check().Export().Run(); err != nil {
		panic(err)
	}

}


