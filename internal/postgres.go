package internal

import (
	"bytes"
	"errors"
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

func (p Postgres) CheckPath(dump config.Config) error {
	cmd := CreateCheckBinaryPathCommand(dump)
	var out, errBuf bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	err := cmd.Run()
	if err != nil {
		return errors.New("path does not located")
	}
	return nil
}

func (p Postgres) Export(dump config.Config) error {
	var out, errBuf bytes.Buffer

	cmd := CreateExportCommand(dump)
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

	cmd := CreateImportCommand(dump)
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	err := cmd.Run()

	if err != nil {
		return errors.New(errBuf.String())
	}

	return nil
}
