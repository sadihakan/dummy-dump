package config

import (
	"github.com/sadihakan/dummy-dump/util"
)

var sourceTypes = []string{
	string(MySQL),
	string(PostgreSQL),
	string(MSSQL),
	string(Sqlite),
	string(Sqlcompact),
}

type SourceType string

const (
	MySQL      SourceType = "mysql"
	PostgreSQL SourceType = "postgres"
	MSSQL      SourceType = "mssql"
	Sqlite     SourceType = "sqlite"
	Sqlcompact SourceType = "sqlcompact"
)

func (s SourceType) IsValid() bool {
	e, _ := util.InArray(string(s), sourceTypes)
	return e
}
