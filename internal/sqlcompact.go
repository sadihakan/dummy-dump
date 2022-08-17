package internal

import (
	"context"
	"github.com/sadihakan/dummy-dump/config"
	"os"
	"path/filepath"
)

type SqlCompact struct {
	Dump
}

func (s SqlCompact) Check(_ context.Context) error {
	//
	return nil
}

func (s SqlCompact) CheckPath(_ context.Context, dump config.Config) error {
	//
	return nil
}

func (s SqlCompact) Export(_ context.Context, dump config.Config) error {
	_, err := os.Stat(filepath.Join(dump.BackupFilePath, dump.BackupName))
	return err
}

func (s SqlCompact) Import(_ context.Context, dump config.Config) error {
	//
	return nil
}
