package config

import "github.com/sadihakan/dummy-dump/errors"

type Config struct {
	Source     SourceType
	Import     bool
	Export     bool
	User       string
	Path       string
	DB         string
	BinaryPath string
}

func (c Config) checkAll() error {
	if c.Source == "" {
		return errors.New(errors.ConfigSourceNil)
	}

	if c.Source == "" {
		return errors.New(errors.ConfigSourceNil)
	}

	if c.User == "" {
		return errors.New(errors.ConfigUserNil)
	}

	if c.Path == "" {
		return errors.New(errors.ConfigPathNotExist)
	}

	if c.DB == "" {
		return errors.New(errors.ConfigDbNotExist)
	}

	if c.Import == c.Export {
		return errors.New(errors.ConfigMethodError)
	}
	return nil
}

func (c Config) CheckConfigPostgreSQL() error {
	if err := c.checkAll(); err != nil {
		return  err
	}

	if c.BinaryPath == "" {
		return errors.New(errors.ConfigBinaryPathNotExist)
	}

	return nil
}

func (c Config) CheckConfigMySQL() error {
	if err := c.checkAll(); err != nil {
		return  err
	}

	if c.BinaryPath == "" {
		return errors.New(errors.ConfigBinaryPathNotExist)
	}

	return nil
}

func (c Config) CheckConfigMsSQL() error {
	if err := c.checkAll(); err != nil {
		return  err
	}
	return nil
}