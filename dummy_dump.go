package dummy_dump

import (
	"github.com/sadihakan/dummy-dump/errors"
	"github.com/sadihakan/dummy-dump/internal"
	"github.com/sadihakan/dummy-dump/model"
	"github.com/sadihakan/dummy-dump/util"
)

type DummyDump struct {
	c     *model.Config
	dump  internal.Dump
	Error error
}

func New(config *model.Config) (*DummyDump, error) {
	dd := new(DummyDump)
	dd.c = config

	if err := dd.configParser(); err != nil {
		return nil, err
	}

	return dd, nil
}

func (dd *DummyDump) configParser() (err error) {
	if dd.c.Source == "" {
		return errors.New(model.CONFIG_SOURCE_NIL)
	}

	if dd.c.User == "" {
		return errors.New(model.CONFIG_USER_NIL)
	}

	if dd.c.Path == "" {
		return errors.New(model.CONFIG_USER_NIL)
	}

	if dd.c.DB == "" {
		return errors.New(model.CONFIG_DB_NOT_EXIST)
	}

	if dd.c.BinaryPath == "" {
		return errors.New(model.CONFIG_BINARY_PATH_NOT_EXIST)
	}

	if dd.c.Import == dd.c.Export {
		return errors.New(model.CONFIG_METHOD_ERROR)
	}

	switch dd.c.Source {
	case model.PostgreSQL:
		dd.dump = internal.Postgres{}
		break
	case model.MySQL:
		//dd.dump = internal.MySQL{}
		break
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

	err := dd.dump.Import(config.BinaryPath, config.User, config.Path)

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
