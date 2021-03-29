package dummy_dump

import (
	"github.com/sadihakan/dummy-dump/config"
	"github.com/sadihakan/dummy-dump/internal"
	"testing"
)

func TestNew(t *testing.T) {
	dd, err := New(&config.Config{
		Source:     internal.PostgreSQL,
		Import:     true,
		Export:     false,
		User:       "sadihakan",
		Path:       "/path",
		DB:         "db",
		BinaryPath: "/binaryPath",
	})

	if err != nil {
		t.Fatal(err)
	}

	dd.Check().Import().Run()
}