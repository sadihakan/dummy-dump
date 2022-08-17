package internal

import (
	"bytes"
	"context"
	"errors"
	"github.com/sadihakan/dummy-dump/config"
	"os"
)

type Oracle struct {
	Dump
}

func (o Oracle) Check(ctx context.Context) error {
	cmd := CreateCheckBinaryCommand(ctx, config.Oracle)
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

func (o Oracle) CheckPath(ctx context.Context, dump config.Config) error {
	cmd := CreateCheckBinaryPathCommand(ctx, dump)
	var out, errBuf bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	err := cmd.Run()
	if err != nil {
		return errors.New("oracle path does not located")
	}
	return nil
}

func (o Oracle) Export(ctx context.Context, dump config.Config) error {
	var out, errBuf bytes.Buffer

	cmd := CreateExportCommand(ctx, dump)
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	err := cmd.Run()

	if err != nil {
		return errors.New(errBuf.String())
	}

	return nil
}

func (o Oracle) Import(ctx context.Context, dump config.Config) error {
	var out, errBuf bytes.Buffer

	cmd := CreateImportCommand(ctx, dump)
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	err := cmd.Run()

	if err != nil {
		return errors.New(errBuf.String())
	}

	return nil
}
