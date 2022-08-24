package internal

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/sadihakan/dummy-dump/config"
)

const (
	dbname       = "deneme"
	skippassword = "--skip-password" // use this when you add -p arg
)

//to be able to access mysql without sudo do this :GRANT ALL PRIVILEGES ON *.* TO 'root'@'localhost' WITH GRANT OPTION;
// or create another user besides root

// MySQL ...

type MySQL struct {
	Dump
}

func (m MySQL) CheckPath(ctx context.Context, dump config.Config) error {
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

func (m MySQL) Export(ctx context.Context, dump config.Config) error {
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

func (m MySQL) Import(ctx context.Context, dump config.Config) error {
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
