package util

import (
	"github.com/sadihakan/dummy-dump/errors"
	"github.com/sadihakan/dummy-dump/model"
)

func CheckConfigPostgreSQL(c model.Config) error {
	if c.Source == "" {
		return errors.New(model.CONFIG_SOURCE_NIL)
	}

	if c.Source == "" {
		return errors.New(model.CONFIG_SOURCE_NIL)
	}

	if c.User == "" {
		return errors.New(model.CONFIG_USER_NIL)
	}

	if c.Path == "" {
		return errors.New(model.CONFIG_USER_NIL)
	}

	if c.DB == "" {
		return errors.New(model.CONFIG_DB_NOT_EXIST)
	}

	if c.BinaryPath == "" {
		return errors.New(model.CONFIG_BINARY_PATH_NOT_EXIST)
	}

	if c.Import == c.Export {
		return errors.New(model.CONFIG_METHOD_ERROR)
	}

	return nil
}

func CheckConfigMySQL(c model.Config) error {
	if c.Source == "" {
		return errors.New(model.CONFIG_SOURCE_NIL)
	}

	if c.Source == "" {
		return errors.New(model.CONFIG_SOURCE_NIL)
	}

	if c.User == "" {
		return errors.New(model.CONFIG_USER_NIL)
	}

	if c.Path == "" {
		return errors.New(model.CONFIG_USER_NIL)
	}

	if c.DB == "" {
		return errors.New(model.CONFIG_DB_NOT_EXIST)
	}

	if c.BinaryPath == "" {
		return errors.New(model.CONFIG_BINARY_PATH_NOT_EXIST)
	}

	if c.Import == c.Export {
		return errors.New(model.CONFIG_METHOD_ERROR)
	}

	return nil
}

func CheckConfigMsSQL(c model.Config) error {
	if c.Source == "" {
		return errors.New(model.CONFIG_SOURCE_NIL)
	}

	if c.Source == "" {
		return errors.New(model.CONFIG_SOURCE_NIL)
	}

	if c.User == "" {
		return errors.New(model.CONFIG_USER_NIL)
	}

	if c.Path == "" {
		return errors.New(model.CONFIG_USER_NIL)
	}

	if c.DB == "" {
		return errors.New(model.CONFIG_DB_NOT_EXIST)
	}

	if c.Import == c.Export {
		return errors.New(model.CONFIG_METHOD_ERROR)
	}

	return nil
}