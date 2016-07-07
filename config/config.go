package config

import (
	"bytes"
	"errors"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"upshift/utils"
)

type Config struct {
	machine  MachineConfig
	repo     RepoConfig
	settings Settings
}

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

type MachineConfig struct {
	Password             string
	IOSDeveloperAccount  string
	IOSCertificatePath   string
	AndroidSDKUpdateTime int32
}

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

func Get() *Config {
	once.Do(func() {
		conf = &Config{}
		conf.PrepareSettings()
	})
	return conf
}

func (c *Config) PrepareSettings() error {

	// Defaults that we use
	appVersion := "0.8.4"
	testDevice := "iPhone 6"
	xcodeVersion := "7.3.1"

	// Read the Repo config first
	err := c.ReadRepoConfig()
	if err != nil {
		return err
	}

	// Move repo to settings
	c.settings.URL = c.repo.URL
	c.settings.Remote = c.repo.Remote
	c.settings.CleanBeforeBuild = c.repo.CleanBeforeBuild
	c.settings.UninstallOlderBuilds = c.repo.UninstallOlderBuilds
	c.settings.IOSDeveloperAccount = c.repo.IOSDeveloperAccount
	c.settings.IOSCertificatePath = c.repo.IOSCertificatePath
	c.settings.IOSProjectName = c.repo.IOSProjectName
	c.settings.IOSUseWorkspace = c.repo.IOSUseWorkspace
	c.settings.IOSScheme = c.repo.IOSScheme
	c.settings.IOSTestDevice = c.repo.IOSTestDevice
	c.settings.IOSXcodeVersion = c.repo.IOSXcodeVersion
	c.settings.AndroidPackageName = c.repo.AndroidPackageName
	c.settings.AndroidMainActivityName = c.repo.AndroidMainActivityName

	// Then read the machin config
	err = c.ReadMachineConfig()
	if err != nil {
		return err
	}

	c.settings.Password = c.machine.Password
	c.settings.AndroidSDKUpdateTime = c.machine.AndroidSDKUpdateTime
	if c.settings.IOSDeveloperAccount == "" {
		c.settings.IOSDeveloperAccount = c.machine.IOSDeveloperAccount
	}
	if c.settings.IOSCertificatePath == "" {
		c.settings.IOSCertificatePath = c.machine.IOSCertificatePath
	}

	// Finally add defaults
	if c.settings.IOSTestDevice == "" || c.settings.IOSTestDevice == "testDevice" {
		c.settings.IOSTestDevice = testDevice
	}
	if c.settings.IOSXcodeVersion == "" {
		c.settings.IOSXcodeVersion = xcodeVersion
	}
	c.settings.AppVersion = appVersion

	return nil
}

func (c *Config) ReadRepoConfig() error {
	repoPath := "config.toml"

	err := c.readConfig(repoPath)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) ReadMachineConfig() error {
	machinePath := filepath.Join(os.Getenv("HOME"), ".upshift", "config")

	err := c.readConfig(machinePath)
	if err != nil {
		return err
	}

	return nil
}

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

func (c *Config) GetRootPassword() (string, error) {
	// Check if it is defined in the environment variable
	RootPassword := os.Getenv("RootPassword")
	if RootPassword != "" {
		return RootPassword, nil
	}

	// If not, check if config has it
	if c.settings.Password != "" {
		return c.settings.Password, nil
	}

	// Don't have it, throw an error
	return "", errors.New("We can't do this without the root password, you need to set it up either in your env or the machine config")
}
