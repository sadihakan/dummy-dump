package internal

import (
	"context"
	"github.com/sadihakan/dummy-dump/config"
	"os"
	"path/filepath"
)

type Sqlite struct {
	Dump
}

func (s Sqlite) Check(_ context.Context) error {
	//
	return nil
}

func (s Sqlite) CheckPath(_ context.Context, dump config.Config) error {
	//
	return nil
}

func (s Sqlite) Export(_ context.Context, dump config.Config) error {
	_, err := os.Stat(filepath.Join(dump.BackupFilePath, dump.BackupName))
	return err
}

func (s Sqlite) Import(_ context.Context, dump config.Config) error {
	//
	return nil
}
