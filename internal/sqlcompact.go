package internal

import (
	"github.com/sadihakan/dummy-dump/config"
	"os"
	"path/filepath"
)

type SqlCompact struct {
	Dump
}

func (s SqlCompact) Check() error {
	//
	return nil
}

func (s SqlCompact) CheckPath(dump config.Config) error {
	//
	return nil
}

func (s SqlCompact) Export(dump config.Config) error {
	_, err := os.Stat(filepath.Join(dump.BackupFilePath, dump.BackupName))
	return err
}

func (s SqlCompact) Import(dump config.Config) error {
	//
	return nil
}
