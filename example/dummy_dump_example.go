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
		Source:         config.MySQL,
		Import:         false,
		Export:         true,
		User:           "hakankosanoglu",
		Password:       "",
		DB:             "hell",
		Host:           "localhost",
		Port:           5432,
		BackupFilePath: "/Users/hakankosanoglu/Desktop",
		BackupName:     "aa.backup",
		BinaryPath:     "aa",
	})

	if err != nil {
		fmt.Println("DummyDump error ", err)
	}

	_, err = dd2.CheckPath().Export().Run()

	if err != nil {
		fmt.Println("Run error: ", err)
	}

}
