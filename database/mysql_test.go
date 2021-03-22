package database

import (
	"testing"
)

func TestMySQL_Check(t *testing.T) {
	var d Dump
	d=MySQL{}
	err:=d.Check()
	if err != nil {
		t.Errorf("MySQL Check() Error: %s",err.Error())
	}
}

func TestMySQL_Import(t *testing.T) {
	var d Dump
	d=MySQL{}
	err:=d.Import("testuser","/home/onur123/go/src/github.com/onurcevik/DummyDump/1616151265428723615.backup.sql")
	if err != nil {
		t.Errorf("MySQL Import() Error: %s",err.Error())
	}
}

func TestMySQL_Export(t *testing.T) {
	var d Dump
	d=MySQL{}
	err:=d.Export("testuser","deneme")
	if err != nil {
		t.Errorf("MySQL Export() Error: %s",err.Error())
	}
}