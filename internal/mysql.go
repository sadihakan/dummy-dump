package internal

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/sadihakan/dummy-dump/config"
	"io/ioutil"
	"os"
	"time"
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

func (p MySQL) CheckPath(dump config.Config) error {
	cmd := CreateCheckBinaryPathCommand(dump)
	var out, errBuf bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	err := cmd.Run()
	if err != nil {
		return errors.New("path does not located")
	}
	return nil
}

func (m MySQL) Export(dump config.Config) error {
	filename := fmt.Sprintf("%d.backup", time.Now().UTC().UnixNano())
	cmd := CreateExportCommand(dump)
	fmt.Println(cmd)
	var outb, errBuf bytes.Buffer
	cmd.Stderr = &errBuf
	cmd.Stdin = os.Stdin
	cmd.Stdout = &outb
	err := cmd.Run()
	if err != nil {
		return errors.New(errBuf.String())
	}

	err = ioutil.WriteFile(filename, outb.Bytes(), 0644)
	return err
}

func (m MySQL) Import(dump config.Config) error {
	cmd := CreateImportCommand(dump)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	return err
}
