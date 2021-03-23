package model

import (
	"github.com/sadihakan/DummyDump/util"
)

var sourceTypes = []string{
	string(MySQL),
	string(PostgreSQL),
}

type SOURCE_TYPE string

const (
	MySQL SOURCE_TYPE = "mysql"
	PostgreSQL SOURCE_TYPE = "postgres"
)

func (s SOURCE_TYPE) IsValid() bool {
	e, _ := util.InArray(string(s), sourceTypes)
	return e
}