package internal

import (
	"bytes"
	"fmt"
	"github.com/sadihakan/dummy-dump/model"
	"os"
)

type Postgres struct {
	Dump
}

func (p Postgres) Check() error {
	cmd := CreateCheckBinaryCommand(model.PostgreSQL)
	var out, errBuf bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	err := cmd.Run()
	if err != nil {
		_, _ = os.Stderr.WriteString(err.Error())
		return err
	}
	return nil
}

func (p Postgres) Export(binaryPath string, user string, database string) error {
	var out, errBuf bytes.Buffer

	user = fmt.Sprintf("--username=%s", user)
	database = fmt.Sprintf("--dbname=%s", database)

	cmd := CreateExportCommand(binaryPath, model.PostgreSQL, user, database)
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	err := cmd.Run()

	if err != nil {
		return err
	}

	return nil
}

func (p Postgres) Import(binaryPath string, user string, database string, path string) error {
	var out, errBuf bytes.Buffer

	user = fmt.Sprintf("--username=%s", user)

	cmd := CreateImportCommand(binaryPath, model.PostgreSQL, user,database, path)
	fmt.Println(cmd)
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	err := cmd.Run()

	if err != nil {
		return err
	}

	return nil
}
