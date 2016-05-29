package utils

import (
	"testing"
)

func TestIsDocker(t *testing.T) {
	isDocker, _ := IsDocker()
	if isDocker == true {
		t.Fail()
		t.Log("We are certainly not inside docker. Are we?")
	}
}
