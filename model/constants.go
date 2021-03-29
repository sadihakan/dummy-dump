package model

import (
	"github.com/sadihakan/dummy-dump/util"
)

var sourceTypes = []string{
	string(MySQL),
	string(PostgreSQL),
	string(MSSQL),
}

type SOURCE_TYPE string

const (
	MySQL      SOURCE_TYPE = "mysql"
	PostgreSQL SOURCE_TYPE = "postgres"
	MSSQL      SOURCE_TYPE = "mssql"
)

func (s SOURCE_TYPE) IsValid() bool {
	e, _ := util.InArray(string(s), sourceTypes)
	return e
}

type DummyDumpError string

const (
	CONFIG_USER_NIL              DummyDumpError = "user can not be nil"
	CONFIG_SOURCE_NIL            DummyDumpError = "select source"
	CONFIG_PATH_NOT_EXIST        DummyDumpError = "path is not exist"
	CONFIG_DB_NOT_EXIST          DummyDumpError = "DB can not be nil"
	CONFIG_BINARY_PATH_NOT_EXIST DummyDumpError = "binary path can not be nil"
	CONFIG_METHOD_ERROR          DummyDumpError = "select method"
)
