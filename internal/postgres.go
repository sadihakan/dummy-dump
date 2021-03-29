package internal

import (
	"bytes"
	"fmt"
	"github.com/sadihakan/dummy-dump/config"
	"os"
)

type Postgres struct {
	Dump
}

func (p Postgres) Check() error {
	cmd := CreateCheckBinaryCommand(config.PostgreSQL)
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

func (p Postgres) Export(dump config.Config) error {
	var out, errBuf bytes.Buffer

	user := fmt.Sprintf("--username=%s", dump.User)
	database := fmt.Sprintf("--dbname=%s", dump.DB)

	cmd := CreateExportCommand(dump.BinaryPath, config.PostgreSQL, user, database)
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	err := cmd.Run()

	if err != nil {
		return err
	}

	return nil
}

func (p Postgres) Import(dump config.Config) error {
	var out, errBuf bytes.Buffer

	user := fmt.Sprintf("--username=%s", dump.User)

	cmd := CreateImportCommand(dump.BinaryPath, config.PostgreSQL, user, dump.DB, dump.Path)
	fmt.Println(cmd)
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	err := cmd.Run()

	if err != nil {
		return err
	}

	return nil
}
