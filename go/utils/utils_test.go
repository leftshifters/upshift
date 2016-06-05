package utils

import (
	"path/filepath"
	"testing"
)

func TestIsDocker(t *testing.T) {
	isDocker, _ := IsDocker()
	if isDocker == true {
		t.Fail()
		t.Log("We are certainly not inside docker. Are we?")
	}
}

func TestReadIfFileExists(t *testing.T) {
	mainGoFile, err := filepath.Abs("../main.go")
	_, err = ReadIfFileExists(mainGoFile)
	if err != nil {
		t.Fail()
		t.Log("main.go should exist in the folder outside")
	}
}
