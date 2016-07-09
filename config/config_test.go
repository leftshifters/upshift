package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_readConfig(t *testing.T) {
	c := Get()
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
	c := Get()
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
	assert.Equal(t, "8.0", c.repo.IOSXcodeVersion)
	assert.Equal(t, "testPackage", c.repo.AndroidPackageName)
	assert.Equal(t, "testActivity", c.repo.AndroidMainActivityName)
}

func Test_ReadMachineConfig(t *testing.T) {
	c := Get()
	_ = c.ReadMachineConfig()
	assert.Equal(t, os.Getenv("HOME")+"/code/ios-distribution-certificates", c.machine.IOSCertificatePath)
}

func Test_WriteMachineConfig(t *testing.T) {
	c := Get()
	_ = c.ReadMachineConfig()

	// Change android sdk update time to 100
	c.machine.AndroidSDKUpdateTime = 100
	_ = c.WriteMachineConfig()

	// Read config again to check
	_ = c.ReadMachineConfig()
	assert.Equal(t, 100.0, float64(c.machine.AndroidSDKUpdateTime))
}

func Test_PrepareSettings(t *testing.T) {
	c := Get()
	_ = c.PrepareSettings()

	assert.Equal(t, "testDeveloperAccount", c.Settings.IOSDeveloperAccount)
	assert.Equal(t, "iPhone 6", c.Settings.IOSTestDevice)
}

func Test_WriteRepoConfig(t *testing.T) {
	c := Get()
	_ = c.ReadRepoConfig()

	// Change XcodeVersion to 8.0
	c.repo.IOSXcodeVersion = "8.0"
	_ = c.WriteRepoConfig()
	// Reset local value
	c.repo.IOSXcodeVersion = "7.3.1"

	// Read config again to check
	_ = c.ReadRepoConfig()
	assert.Equal(t, "8.0", c.repo.IOSXcodeVersion)
}

func Test_Get(t *testing.T) {
	// Create multiple instances and see if you get the correct one
	c := Get()
	c.machine.Password = "fakeTestPassword"

	d := Get()
	assert.Equal(t, "fakeTestPassword", d.machine.Password)
}

func Test_GetRootPassword(t *testing.T) {
	// Get the old password from env
	oldPassword := os.Getenv("RootPassword")

	// Case #1 : If a password is set in env, do we read it

	// Set a new temporary password in env
	os.Setenv("RootPassword", "tempRootPassword")

	c := Get()
	password, _ := c.GetRootPassword()
	assert.Equal(t, "tempRootPassword", password)

	// Case #2 : If there is no password in env, do we read the password from machine config

	// Set an empty password in env
	os.Setenv("RootPassword", "")
	// Set a password in settings
	c.Settings.Password = "tempRootPassword2"

	// Check if you get the correct value
	password, _ = c.GetRootPassword()
	assert.Equal(t, "tempRootPassword2", password)

	// Case #3 : No password in env or machine config, should return error

	// Set an empty password in env
	os.Setenv("RootPassword", "")
	// Set a password in settings
	c.Settings.Password = ""

	// Check if you get the correct value
	_, err := c.GetRootPassword()
	assert.Contains(t, err.Error(), "without the root password")

	// Reset the old password
	os.Setenv("RootPassword", oldPassword)
}

func Test_IsCI(t *testing.T) {
	// Get original state of GITLAB_CI
	currentCI, exists := os.LookupEnv("GITLAB_CI")

	// Set the variable in env
	os.Setenv("GITLAB_CI", "true")

	// Check status
	c := Get()
	ci := c.IsCI()
	assert.Equal(t, true, ci)

	// Remove the variable in env
	os.Unsetenv("GITLAB_CI")

	// Check status
	ci = c.IsCI()
	assert.Equal(t, false, ci)

	// Reset
	if exists == true {
		os.Setenv("GITLAB_CI", currentCI)
	} else {
		os.Unsetenv("GITLAB_CI")
	}
}

func Test_IsDocker(t *testing.T) {
	// Case #1 : Trying on my mac
	c := Get()
	docker := c.IsDocker()
	assert.Equal(t, false, docker)

	// Case #2 : Not sure how to test this on docker
}
