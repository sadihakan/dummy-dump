package main

import (
	dummydump "github.com/sadihakan/dummy-dump"
	"github.com/sadihakan/dummy-dump/model"
)

func main() {
	dd, err := dummydump.New(&model.Config{
		Source:     model.PostgreSQL,
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

	dd.Check().Import().Run()
}


