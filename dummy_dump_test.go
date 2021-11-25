package dummy_dump

import (
	"encoding/json"
	"github.com/sadihakan/dummy-dump/config"
	"testing"
)

func TestNew(t *testing.T) {
	dd, err := New(&config.Config{
		Source:         config.PostgreSQL,
		Import:         true,
		Export:         false,
		User:           "sadihakan",
		BackupFilePath: "/path",
		DB:             "db",
		BinaryPath:     "/binaryPath",
	})

	if err != nil {
		t.Fatal(err)
	}

	dd.Check().Import().Run()
}

func TestNew2(t *testing.T) {
	cfg := config.Config{}
	err := json.Unmarshal([]byte(`{"Source":"postgres","Import":false,"Export":true,"User":"testuser","Password":"123456","DB":"testdb","Host":"localhost","Port":5432,"Service":"","BackupFilePath":"/Users/transferchain-mobile1/backup-app/backups","BackupName":"testdb.backup","BinaryPath":"/Applications/Postgres.app/Contents/Versions/latest/bin/pg_dump"}`), &cfg)
	if err != nil {
		t.Error(err)
	}
	dd, err := New(&cfg)
	if err != nil {
		t.Error(err)
	}
	dd, err = dd.Check().Run()
	if err != nil {
		t.Error(err)
	}
}