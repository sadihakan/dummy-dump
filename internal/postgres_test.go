package internal

import (
	"github.com/sadihakan/dummy-dump/config"
	"github.com/sadihakan/dummy-dump/util"
	"path/filepath"
	"testing"
)

func TestCheck(t *testing.T) {
	var dump Dump
	dump = Postgres{}

	dump.Check(nil)
}

func TestCheckWithError(t *testing.T) {
	var dump Dump
	dump = Postgres{}

	err := dump.Check(nil)

	if err != nil {
		t.Fatal(err)
	}
}

func TestExport(t *testing.T) {

	var dump Dump
	dump = Postgres{}

	config := config.Config{
		Source:         "",
		Import:         false,
		Export:         true,
		User:           "hakankosanoglu",
		Password:       "",
		BackupFilePath: "",
		DB:             "test",
		BinaryPath:     "pg_dump",
	}

	dump.Export(nil, config)
}

func TestExportWithError(t *testing.T) {

	var dump Dump
	dump = Postgres{}

	config := config.Config{
		Source:         "",
		Import:         false,
		Export:         true,
		User:           "none",
		Password:       "",
		BackupFilePath: "",
		DB:             "test",
		BinaryPath:     "pg_dump",
	}

	dump.Export(nil, config)
}

func TestImport(t *testing.T) {

	var dump Dump
	dump = Postgres{}

	file := filepath.Join(util.GetDirectory(), "test.backup")

	config := config.Config{
		Source:         "",
		Import:         true,
		Export:         false,
		User:           "hakankosanoglu",
		Password:       "",
		BackupFilePath: file,
		DB:             "",
		BinaryPath:     "pg_restore",
	}

	dump.Import(nil, config)
}

func TestImportWithError(t *testing.T) {

	var dump Dump
	dump = Postgres{}
	file := "test"

	config := config.Config{
		Source:         "",
		Import:         true,
		Export:         false,
		User:           "hakankosanoglu",
		Password:       "",
		BackupFilePath: file,
		DB:             "",
		BinaryPath:     "pg_restore",
	}

	dump.Import(nil, config)
}
