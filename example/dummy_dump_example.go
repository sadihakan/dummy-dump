package main

import (
	dummydump "github.com/sadihakan/dummy-dump"
	"github.com/sadihakan/dummy-dump/config"
)

func main() {

	cfg := &config.Config{
		Source:     config.MySQL,
		Import:     false,
		Export:     true,
		User:       "testuser",
		Password:   "123456",
		Path:       "C:\\Program Files\\MySQL\\MySQL Server 8.0\\bin\\mysqldump.exe",
		DB:         "deneme",
		BinaryPath: "C:\\Program Files\\MySQL\\MySQL Server 8.0\\bin\\mysqldump.exe",
		BackupName: "C:\\Users\\onur\\Desktop\\blah\\denefasdfme.bak",
		Host:       "localhost",
		Port:       3306,
	}

	dd, err := dummydump.New(cfg)
	if err != nil {
		panic(err)
	}

	if _, err = dd.CheckPath().Export().Run(); err != nil {
		panic(err)
	}

}
