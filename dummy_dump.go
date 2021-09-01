package dummy_dump

import (
	"fmt"
	"github.com/sadihakan/dummy-dump/config"
	"github.com/sadihakan/dummy-dump/errors"
	"github.com/sadihakan/dummy-dump/internal"
	"github.com/sadihakan/dummy-dump/util"
	"path/filepath"
)

// DummyDump ..
type DummyDump struct {
	c     *config.Config
	dump  internal.Dump
	Error error
}

// New ..
func New(cfg ...*config.Config) (*DummyDump, error) {
	dd := new(DummyDump)

	if len(cfg) > 0 {
		dd.c = cfg[0]

		if err := dd.configParser(); err != nil {
			return nil, err
		}

	} else {
		dd.c = new(config.Config)
	}

	return dd, nil
}

// SetUser ..
func (dd *DummyDump) SetUser(username, password string, sourceType config.SourceType) {
	dd.c.Source = sourceType
	dd.c.User = username
	dd.c.Password = password
	dd.configParserWithoutCheck()
}

// SetBinaryPath ..
func (dd *DummyDump) SetBinaryPath(binaryPath string, sourceType config.SourceType, importArg bool, exportArg bool) {
	dd.c.Source = sourceType
	dd.c.BinaryPath = binaryPath
	dd.c.Export = exportArg
	dd.c.Import = importArg
	dd.configParserWithoutCheck()
}

// SetBinaryPath ..
func (dd *DummyDump) SetBinaryConfig(source config.SourceType, importArg bool, exportArg bool) {
	dd.c.Source = source
	dd.c.Import = importArg
	dd.c.Export = exportArg
}

func (dd *DummyDump) configParserWithoutCheck() {
	switch dd.c.Source {
	case config.PostgreSQL:
		dd.dump = internal.Postgres{}
		break
	case config.MySQL:
		dd.dump = internal.MySQL{}
		break
	case config.MSSQL:
		dd.dump = internal.MSSQL{}
	case config.Oracle:
		dd.dump = internal.Oracle{}
	case config.Sqlite:
		dd.dump = internal.Sqlite{}
	case config.Sqlcompact:
		dd.dump = internal.SqlCompact{}
	default:
		panic(errors.New("not implemented"))
	}
}

func (dd *DummyDump) configParser() (err error) {
	switch dd.c.Source {
	case config.PostgreSQL:
		if err = dd.c.CheckConfigPostgreSQL(); err != nil {
			return err
		}
		dd.dump = internal.Postgres{}
		break
	case config.MySQL:
		if err = dd.c.CheckConfigMySQL(); err != nil {
			return err
		}
		dd.dump = internal.MySQL{}
		break
	case config.MSSQL:
		if err = dd.c.CheckConfigMsSQL(); err != nil {
			return err
		}
		dd.dump = internal.MSSQL{}
	case config.Oracle:
		if err = dd.c.CheckConfigOracle(); err != nil {
			return err
		}
		dd.dump = internal.Oracle{}
	case config.Sqlcompact:
		dd.dump = internal.SqlCompact{}
	case config.Sqlite:
		dd.dump = internal.Sqlite{}
	default:
		err = errors.New("not implemented")
	}

	return err
}

func (dd *DummyDump) Import() *DummyDump {
	dumpConfig := dd.c

	if !util.PathExists(dumpConfig.BackupFilePath) && dd.Error == nil {
		dd.Error = errors.New(errors.ConfigPathNotExist)
	}

	err := dd.dump.Import(*dumpConfig)

	if err != nil && dd.Error == nil {
		dd.Error = err
	}

	return dd
}

func (dd *DummyDump) Export() *DummyDump {
	dumpConfig := dd.c
	err := dd.dump.Export(*dumpConfig)

	if err != nil && dd.Error == nil {
		dd.Error = err
	}

	return dd
}

func (dd *DummyDump) Check() *DummyDump {
	dd.Error = dd.dump.Check()

	return dd
}

func (dd *DummyDump) CheckPath() *DummyDump {
	dumpConfig := dd.c

	dd.Error = dd.dump.CheckPath(*dumpConfig)

	return dd
}

func (dd *DummyDump) Run() (*DummyDump, error) {
	if dd.Error != nil {
		return dd, dd.Error
	}

	return dd, nil
}

func (dd *DummyDump) GetBinary() (binaryPath string, version string) {
	dumpConfig := dd.c
	binaryPath, err := internal.CheckBinary(dumpConfig.BinaryPath, dumpConfig.Source, dumpConfig.Import, dumpConfig.Export)
	version, err = internal.CheckVersion(binaryPath, dumpConfig.Source)

	if err != nil {
		dd.Error = err
	}

	return binaryPath, version
}

//temprorary
func (dd *DummyDump) RestoreMSSQLToDB(cfg config.Config) error {
	ms := internal.MSSQL{}
	db, err := ms.NewDB(cfg)
	if err != nil {
		return err
	}
	importQuery := fmt.Sprintf(`RESTORE DATABASE [%s] FROM DISK = '%s' WITH REPLACE`,
		cfg.DB,
		filepath.Join(cfg.BackupFilePath, cfg.BackupName))
	_, err = db.Exec(importQuery)
	if err != nil {
		return err
	}
	return nil
}
