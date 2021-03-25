package internal

import (
	"bytes"
	"fmt"
	"github.com/sadihakan/DummyDump/model"
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
// or create another user besides root

// MySQL ...
<<<<<<< HEAD
type MySQL struct {
	Dump
}
=======
type MySQL struct{}
>>>>>>> 9dcf83a030589ec764d633be08cccab1e1c7e59e

func (m MySQL) Check() error {
	cmd := exec.Command("mysql", "--version")
	err := cmd.Run()
	if err != nil {
		_, _ = os.Stderr.WriteString(err.Error())
		return err
	}
	return nil
}

func (m MySQL) Export(binaryPath string, user string, database string) error {
	filename := fmt.Sprintf("%d.backup", time.Now().UTC().UnixNano())
	cmd := CreateExportCommand(binaryPath, model.MySQL, user, database)
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

func (m MySQL) Import(binaryPath string, user string, database string) error {
	cmd := CreateImportCommand(binaryPath, model.MySQL, user, database)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	return err
}
