package internal

import (
	"context"
	"github.com/sadihakan/dummy-dump/config"
)

// Dump ...
type Dump interface {
	CheckPath(ctx context.Context, dump config.Config) error
	Export(ctx context.Context, dump config.Config) error
	Import(ctx context.Context, dump config.Config) error
}
