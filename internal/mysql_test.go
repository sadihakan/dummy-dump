package internal

import (
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
	binaryPath := "mysql"
	d.Import(binaryPath, "testuser", "deneme","/home/onur123/go/src/github.com/onurcevik/DummyDump/test.sql")

}

func TestMySQL_ImportWithError(t *testing.T) {
	var d Dump
	d = MySQL{}

	binaryPath := "mysql"
	err := d.Import(binaryPath, "testuser", "deneme","/home/onur123/go/src/github.com/onurcevik/DummyDump/test.sql")

	if err != nil {
		t.Errorf("MySQL Import() Error: %s", err.Error())
	}
}

func TestMySQL_Export(t *testing.T) {
	var d Dump
	d = MySQL{}
	binaryPath := "mysqldump"
	d.Export(binaryPath, "testuser", "deneme")

}

func TestMySQL_ExportWithError(t *testing.T) {
	var d Dump
	d = MySQL{}
	binaryPath := "mysql"
	err := d.Export(binaryPath, "testuser", "deneme")
	if err != nil {
		t.Errorf("MySQL Export() Error: %s", err.Error())
	}
}
