package main

import (
	"fmt"
	dummydump "github.com/sadihakan/dummy-dump"
	"github.com/sadihakan/dummy-dump/config"
)

func main() {
	dd, err := dummydump.New()

	if err != nil {
		fmt.Println("DummyDump new error: ", err)
	}

	dd.SetBinaryConfig(config.PostgreSQL, false, true)

	binary, version := dd.GetBinary()
	fmt.Println("Bin: ", binary)
	fmt.Println("Version: ", version)

	dd2, err := dummydump.New(&config.Config{
		Source:         config.PostgreSQL,
		Import:         false,
		Export:         true,
		User:           "communication_prod",
		Password:       "a8LvLKnkC784xvUg2F6cxHsTM3bxxJ8G",
		DB:             "communication_prod_db",
		Host:           "34.159.36.0",
		Port:           5432,
		BackupFilePath: "/Users/hakankosanoglu/Desktop",
		BackupName:     "hell.backup",
		BinaryPath:     binary,
	})

	if err != nil {
		fmt.Println("DummyDump error ", err)
	}

	_, err = dd2.CheckPath().Export().Run()

	if err != nil {
		fmt.Println("Run error: ", err)
	}

}
