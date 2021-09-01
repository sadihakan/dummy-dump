package internal

import (
	"bytes"
	"errors"
	"github.com/sadihakan/dummy-dump/config"
	"os"
)

type Oracle struct {
	Dump
}

func (o Oracle) Check() error {
	cmd := CreateCheckBinaryCommand(config.Oracle)
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

func (o Oracle) CheckPath(dump config.Config) error {
	cmd := CreateCheckBinaryPathCommand(dump)
	var out, errBuf bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	err := cmd.Run()
	if err != nil {
		return errors.New("oracle path does not located")
	}
	return nil
}

func (o Oracle) Export(dump config.Config) error {
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

func (o Oracle) Import(dump config.Config) error {
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
