package internal

import (
	"bytes"
	"fmt"
	"github.com/sadihakan/dummy-dump/config"
	"os"
)

type Postgres struct {
	Dump
}

func (p Postgres) Check() error {
	cmd := CreateCheckBinaryCommand(config.PostgreSQL)
	var out, errBuf bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	err := cmd.Run()
	if err != nil {
		_, _ = os.Stderr.WriteString(err.Error())
		return err
	}
	return nil
}

func (p Postgres) Export(config config.Config) error {
	var out, errBuf bytes.Buffer

	user := fmt.Sprintf("--username=%s", config.User)
	database := fmt.Sprintf("--dbname=%s", config.DB)

	cmd := CreateExportCommand(config.BinaryPath, config.Source, user, database)
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	err := cmd.Run()

	if err != nil {
		return err
	}

	return nil
}

func (p Postgres) Import(config config.Config) error {
	var out, errBuf bytes.Buffer

	user := fmt.Sprintf("--username=%s", config.User)

	cmd := CreateImportCommand(config.BinaryPath, config.Source, user,config.DB, config.Path)
	fmt.Println(cmd)
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	err := cmd.Run()

	if err != nil {
		return err
	}

	return nil
}
