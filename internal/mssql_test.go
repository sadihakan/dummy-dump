package internal

import (
	"testing"
)

func TestMSSQL_Check(t *testing.T) {
	var d Dump
	d=MSSQL{}
	d.Check()
}

func TestMSSQL_CheckWithError(t *testing.T) {
	var d Dump
	d=MSSQL{}
	err:=d.Check()
	if err != nil {
		t.Errorf("MSSQL Check() Error: %s", err.Error())
	}

}

func TestMSSQL_Import(t *testing.T) {
	var d Dump
	d=MSSQL{}
	var  binary string
	binary = CheckBinary(binary, MSSQL,true,false)
	d.Import(binary,"testuser","deneme","C:\\Program Files\\Microsoft SQL Server\\MSSQL15.SQLEXPRESS\\MSSQL\\Backup\\deneme.bak")
}

func TestMSSQL_ImportWithError(t *testing.T) {
	var d Dump
	d=MSSQL{}
	var  binary string
	binary = CheckBinary(binary, MSSQL,true,false)
	err:=d.Import(binary,"testuser","deneme","C:\\Program Files\\Microsoft SQL Server\\MSSQL15.SQLEXPRESS\\MSSQL\\Backup\\deneme.bak")
	if err != nil {
		t.Errorf("MSSQL Import() Error: %s", err.Error())
	}
}

func TestMSSQL_Export(t *testing.T) {
	var d Dump
	d=MSSQL{}
	var  binary string
	binary = CheckBinary(binary, MSSQL,true,false)
	d.Export(binary,"testuser","deneme")
}

func TestMSSQL_ExportWithError(t *testing.T) {
	var d Dump
	d=MSSQL{}
	var  binary string
	binary = CheckBinary(binary, MSSQL,true,false)
	err:=d.Export(binary,"testuser","deneme")
	if err != nil {
		t.Errorf("MSSQL Import() Error: %s", err.Error())
	}
}
