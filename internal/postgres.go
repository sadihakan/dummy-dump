package internal

import (
	"bytes"
	"fmt"
	"github.com/sadihakan/dummy-dump/model"
	"os"
)

type Postgres struct {
	Dump
}

func (p Postgres) Check() error {
	cmd := CreateCheckBinaryCommand(model.PostgreSQL)
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

func (p Postgres) Export(config model.Config) error {
	var out, errBuf bytes.Buffer

	user := fmt.Sprintf("--username=%s", config.User)
	database := fmt.Sprintf("--dbname=%s", config.DB)

	cmd := CreateExportCommand(config.BinaryPath, model.PostgreSQL, user, database)
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	err := cmd.Run()

	if err != nil {
		return err
	}

	return nil
}

func (p Postgres) Import(config model.Config) error {
	var out, errBuf bytes.Buffer

	user := fmt.Sprintf("--username=%s", config.User)

	cmd := CreateImportCommand(config.BinaryPath, model.PostgreSQL, user,config.DB, config.DB)
	fmt.Println(cmd)
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	err := cmd.Run()

	if err != nil {
		return err
	}

	return nil
}
