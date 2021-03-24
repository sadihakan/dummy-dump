package internal

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
	db := "test"

	binaryPath := "pg_restore"

	dump.Export(binaryPath, user, db)
}

func TestExportWithError(t *testing.T) {

	var dump Dump
	dump = Postgres{}

	user := "none"
	db := "test"

	binaryPath := "pg_restore"

	dump.Export(binaryPath, user, db)
}

func TestImport(t *testing.T) {

	var dump Dump
	dump = Postgres{}

	user := "hakankosanoglu"
	file := filepath.Join(util.GetDirectory(), "test.backup")

	binaryPath := "pg_dump"

	dump.Import(binaryPath, user, file)
}

func TestImportWithError(t *testing.T) {

	var dump Dump
	dump = Postgres{}

	user := "hakankosanoglu"
	file := "test"

	binaryPath := "pg_dump"

	dump.Import(binaryPath, user, file)
}
