package actions

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Xcodebuild_LoadSettings(t *testing.T) {
	// Get the current working directory
	currentWD, _ := os.Getwd()

	// Move to the new directory
	os.Chdir(filepath.Join(os.Getenv("HOME"), "code", "deezeno-ios"))

	var xcodebuild Xcodebuild

	err := xcodebuild.LoadSettings()
	assert.Equal(t, "project", xcodebuild.Type)
	assert.Nil(t, err)

	err = xcodebuild.FindSchemes()
	assert.Nil(t, err)
	assert.Equal(t, "deezeno-ios", xcodebuild.Scheme)

	err = xcodebuild.Build()
	assert.Nil(t, err)

	var simulator IOSSimulator
	simulator.StartSimulator("iPhone 6 (9.3)")

	err = xcodebuild.Run()
	assert.Nil(t, err)

	err = xcodebuild.InstallCertificates()
	assert.Nil(t, err)

	err = xcodebuild.Archive()
	assert.Nil(t, err)

	err = xcodebuild.SetupExportPlist()
	assert.Nil(t, err)

	err = xcodebuild.ExportIPA()
	assert.Nil(t, err)

	err = xcodebuild.IncrementBuildNumber()
	assert.Nil(t, err)

	// Move back to older working directory
	os.Chdir(currentWD)
}

func Test_Xcodebuild_SwitchXcode(t *testing.T) {
	// Get the current working directory
	currentWD, _ := os.Getwd()

	// Move to the new directory
	os.Chdir(filepath.Join(os.Getenv("HOME"), "code", "deezeno-ios"))

	var xcodebuild Xcodebuild

	xcodebuild.XcodeVersion = "7.3"
	err := xcodebuild.SwitchXcode()
	assert.Nil(t, err)

	xcodebuild.XcodeVersion = "7.3.1"
	err = xcodebuild.SwitchXcode()
	assert.Nil(t, err)

	// Move back to older working directory
	os.Chdir(currentWD)
}
