package util

import "testing"

func TestInArray(t *testing.T) {
	var sourceTypes = []string{
		"mysql",
		"postgres",
	}
	exists, index := InArray("postgres", sourceTypes)
	if !exists {
		t.Errorf("Error at InArray(): Elemenent doesnt exist in array.")
	}
	if index != 1 {
		t.Errorf("Error at InArray(): Elemenent index is wrong.")

	}
	//100 percent coverage
	nExist,nIndex:=InArray("non-existent element",sourceTypes)
	if nExist||nIndex!=-1{
		t.Errorf("Error at InArray(): Function doesnt work as expected")
	}
}
