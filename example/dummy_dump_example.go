package main

import (
	"context"
	"fmt"
	dummydump "github.com/sadihakan/dummy-dump"
	"github.com/sadihakan/dummy-dump/config"
)

func main() {
	dd, err := dummydump.New()

	if err != nil {
		fmt.Println("DummyDump new error: ", err)
	}

	dd.SetBinaryConfig(config.MySQL, false, true)

	ctx := context.Background()

	binary, version := dd.GetBinary(ctx)
	fmt.Println("Bin: ", binary)
	fmt.Println("Version: ", version)

	dd2, err := dummydump.New(&config.Config{
		Source:         config.MySQL,
		Import:         false,
		Export:         true,
		User:           "root",
		Password:       "123456",
		DB:             "testdb",
		Host:           "localhost",
		Port:           3306,
		BackupFilePath: "C:\\Users\\Administrator\\Desktop",
		BackupName:     "testdb.backup",
		BinaryPath:     binary,
	})

	if err != nil {
		fmt.Println("DummyDump error ", err)
	}

	_, err = dd2.Export(ctx).Run()

	if err != nil {
		fmt.Println("Run error: ", err)
	}

}
