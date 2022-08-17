package internal

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/microsoft/go-mssqldb"
	"github.com/sadihakan/dummy-dump/config"
	"github.com/sadihakan/dummy-dump/util"
	"net/url"
	"path/filepath"

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
		Host:     fmt.Sprintf("%s:%d", dump.Host, dump.Port),
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

func (ms MSSQL) Check(_ context.Context) error {
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

func (ms MSSQL) CheckPath(_ context.Context, dump config.Config) error {
	return nil
}

func (ms MSSQL) Export(ctx context.Context, dump config.Config) error {
	db, err := ms.NewDB(dump)
	if err != nil {
		return err
	}

	var location string

	if dump.BackupFilePath == "." || dump.BackupFilePath == "" || dump.BackupFilePath == " " {
		p := util.GetBackupDirectory()
		filename := fmt.Sprintf(`%s\%s`, p, dump.BackupName)
		location = filename

	} else {
		location = fmt.Sprintf(`%s\%s`, dump.BackupFilePath, dump.BackupName)
	}

	exportQuery := fmt.Sprintf(`BACKUP DATABASE [%s] TO DISK = '%s'`,
		dump.DB,
		location)
	_, err = db.ExecContext(ctx, exportQuery)
	if err != nil {
		return err
	}
	return nil
}

func (ms MSSQL) Import(ctx context.Context, dump config.Config) error {
	db, err := ms.NewDB(dump)
	if err != nil {
		return err
	}
	importQuery := fmt.Sprintf(`RESTORE DATABASE [%s] FROM DISK = '%s' WITH REPLACE`,
		dump.DB,
		filepath.Join(dump.BackupFilePath, dump.BackupName))
	_, err = db.ExecContext(ctx, importQuery)
	if err != nil {
		return err
	}
	return nil
}
