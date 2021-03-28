package internal

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/sadihakan/dummy-dump/model"
	"os"
	"os/exec"
)

type MSSQL struct {
	Dump
}

func (ms MSSQL) Check() error {
	cmd := exec.Command("reg", "query", "HKEY_LOCAL_MACHINE\\SOFTWARE\\Microsoft", "/f", "mssql", "/k")
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func (ms MSSQL) Export(binaryPath string, user string, database string) error {
	cmd := CreateExportCommand(binaryPath, model.MSSQL, user, database)
	fmt.Println(cmd)
	fmt.Println("---------")
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	return err
}

func (ms MSSQL) Import(binaryPath string, user string, database string, path string) error {
	cmd := CreateImportCommand(binaryPath, model.MSSQL, user, database, path)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	return err
}
