package internal

import (
	"github.com/sadihakan/dummy-dump/config"
	"testing"
)

func TestMSSQL_Check(t *testing.T) {
	var d Dump
	d = MSSQL{}
	d.Check(nil)
}

func TestMSSQL_CheckWithError(t *testing.T) {
	var d Dump
	d = MSSQL{}
	err := d.Check(nil)
	if err != nil {
		t.Errorf("MSSQL Check() Error: %s", err.Error())
	}

}

func TestMSSQL_Import(t *testing.T) {
	var d Dump
	d = MSSQL{}
	config := config.Config{
		Source:         "",
		Import:         false,
		Export:         false,
		User:           "",
		Password:       "",
		BackupFilePath: "",
		DB:             "",
		BinaryPath:     "",
		BackupName:     "",
	}
	d.Import(nil, config)
}

func TestMSSQL_ImportWithError(t *testing.T) {
	var d Dump
	d = MSSQL{}

	config := config.Config{
		Source:         "",
		Import:         false,
		Export:         false,
		User:           "",
		Password:       "",
		BackupFilePath: "",
		DB:             "",
		BinaryPath:     "",
		BackupName:     "",
	}
	err := d.Import(nil, config)
	if err != nil {
		t.Errorf("MSSQL Import() Error: %s", err.Error())
	}
}

func TestMSSQL_Export(t *testing.T) {
	var d Dump
	d = MSSQL{}

	config := config.Config{
		Source:         "",
		Import:         false,
		Export:         false,
		User:           "",
		Password:       "",
		BackupFilePath: "",
		DB:             "",
		BinaryPath:     "",
		BackupName:     "",
	}
	d.Export(nil, config)
}

func TestMSSQL_ExportWithError(t *testing.T) {
	var d Dump
	d = MSSQL{}

	config := config.Config{
		Source:         "",
		Import:         false,
		Export:         false,
		User:           "",
		Password:       "",
		BackupFilePath: "",
		DB:             "",
		BinaryPath:     "",
		BackupName:     "",
	}
	err := d.Export(nil, config)
	if err != nil {
		t.Errorf("MSSQL Import() Error: %s", err.Error())
	}
}
