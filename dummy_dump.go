package dummy_dump

import (
	"github.com/sadihakan/dummy-dump/errors"
	"github.com/sadihakan/dummy-dump/internal"
	"github.com/sadihakan/dummy-dump/model"
	"github.com/sadihakan/dummy-dump/util"
)

// DummyDump ..
type DummyDump struct {
	c     *model.Config
	dump  internal.Dump
	Error error
}

// New ..
func New(config *model.Config) (*DummyDump, error) {
	dd := new(DummyDump)
	dd.c = config

	if err := dd.configParser(); err != nil {
		return nil, err
	}

	return dd, nil
}

func (dd *DummyDump) configParser() (err error) {
	switch dd.c.Source {
	case model.PostgreSQL:
		if err = util.CheckConfigPostgreSQL(*dd.c); err != nil {
			return err
		}
		dd.dump = internal.Postgres{}
		break
	case model.MySQL:
		if err = util.CheckConfigMySQL(*dd.c); err != nil {
			return err
		}
		dd.dump = internal.MySQL{}
		break
	case model.MSSQL:
		if err = util.CheckConfigMsSQL(*dd.c); err != nil {
			return err
		}
		dd.dump=internal.MSSQL{}
	default:
		err = errors.New("not implemented")
	}

	return err
}

func (dd *DummyDump) Import() *DummyDump {
	config := dd.c

	if !util.PathExists(config.Path) {
		dd.Error = errors.New(model.CONFIG_PATH_NOT_EXIST)
	}

	err := dd.dump.Import(config.BinaryPath,config.User,config.DB,config.Path)

	if err != nil {
		dd.Error = err
	}

	return dd
}

func (dd *DummyDump) Export() *DummyDump {
	config := dd.c

	err := dd.dump.Export(config.BinaryPath, config.User, config.DB)

	if err != nil {
		dd.Error = err
	}

	return dd
}

func (dd *DummyDump) Check() *DummyDump {
	dd.Error = dd.dump.Check()

	return dd
}

func (dd *DummyDump) Run() (*DummyDump, error) {
	if dd.Error != nil {
		return dd, dd.Error
	}

	return dd, nil
}
