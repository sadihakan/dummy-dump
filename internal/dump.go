package internal

import "github.com/sadihakan/dummy-dump/model"

// Dump ...
type Dump interface {
	Check() error
	Export(config model.Config) error
	Import(config model.Config) error
}
