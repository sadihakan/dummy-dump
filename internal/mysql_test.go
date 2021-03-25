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
<<<<<<< HEAD
	binaryPath := "mysql"
	d.Import(binaryPath, "testuser", "/home/onur123/go/src/github.com/onurcevik/DummyDump/test.sql")
=======
	d.Import("testuser", "/home/onur123/go/src/github.com/onurcevik/dummy-dump/test.sql")
>>>>>>> 9dcf83a030589ec764d633be08cccab1e1c7e59e
}

func TestMySQL_ImportWithError(t *testing.T) {
	var d Dump
	d = MySQL{}
<<<<<<< HEAD
	binaryPath := "mysql"
	err := d.Import(binaryPath, "testuser", "/home/onur123/go/src/github.com/onurcevik/DummyDump/test.sql")
=======
	err := d.Import("testuser", "/home/onur123/go/src/github.com/onurcevik/dummy-dump/test.sql")
>>>>>>> 9dcf83a030589ec764d633be08cccab1e1c7e59e
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
