package main

import (
	dummydump "github.com/sadihakan/dummy-dump"
	"github.com/sadihakan/dummy-dump/config"
)

func main() {
	dd, err := dummydump.New(&config.Config{
		Source:     config.PostgreSQL,
		Import:     true,
		Export:     false,
		User:       "sadihakan",
		Path:       "/path",
		DB:         "db",
		BinaryPath: "/binaryPath",
	})

	if err != nil {
		panic(err)
	}

	if _, err = dd.Check().Export().Run(); err != nil {
		panic(err)
	}

}


