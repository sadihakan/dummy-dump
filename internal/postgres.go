package internal

import (
	"bytes"
	"errors"
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
		return errors.New(errBuf.String())
	}
	return nil
}

func (p Postgres) Export(dump config.Config) error {
	var out, errBuf bytes.Buffer

	user := fmt.Sprintf("--username=%s", dump.User)
	database := fmt.Sprintf("--dbname=%s", dump.DB)

	cmd := CreateExportCommand(dump.BinaryPath, config.PostgreSQL, user, database, dump.Path, dump.BackupName)
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	err := cmd.Run()

	if err != nil {
		return errors.New(errBuf.String())
	}

	return nil
}

func (p Postgres) Import(dump config.Config) error {
	var out, errBuf bytes.Buffer

	user := fmt.Sprintf("--username=%s", dump.User)
	cmd := CreateImportCommand(dump.BinaryPath, config.PostgreSQL, user, dump.DB, dump.Path)
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	err := cmd.Run()

	if err != nil {
		return errors.New(errBuf.String())
	}

	return nil
}
