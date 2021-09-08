package config

import (
	"github.com/sadihakan/dummy-dump/errors"
)

type Config struct {
	//source: mysql-postgres-mssql-oracle-sqlite-sqlcompact
	Source SourceType

	//methods
	Import bool
	Export bool

	//database configuration
	User     string
	Password string
	DB       string
	Host     string
	Port     int

	//For Oracle
	Service string

	//os configuration
	BackupFilePath string //path where to save or retrieve
	BackupName     string //name which to save or retrieve, don't forget to add .bak or .backup!
	BinaryPath     string //etc: usr/bin/pg_dump
}

func (c *Config) checkAll() error {
	if c.Source == "" {
		return errors.New(errors.ConfigSourceNil)
	}

	if c.User == "" {
		return errors.New(errors.ConfigUserNil)
	}

	if c.Import == true && c.BackupFilePath == "" {
		return errors.New(errors.ConfigPathNotExist)
	}

	if c.DB == "" {
		return errors.New(errors.ConfigDbNotExist)
	}

	if c.Import == c.Export {
		return errors.New(errors.ConfigMethodError)
	}

	if c.Export && c.BackupFilePath == "" {
		c.BackupFilePath = "."
	}

	return nil
}

func (c *Config) CheckConfigPostgreSQL() error {

	if err := c.checkAll(); err != nil {
		return err
	}

	if c.BinaryPath == "" {
		return errors.New(errors.ConfigBinaryPathNotExist)
	}

	return nil
}

func (c *Config) CheckConfigMySQL() error {
	if err := c.checkAll(); err != nil {
		return err
	}

	if c.BinaryPath == "" {
		return errors.New(errors.ConfigBinaryPathNotExist)
	}

	return nil
}

func (c *Config) CheckConfigMsSQL() error {
	if err := c.checkAll(); err != nil {
		return err
	}

	return nil
}

func (c *Config) CheckConfigOracle() error {
	if err := c.checkAll(); err != nil {
		return err
	}
	if c.BinaryPath == "" {
		return errors.New(errors.ConfigBinaryPathNotExist)
	}
	return nil
}
