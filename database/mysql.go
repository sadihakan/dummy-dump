package database

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

const dbname = "mysql"

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
	cmd:= exec.Command("sudo","mysqldump","-u",user,"-p",database)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	bytes, err := ioutil.ReadAll(stdout)
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
	cmd:=exec.Command("sudo","mysql", "-u", user, "-p", dbname,"-e", "source "+ path )
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}