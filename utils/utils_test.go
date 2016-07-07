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

func Test_CreateList(t *testing.T) {
	testMessage := "On branch go\nYour branch is up-to-date with 'origin/go'.\nnothing to commit, working directory clean"

	// Create an array of all lines
	testArray := CreateList(testMessage, []string{})
	assert.Equal(t, 3, len(testArray))

	// Create an array of all lines except the ones with 'nothing to commit'
	testArray2 := CreateList(testMessage, []string{"nothing to commit"})
	assert.Equal(t, 2, len(testArray2))
}

func Test_FileTail(t *testing.T) {
	failTailTest, _ := filepath.Abs("failTailTest")
	data, err := FileTail(failTailTest, 5)
	assert.Contains(t, err.Error(), "File does not exist")

	tailTest, _ := filepath.Abs("tailTest")
	data, _ = FileTail(tailTest, 5)
	assert.Equal(t, "43210", data)

	data, _ = FileTail(tailTest, 15)
	assert.Equal(t, "9876543210", data)
}

func Test_FileRead(t *testing.T) {
	failTailTest, _ := filepath.Abs("failTailTest")
	data, err := FileRead(failTailTest)
	assert.Contains(t, err.Error(), "File does not exist")

	tailTest, _ := filepath.Abs("tailTest")
	data, _ = FileRead(tailTest)
	assert.Equal(t, "9876543210", data)
}
