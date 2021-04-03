package main

import (
	dummydump "github.com/sadihakan/dummy-dump"
	"github.com/sadihakan/dummy-dump/config"
)

func main() {
	dd, err := dummydump.New()

	if err != nil {
		panic(err)
	}

	dd.SetBinaryPath("/usr/bin/psql", config.PostgreSQL, false, true)

	if _, err = dd.CheckPath().Run(); err != nil {
		panic(err)
	}

}
