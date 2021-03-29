package internal

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/sadihakan/dummy-dump/model"
	"github.com/sadihakan/dummy-dump/util"
	"net/url"
	"os"
	"os/exec"
	"time"
)

type MSSQL struct {
	Dump
}

func (ms MSSQL) NewDB(config model.Config) (*sql.DB,error){
	urlQuery:=url.Values{}
	urlQuery.Add("app name","Backup App")

	connURL:= &url.URL{
		Scheme: "sqlserver",
		User: url.UserPassword(config.User,config.Password),
		Host: fmt.Sprintf("%s:%d","localhost",1433),
		RawQuery: urlQuery.Encode(),
	}
	db,err:=sql.Open("sqlserver",url.QueryEscape(connURL.String()))
	err=db.Ping()
	if err != nil {
		return nil,err
	}
	return db,nil
}


func (ms MSSQL) Check() error {
	cmd := exec.Command("reg", "query", "HKEY_LOCAL_MACHINE\\SOFTWARE\\Microsoft", "/f", "mssql", "/k")
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func (ms MSSQL) Export(config model.Config) error {
	db,err:=ms.NewDB(config)

	if err != nil {
		return err
	}
	today := time.Now().UTC().UnixNano()
	filename := fmt.Sprintf("%d.bak", today)
	exportQuery := fmt.Sprintf(`BACKUP DATABASE [%s] TO DISK = '%s' WITH STATS = 10`,
		config.DB,
		util.GetMSSQLBackupDirectory()+`\`+fmt.Sprintf("%d", filename))

	_,err=db.Exec(exportQuery)
	if err != nil {
		return err
	}
	return nil
}

func (ms MSSQL) Import(config model.Config) error {
	db,err:=ms.NewDB(config)
	if err != nil {
		return err
	}
	importQuery := fmt.Sprintf(`RESTORE DATABASE [%s] FROM DISK = '%s'`,
		config.DB,
		config.Path)
	_,err=db.Exec(importQuery)
	if err != nil {
		return err
	}

	return nil
}
