package config

import (
	"github.com/sadihakan/dummy-dump/util"
)

var sourceTypes = []string{
	string(MySQL),
	string(PostgreSQL),
	string(MSSQL),
}

type SourceType string

const (
	MySQL      SourceType = "mysql"
	PostgreSQL SourceType = "postgres"
	MSSQL      SourceType = "mssql"
)

func (s SourceType) IsValid() bool {
	e, _ := util.InArray(string(s), sourceTypes)
	return e
}
