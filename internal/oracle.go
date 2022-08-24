package internal

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/sadihakan/dummy-dump/config"
)

type Oracle struct {
	Dump
}

func (o Oracle) CheckPath(ctx context.Context, dump config.Config) error {
	cmd := CheckBinaryPathCommand(ctx, dump)

	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return errors.New(fmt.Sprint(err) + ": " + stderr.String())
	}

	return nil
}

func (o Oracle) Export(ctx context.Context, dump config.Config) error {
	cmd := CreateExportCommand(ctx, dump)

	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return errors.New(fmt.Sprint(err) + ": " + stderr.String())
	}

	return nil
}

func (o Oracle) Import(ctx context.Context, dump config.Config) error {
	cmd := CreateImportCommand(ctx, dump)

	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return errors.New(fmt.Sprint(err) + ": " + stderr.String())
	}

	return nil
}
