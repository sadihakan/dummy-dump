package dummy_dump

import (
	"github.com/sadihakan/dummy-dump/model"
	"testing"
)

func TestNew(t *testing.T) {
	dd, err := New(&model.Config{
		Source:     model.PostgreSQL,
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