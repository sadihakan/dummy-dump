package dummy_dump

import (
	"github.com/sadihakan/dummy-dump/config"
	"testing"
)

func TestNew(t *testing.T) {
	dd, err := New(&config.Config{
		Source:     config.PostgreSQL,
		Import:     true,
		Export:     false,
		User:       "sadihakan",
		BackupFilePath:       "/path",
		DB:         "db",
		BinaryPath: "/binaryPath",
	})

	if err != nil {
		t.Fatal(err)
	}

	dd.Check().Import().Run()
}