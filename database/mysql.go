package database

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

const (
	dbname = "deneme"
	skippassword="--skip-password"
)

type MySQL struct {
	Dump
}

func (m MySQL) Check() error {
	cmd := exec.Command("mysql", "--version")

	err := cmd.Run()
	if err != nil {
		_, _ = os.Stderr.WriteString(err.Error())
		return err
	}

	return err
}


func (m MySQL) Export(user ,database string) error{
	filename := fmt.Sprintf("%d.backup", time.Now().UTC().UnixNano())
	cmd:= exec.Command("sudo","mysqldump","-u",user,"-p",skippassword,database)
	var outb bytes.Buffer
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout=&outb

	fmt.Println()

	if err := cmd.Run(); err != nil {
		return err
	}
	bytes, err := ioutil.ReadAll(&outb)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename+".sql", bytes, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (m MySQL) Import(user,path string) error{
	//cmd:= exec.Command("bash", "-c","sudo mysql -u"+dbuser +" -p < "+dir)
	//cmd:= exec.Command("mysqlimport", "-uroot","-p",dbname,dir)
	cmd:=exec.Command("sudo","mysql", "-u", user, "-p",skippassword ,dbname,"-e", "source "+ path )
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}