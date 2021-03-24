package util

import "testing"

func TestPathExists(t *testing.T) {
	pExists:=PathExists("/home/")
	if !pExists{
		t.Errorf("Error at PathExists(): path doesnt exist")
	}
	pNotExists:=PathExists("/nonexistentpath")
	if pNotExists{
		t.Errorf("Error at PathExists(): path not supposed to exist. (Change test variable if it exists)")
	}

}