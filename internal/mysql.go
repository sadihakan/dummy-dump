package internal

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

const (
	dbname       = "deneme"
	skippassword = "--skip-password" // use this when you add -p arg
)

//to be able to access mysql without sudo do this :GRANT ALL PRIVILEGES ON *.* TO 'root'@'localhost' WITH GRANT OPTION;

// MySQL ...
type MySQL struct{}

func (m MySQL) Check() error {
	cmd := exec.Command("mysql", "--version")
	err := cmd.Run()
	if err != nil {
		_, _ = os.Stderr.WriteString(err.Error())
		return err
	}
	return nil
}

func (m MySQL) Export(user, database string) error {
	filename := fmt.Sprintf("%d.backup", time.Now().UTC().UnixNano())
	cmd := exec.Command("mysqldump", "-u", user, database)
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
	if err != nil {
		return err
	}
	return err
}

func (m MySQL) Import(user, path string) error {
	cmd := exec.Command("mysql", "-u", user, dbname, "-e", "source "+path)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	return err
}
