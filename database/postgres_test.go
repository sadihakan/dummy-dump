package database

import (
	"github.com/sadihakan/DummyDump/util"
	"path/filepath"
	"testing"
)

func TestCheck(t *testing.T) {
	var dump Dump
	dump = Postgres{}

	dump.Check()
}

func TestCheckWithError(t *testing.T) {
	var dump Dump
	dump = Postgres{}

	err := dump.Check()

	if err != nil {
		t.Fatal(err)
	}
}

func TestExport(t *testing.T) {

	var dump Dump
	dump = Postgres{}

	user := "hakankosanoglu"
	dbName := "test"

	dump.Export(user, dbName)
}

func TestExportWithError(t *testing.T) {

	var dump Dump
	dump = Postgres{}

	user := "none"
	dbName := "test"

	dump.Export(user, dbName)
}

func TestImport(t *testing.T) {

	var dump Dump
	dump = Postgres{}

	user := "hakankosanoglu"
	file := filepath.Join(util.GetDirectory(), "1616145448738317000.backup")

	dump.Import(user, file)
}

func TestImportWithError(t *testing.T) {

	var dump Dump
	dump = Postgres{}

	user := "hakankosanoglu"
	file := "test"

	dump.Import(user, file)
}
