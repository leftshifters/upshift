package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_readConfig(t *testing.T) {
	var c Config
	_ = c.readConfig("machine.toml")
	assert.Equal(t, "testPassword", c.machine.Password)
	assert.Equal(t, "testDeveloperAccount", c.machine.IOSDeveloperAccount)
	assert.Equal(t, "testCertificatePath", c.machine.IOSCertificatePath)
	assert.Equal(t, 5.0, float64(c.machine.AndroidSDKUpdateTime))

	_ = c.readConfig("repo.toml")
	assert.Equal(t, "testGitURL", c.repo.URL)
	assert.Equal(t, "testGitRemote", c.repo.Remote)
	assert.Equal(t, true, c.repo.CleanBeforeBuild)
	assert.Equal(t, true, c.repo.UninstallOlderBuilds)
	assert.Equal(t, "testDeveloperAccount", c.repo.IOSDeveloperAccount)
	assert.Equal(t, "testCertificatePath", c.repo.IOSCertificatePath)
	assert.Equal(t, "testProjectName", c.repo.IOSProjectName)
	assert.Equal(t, true, c.repo.IOSUseWorkspace)
	assert.Equal(t, "testScheme", c.repo.IOSScheme)
	assert.Equal(t, "testDevice", c.repo.IOSTestDevice)
	assert.Equal(t, "7.3.1", c.repo.IOSXcodeVersion)
	assert.Equal(t, "testPackage", c.repo.AndroidPackageName)
	assert.Equal(t, "testActivity", c.repo.AndroidMainActivityName)

	err := c.readConfig("fileDoesNotExist")
	assert.Equal(t, "File does not exist "+os.Getenv("GOPATH")+"/src/upshift/config/fileDoesNotExist", err.Error())

	err = c.readConfig("invalid.toml")
	assert.Contains(t, err.Error(), "Damn, we couldn't understand your TOML file")
}

func Test_ReadRepoConfig(t *testing.T) {
	var c Config
	_ = c.ReadRepoConfig()
	assert.Equal(t, "testGitURL", c.repo.URL)
	assert.Equal(t, "testGitRemote", c.repo.Remote)
	assert.Equal(t, true, c.repo.CleanBeforeBuild)
	assert.Equal(t, true, c.repo.UninstallOlderBuilds)
	assert.Equal(t, "testDeveloperAccount", c.repo.IOSDeveloperAccount)
	assert.Equal(t, "testCertificatePath", c.repo.IOSCertificatePath)
	assert.Equal(t, "testProjectName", c.repo.IOSProjectName)
	assert.Equal(t, true, c.repo.IOSUseWorkspace)
	assert.Equal(t, "testScheme", c.repo.IOSScheme)
	assert.Equal(t, "testDevice", c.repo.IOSTestDevice)
	assert.Equal(t, "7.3.1", c.repo.IOSXcodeVersion)
	assert.Equal(t, "testPackage", c.repo.AndroidPackageName)
	assert.Equal(t, "testActivity", c.repo.AndroidMainActivityName)
}

func Test_ReadMachineConfig(t *testing.T) {
	var c Config
	_ = c.ReadMachineConfig()
	assert.Equal(t, os.Getenv("HOME")+"/code/ios-distribution-certificates", c.machine.IOSCertificatePath)
}

func Test_WriteMachineConfig(t *testing.T) {
	var c Config
	_ = c.ReadMachineConfig()

	// Change android sdk update time to 100
	c.machine.AndroidSDKUpdateTime = 100
	_ = c.WriteMachineConfig()

	// Read config again to check
	_ = c.ReadMachineConfig()
	assert.Equal(t, 100.0, float64(c.machine.AndroidSDKUpdateTime))
}

func Test_PrepareSettings(t *testing.T) {
	var c Config
	_ = c.PrepareSettings()

	assert.Equal(t, "testDeveloperAccount", c.settings.IOSDeveloperAccount)
	assert.Equal(t, "iPhone 6", c.settings.IOSTestDevice)
}
