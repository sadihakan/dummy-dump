package internal

import (
	"bytes"
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

func (m MySQL) Export(dump config.Config) error {
	filename := fmt.Sprintf("%d.backup", time.Now().UTC().UnixNano())
	cmd := CreateExportCommand(dump.BinaryPath, config.MySQL, dump.User, dump.DB)
	var outb bytes.Buffer
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = &outb
	err := cmd.Run()
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(&outb)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename+".sql", b, 0644)
	return err
}

func (m MySQL) Import(dump config.Config) error {
	cmd := CreateImportCommand(dump.BinaryPath, config.MySQL, dump.User, dump.DB, dump.Path)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	return err
}
