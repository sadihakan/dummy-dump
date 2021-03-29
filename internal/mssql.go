package internal

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/sadihakan/dummy-dump/config"

	"github.com/sadihakan/dummy-dump/util"
	"net/url"

	"os"
	"os/exec"
	"time"
)

type MSSQL struct {
	Dump
}

func (ms MSSQL) NewDB(config config.Config) (*sql.DB, error) {
	urlQuery := url.Values{}
	urlQuery.Add("app name", "Backup App")

	connURL := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(config.User, config.Password),
		Host:     fmt.Sprintf("%s:%d", "localhost", 1433),
		RawQuery: urlQuery.Encode(),
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

func (ms MSSQL) Export(dump config.Config) error {
	cmd := CreateExportCommand(binaryPath, config.MSSQL, user, database)
	fmt.Println(cmd)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	return err
}

func (ms MSSQL) Import(config config.Config) error {
	db, err := ms.NewDB(config)
	if err != nil {
		return err
	}
	importQuery := fmt.Sprintf(`RESTORE DATABASE [%s] FROM DISK = '%s'`,
		config.DB,
		config.Path)
	_, err = db.Exec(importQuery)
	if err != nil {
		return err
	}

	return nil
}
