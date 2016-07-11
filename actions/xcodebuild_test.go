package actions

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Xcodebuild_LoadSettings(t *testing.T) {
	assert.Equal(t, 1, 1)

	// Get the current working directory
	currentWD, _ := os.Getwd()

	// Move to the new directory
	os.Chdir(filepath.Join(os.Getenv("HOME"), "code", "deezeno-ios"))
	newWD, _ := os.Getwd()
	t.Log(newWD)

	var xcodebuild Xcodebuild
	xcodebuild.LoadSettings()
	t.Log(xcodebuild)

	// Move back to older working directory
	os.Chdir(currentWD)
}
