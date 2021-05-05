package internal

import (
	"bytes"
	"errors"
	"github.com/sadihakan/dummy-dump/config"
	"os"
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

func (m MySQL) Check() error {
	cmd := CreateCheckBinaryCommand(config.MySQL)
	err := cmd.Run()
	if err != nil {
		_, _ = os.Stderr.WriteString(err.Error())
		return err
	}
	return nil
}

func (m MySQL) CheckPath(dump config.Config) error {
	cmd := CreateCheckBinaryPathCommand(dump)
	var out, errBuf bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	err := cmd.Run()
	if err != nil {
		return errors.New("mysql path does not located")
	}
	return nil
}

func (m MySQL) Export(dump config.Config) error {
	cmd:= CreateExportCommand(dump)
	var outb, errBuf bytes.Buffer
	cmd.Stderr = &errBuf
	cmd.Stdin = os.Stdin
	cmd.Stdout = &outb
	if err := cmd.Run(); err != nil {
		return errors.New(errBuf.String())
	}
	return nil
}

func (m MySQL) Import(dump config.Config) error {
	cmd := CreateImportCommand(dump)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	return err
}
