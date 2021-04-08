package internal

import (
	"github.com/sadihakan/dummy-dump/config"
	"testing"
)

func TestMySQL_Check(t *testing.T) {
	var d Dump
	d = MySQL{}
	d.Check()
}

func TestMySQL_CheckWithError(t *testing.T) {
	var d Dump
	d = MySQL{}
	err := d.Check()
	if err != nil {
		t.Errorf("MySQL Check() Error: %s", err.Error())
	}
}

func TestMySQL_Import(t *testing.T) {
	var d Dump
	d = MySQL{}

	config := config.Config{
		Source:         "",
		Import:         false,
		Export:         false,
		User:           "",
		Password:       "",
		BackupFilePath: "",
		DB:             "",
		BinaryPath:     "",
	}
	d.Import(config)

}

func TestMySQL_ImportWithError(t *testing.T) {
	var d Dump
	d = MySQL{}

	config := config.Config{
		Source:         "",
		Import:         false,
		Export:         false,
		User:           "",
		Password:       "",
		BackupFilePath: "",
		DB:             "",
		BinaryPath:     "",
	}
	err := d.Import(config)

	if err != nil {
		t.Errorf("MySQL Import() Error: %s", err.Error())
	}
}

func TestMySQL_Export(t *testing.T) {
	var d Dump
	d = MySQL{}
	config := config.Config{
		Source:         "",
		Import:         false,
		Export:         false,
		User:           "",
		Password:       "",
		BackupFilePath: "",
		DB:             "",
		BinaryPath:     "",
	}
	d.Export(config)

}

func TestMySQL_ExportWithError(t *testing.T) {
	var d Dump
	d = MySQL{}
	config := config.Config{
		Source:         "",
		Import:         false,
		Export:         false,
		User:           "",
		Password:       "",
		BackupFilePath: "",
		DB:             "",
		BinaryPath:     "",
	}
	err := d.Export(config)
	if err != nil {
		t.Errorf("MySQL Export() Error: %s", err.Error())
	}
}
