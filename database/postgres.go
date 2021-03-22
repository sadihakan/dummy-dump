package database

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	//"syscall"
	"time"
)

const (
	PgRestore      = "pg_restore"
	PgDump         = "pg_dump"
	PgDatabase     = "--dbname=postgres"
	PgFlagFileName = "-f"
	PgFlagCreate   = "--create"
	PgFlatFormat   = "--format=c"
)



type Postgres struct{
	Dump
}

func (p Postgres) Check() error {
	cmd := exec.Command("postgres", "-V")
	var out, errBuf bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	err := cmd.Run()
	if err != nil {
		_, _ = os.Stderr.WriteString(err.Error())
		return err
	}
	return nil
}

func (p Postgres) Export(binaryPath string, user string, database string) error {
	today := time.Now().UTC().UnixNano()
	var out, errBuf bytes.Buffer

	user = fmt.Sprintf("--username=%s", user)
	database = fmt.Sprintf("--dbname=%s", database)
	filename := fmt.Sprintf("%d.backup", today)

	if binaryPath == "" {
		binaryPath = PgDump
	}

	cmd := exec.Command(binaryPath, user, database, PgFlagFileName, filename, PgFlagCreate, PgFlatFormat)
	cmd.Stdin = strings.NewReader("password")
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	err := cmd.Run()

	if err != nil {
		return err
	}

	return nil
}

func (p Postgres) Import(binaryPath string, user string, path string) error {
	var out, errBuf bytes.Buffer

	user = fmt.Sprintf("--username=%s", user)

	if binaryPath == "" {
		binaryPath = PgDump
	}

	cmd := exec.Command(PgRestore, user, "-W", PgDatabase, path, PgFlagCreate)
	cmd.Stdin = strings.NewReader("password")
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	err := cmd.Run()

	if err != nil {
		fmt.Println(out.String(), errBuf.String())
		return err
	}

	return nil
}
