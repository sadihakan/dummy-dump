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
	d.Import("testuser", "/home/onur123/go/src/github.com/onurcevik/dummy-dump/test.sql")
}

func TestMySQL_ImportWithError(t *testing.T) {
	var d Dump
	d = MySQL{}
	err := d.Import("testuser", "/home/onur123/go/src/github.com/onurcevik/dummy-dump/test.sql")
	if err != nil {
		t.Errorf("MySQL Import() Error: %s", err.Error())
	}
}

func TestMySQL_Export(t *testing.T) {
	var d Dump
	d = MySQL{}
	d.Export("testuser", "deneme")

}

func TestMySQL_ExportWithError(t *testing.T) {
	var d Dump
	d = MySQL{}
	err := d.Export("testuser", "deneme")
	if err != nil {
		t.Errorf("MySQL Export() Error: %s", err.Error())
	}
}
