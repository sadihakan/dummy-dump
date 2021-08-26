package internal

import (
	"github.com/sadihakan/dummy-dump/config"
	"os"
	"path/filepath"
)

type Sqlite struct {
	Dump
}

func (s Sqlite) Check() error {
	//
	return nil
}

func (s Sqlite) CheckPath(dump config.Config) error {
	//
	return nil
}

func (s Sqlite) Export(dump config.Config) error {
	_, err := os.Stat(filepath.Join(dump.BackupFilePath, dump.BackupName))
	return err
}

func (s Sqlite) Import(dump config.Config) error {
	//
	return nil
}