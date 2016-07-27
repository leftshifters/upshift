package actions

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Gradle_Version(t *testing.T) {
	var g Gradle
	os.Chdir(filepath.Join("..", "android-test"))
	g.Version()
	assert.NotEmpty(t, g.version)
}

func Test_Gradle_AddWrapper(t *testing.T) {
	var g Gradle
	os.Chdir(filepath.Join("..", "android-test"))
	g.Version()
	g.AddWrapper()
	assert.Equal(t, true, g.wrapperInstalled)

	os.Remove(".gradlew")
	g.AddWrapper()
	assert.Equal(t, true, g.wrapperInstalled)
}

func Test_Gradle_Clean(t *testing.T) {
	var g Gradle
	os.Chdir(filepath.Join("..", "android-test"))
	g.Version()
	g.AddWrapper()
	status, err := g.Clean("clean.log")
	assert.Nil(t, err)
	assert.Equal(t, 0, status)
}

func Test_Gradle_Lint(t *testing.T) {
	var g Gradle
	os.Chdir(filepath.Join("..", "android-test"))
	g.Version()
	g.AddWrapper()
	status, err := g.Lint("lint.log")
	assert.Nil(t, err)
	assert.Equal(t, 0, status)
}

func Test_Gradle_Uninstall(t *testing.T) {
	var g Gradle
	os.Chdir(filepath.Join("..", "android-test"))
	g.Version()
	g.AddWrapper()
	// #TODO : Launch a device before running uninstall, or check if devices available

	var emulator AndroidEmulator
	devices, _ := emulator.ConnectedDevices()

	if len(devices) > 0 {
		status, err := g.Uninstall("uninstall.log")
		assert.Nil(t, err)
		assert.Equal(t, 0, status)
	}
}

func Test_Gradle_InstallDebug(t *testing.T) {
	var g Gradle
	os.Chdir(filepath.Join("..", "android-test"))
	g.Version()
	g.AddWrapper()

	// #TODO : Launch a device before running uninstall, or check if devices available
	var emulator AndroidEmulator
	devices, _ := emulator.ConnectedDevices()

	if len(devices) > 0 {
		status, err := g.InstallDebug("install-debug.log")
		assert.Nil(t, err)
		assert.Equal(t, 0, status)
	}
}

func Test_Gradle_Assemble(t *testing.T) {
	var g Gradle
	os.Chdir(filepath.Join("..", "android-test"))
	g.Version()
	g.AddWrapper()
	status, err := g.Assemble("assemble.log")
	assert.Nil(t, err)
	assert.Equal(t, 0, status)
}
