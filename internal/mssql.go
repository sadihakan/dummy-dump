package internal

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/sadihakan/dummy-dump/config"
	"github.com/sadihakan/dummy-dump/util"
	"log"
	"time"

	"net/url"

	"os"
	"os/exec"
)

type MSSQL struct {
	Dump
}

func (ms MSSQL) NewDB(dump config.Config) (*sql.DB, error) {
	urlQuery := url.Values{}
	urlQuery.Add("app name", "Backup App")

	connURL := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(dump.User, dump.Password),
		Host:     fmt.Sprintf("%s:%d", "localhost", 1433),
		RawQuery: urlQuery.Encode(),
		Path:     "/SQLExpress",
	}
	db, err := sql.Open("sqlserver", url.QueryEscape(connURL.String()))
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
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

func (ms MSSQL) CheckPath(dump config.Config) error {
	return nil
}


func (ms MSSQL) Export(dump config.Config) error {
	db, err := ms.NewDB(dump)
	var location string

	if dump.Path == "." {
		today := time.Now().UTC().UnixNano()
		p := util.GetBackupDirectory()
		filename := fmt.Sprintf("%s/%d.backup", p, today)
		location = filename
	} else {
		location = dump.Path
	}

	exportQuery := fmt.Sprintf(`BACKUP DATABASE [%s] TO DISK = '%s'`,
		dump.DB,
		location)
	_, err = db.Exec(exportQuery)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}

func (ms MSSQL) Import(dump config.Config) error {
	db, err := ms.NewDB(dump)
	if err != nil {
		return err
	}
	importQuery := fmt.Sprintf(`RESTORE DATABASE [%s] FROM DISK = '%s'`,
		dump.DB,
		dump.Path)
	_, err = db.Exec(importQuery)
	if err != nil {
		return err
	}

	return nil
}
