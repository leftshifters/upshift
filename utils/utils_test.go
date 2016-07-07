package utils

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func Test_LogMessage(t *testing.T) {
	LogMessage("LogMessage writes like this")
}

func Test_LogInfo(t *testing.T) {
	LogInfo("LogInfo writes like this")
}

func Test_LogError(t *testing.T) {
	LogError("LogError writes like this")
}

func Test_IsDocker(t *testing.T) {
	isDocker := IsDocker()
	assert.Equal(t, false, isDocker)
}

func Test_FileRead(t *testing.T) {
	mainGoFile, err := filepath.Abs("../main.go")
	_, err = FileRead(mainGoFile)
	if err != nil {
		t.Fail()
		t.Log("main.go should exist in the folder outside")
	}
}
