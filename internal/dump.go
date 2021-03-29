package internal

import "github.com/sadihakan/dummy-dump/config"

// Dump ...
type Dump interface {
	Check() error
	Export(config config.Config) error
	Import(config config.Config) error
}
