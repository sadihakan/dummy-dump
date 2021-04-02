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

	dd.SetBinaryPath("/usr/local/opt/postgresql@12/bin/psqla", config.PostgreSQL)

	if _, err = dd.CheckPath().Run(); err != nil {
		panic(err)
	}

}


