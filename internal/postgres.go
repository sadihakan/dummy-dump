package internal

import (
	"bytes"
	"context"
	"errors"
	"github.com/sadihakan/dummy-dump/config"
	"os"
)

type Postgres struct {
	Dump
}

func (p Postgres) Check(ctx context.Context) error {
	cmd := CreateCheckBinaryCommand(ctx, config.PostgreSQL)
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

func (p Postgres) CheckPath(ctx context.Context, dump config.Config) error {
	cmd := CreateCheckBinaryPathCommand(ctx, dump)
	var out, errBuf bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	err := cmd.Run()
	if err != nil {
		return errors.New("psql path does not located")
	}
	return nil
}

func (p Postgres) Export(ctx context.Context, dump config.Config) error {
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

func (p Postgres) Import(ctx context.Context, dump config.Config) error {
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
