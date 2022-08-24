package internal

import (
	"bytes"
	"context"
	"errors"
	"github.com/sadihakan/dummy-dump/config"
	"os"
	"path/filepath"
	"runtime"
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

func (m MySQL) Check(ctx context.Context) error {
	cmd := CreateCheckBinaryCommand(ctx, config.MySQL)
	err := cmd.Run()
	if err != nil {
		_, _ = os.Stderr.WriteString(err.Error())
		return err
	}
	return nil
}

func (m MySQL) CheckPath(ctx context.Context, dump config.Config) error {
	cmd := CreateCheckBinaryPathCommand(ctx, dump)
	var out, errBuf bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	err := cmd.Run()
	if err != nil {
		return errors.New("mysql path does not located")
	}
	return nil
}

func (m MySQL) Export(ctx context.Context, dump config.Config) error {
	if runtime.GOOS == "windows" {
		_, dump.BinaryPath = filepath.Split(dump.BinaryPath)
	}

	cmd := CreateExportCommand(ctx, dump)
	var outb, errBuf bytes.Buffer
	cmd.Stderr = &errBuf
	cmd.Stdin = os.Stdin
	cmd.Stdout = &outb
	if err := cmd.Run(); err != nil {
		return errors.New(errBuf.String())
	}
	return nil
}

func (m MySQL) Import(ctx context.Context, dump config.Config) error {
	cmd := CreateImportCommand(ctx, dump)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	return err
}
