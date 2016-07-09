package config

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"upshift/utils"

	"github.com/BurntSushi/toml"
)

// Config : Construct to store the app's main config
type Config struct {
	machine  MachineConfig
	repo     RepoConfig
	Settings Settings
}

// Settings : Consturct to get overall settings for the app
type Settings struct {
	Password                string
	IOSDeveloperAccount     string // Comes from both repo and machine configs
	IOSCertificatePath      string // Comes from both repo and machine configs
	AndroidSDKUpdateTime    int32
	URL                     string
	Remote                  string
	CleanBeforeBuild        bool
	UninstallOlderBuilds    bool
	IOSProjectName          string
	IOSUseWorkspace         bool
	IOSScheme               string
	IOSTestDevice           string
	IOSXcodeVersion         string
	AndroidPackageName      string
	AndroidMainActivityName string
	AppVersion              string
}

// MachineConfig : Construct to readh machine config
type MachineConfig struct {
	Password             string
	IOSDeveloperAccount  string
	IOSCertificatePath   string
	AndroidSDKUpdateTime int32
}

// RepoConfig : Construct to read the repo config
type RepoConfig struct {
	URL                     string
	Remote                  string
	CleanBeforeBuild        bool
	UninstallOlderBuilds    bool
	IOSDeveloperAccount     string
	IOSCertificatePath      string
	IOSProjectName          string
	IOSUseWorkspace         bool
	IOSScheme               string
	IOSTestDevice           string
	IOSXcodeVersion         string
	AndroidPackageName      string
	AndroidMainActivityName string
}

// Idiomatic singleton from here - http://marcio.io/2015/07/singleton-pattern-in-go/
var conf *Config
var once sync.Once

// Get : Get the singleon to the config
func Get() *Config {
	once.Do(func() {
		conf = &Config{}
		err := conf.PrepareSettings()
		if err != nil {
			utils.LogError("Could not PrepareSettings\n" + err.Error())
		}
	})
	return conf
}

// PrepareSettings : Read all settings and populate the settings block
func (c *Config) PrepareSettings() error {

	// Defaults that we use
	appVersion := "0.8.4"
	testDevice := "iPhone 6"
	xcodeVersion := "7.3.1"

	// Read the Repo config first
	err := c.ReadRepoConfig()
	if err != nil {
		utils.LogInfo("Your can setup your repo config at config.toml\n" + err.Error())
	}

	// Move repo to settings
	c.Settings.URL = c.repo.URL
	c.Settings.Remote = c.repo.Remote
	c.Settings.CleanBeforeBuild = c.repo.CleanBeforeBuild
	c.Settings.UninstallOlderBuilds = c.repo.UninstallOlderBuilds
	c.Settings.IOSDeveloperAccount = c.repo.IOSDeveloperAccount
	c.Settings.IOSCertificatePath = c.repo.IOSCertificatePath
	c.Settings.IOSProjectName = c.repo.IOSProjectName
	c.Settings.IOSUseWorkspace = c.repo.IOSUseWorkspace
	c.Settings.IOSScheme = c.repo.IOSScheme
	c.Settings.IOSTestDevice = c.repo.IOSTestDevice
	c.Settings.IOSXcodeVersion = c.repo.IOSXcodeVersion
	c.Settings.AndroidPackageName = c.repo.AndroidPackageName
	c.Settings.AndroidMainActivityName = c.repo.AndroidMainActivityName

	// Then read the machin config
	err = c.ReadMachineConfig()
	if err != nil {
		utils.LogInfo("Your can setup your machine config at ~/.upshift/config\n" + err.Error())
	}

	c.Settings.Password = c.machine.Password
	c.Settings.AndroidSDKUpdateTime = c.machine.AndroidSDKUpdateTime
	if c.Settings.IOSDeveloperAccount == "" {
		c.Settings.IOSDeveloperAccount = c.machine.IOSDeveloperAccount
	}
	if c.Settings.IOSCertificatePath == "" {
		c.Settings.IOSCertificatePath = c.machine.IOSCertificatePath
	}

	// Finally add defaults
	if c.Settings.IOSTestDevice == "" || c.Settings.IOSTestDevice == "testDevice" {
		c.Settings.IOSTestDevice = testDevice
	}
	if c.Settings.IOSXcodeVersion == "" {
		c.Settings.IOSXcodeVersion = xcodeVersion
	}
	c.Settings.AppVersion = appVersion

	return nil
}

// ReadRepoConfig : Read the config in the repo - config.toml
func (c *Config) ReadRepoConfig() error {
	repoPath := "config.toml"

	err := c.readConfig(repoPath)
	if err != nil {
		return err
	}

	return nil
}

// ReadMachineConfig : Read the machine config at $HOME/.upshift/config
func (c *Config) ReadMachineConfig() error {
	machinePath := filepath.Join(os.Getenv("HOME"), ".upshift", "config")

	err := c.readConfig(machinePath)
	if err != nil {
		return err
	}

	return nil
}

// WriteMachineConfig : write a config from memory to $HOME/.upshift/config
func (c *Config) WriteMachineConfig() error {
	// Make sure the folder exits
	machinePath := filepath.Join(os.Getenv("HOME"), ".upshift")
	err := os.MkdirAll(machinePath, os.ModePerm)
	if err != nil {
		return err
	}

	// Create the file path
	machinePath = filepath.Join(machinePath, "config")

	// Create a byte buffer
	var buffer bytes.Buffer
	e := toml.NewEncoder(&buffer)
	err = e.Encode(c.machine)
	if err != nil {
		return err
	}

	// Write to disk
	err = ioutil.WriteFile(machinePath, []byte(buffer.String()), 0644)
	if err != nil {
		return err
	}

	return nil
}

// WriteRepoConfig : Write the config from memory to config.toml
func (c *Config) WriteRepoConfig() error {
	// Create the file path
	path, err := filepath.Abs("config.toml")

	// Create a byte buffer
	var buffer bytes.Buffer
	e := toml.NewEncoder(&buffer)
	err = e.Encode(c.repo)
	if err != nil {
		return err
	}

	// Write to disk
	err = ioutil.WriteFile(path, []byte(buffer.String()), 0644)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) readConfig(file string) error {
	// Get the absolute path
	path, err := filepath.Abs(file)
	if err != nil {
		return errors.New("Could not create absolute path\n" + err.Error())
	}

	// Check if file exists
	bytes, err := utils.FileRead(path)
	if err != nil {
		return errors.New(err.Error())
	}

	// Convert data to string
	data := string(bytes)

	// Read TOML data
	_, err = toml.Decode(data, &c.machine)
	_, err = toml.Decode(data, &c.repo)

	if err != nil {
		return errors.New("Damn, we couldn't understand your TOML file\n" + err.Error())
	}

	return nil
}

// GetRootPassword : Get the root password setup for the system
func (c *Config) GetRootPassword() (string, error) {
	// Check if it is defined in the environment variable
	RootPassword := os.Getenv("RootPassword")

	if RootPassword != "" {
		return RootPassword, nil
	}

	// If not, check if config has it
	if c.Settings.Password != "" {
		return c.Settings.Password, nil
	}

	// Don't have it, throw an error
	return "", errors.New("We can't do this without the root password, you need to set it up either in your env or the machine config")
}

// IsCI : Check if the currect script is running in a CI
func (c *Config) IsCI() bool {
	// Get GITLAB_CI from the environment
	isGitlab := os.Getenv("GITLAB_CI")
	if isGitlab == "true" {
		return true
	}
	return false
}

// IsDocker : Check if the current script is running in a docker container
func (c *Config) IsDocker() bool {
	// To check if it's docker or not, find out if /proc/1/cgroup has Docker written anywhere
	// We don't need to return an error on this, just a true of false
	cGroupFile := "/proc/1/cgroup"

	if utils.FileExists(cGroupFile) == false {
		// File not found, ceratinly not docker or a linux machine
		return false
	}

	// Read the file, and then check if the work docker is written inside it
	// We can read it directly because we know the file exits
	cGroupBytes, _ := utils.FileRead(cGroupFile)
	cGroupString := string(cGroupBytes)
	return strings.Contains(cGroupString, "docker")

}
