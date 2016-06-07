package config

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"path/filepath"
	colours "upshift/colours"
	"upshift/utils"
)

var conf Config

// Define structs for Config
type Config struct {
	Application ApplicationConfig
	Runner      RunnerConfig
	Build       BuildConfig
	IOS         IOSConfig
	Android     AndroidConfig
}

type ApplicationConfig struct {
	Debug bool
}

type RunnerConfig struct {
	RootPassword string
}

type BuildConfig struct {
	GitRepoURL           string
	GitRepoBranch        string
	CleanBeforeBuild     bool
	UninstallOlderBuilds bool
}

type IOSConfig struct {
	ProjectName  string
	UseWorkspace bool
	Scheme       string
	TestDevice   string
	Xcode        string
}

type AndroidConfig struct {
	PackageName      string
	MainActivityName string
}

func init() {

}

func Get() (Config, error) {
	if &conf == nil {
		fmt.Println("Trying to load config.toml")
		conf, err := Load()
		if err != nil {
			fmt.Println(err.Error())
			return conf, err
		}
	}
	return conf, nil
}

func Load() (Config, error) {
	var conf Config

	fileFullPath, err := filepath.Abs("./config.toml")
	if err != nil {
		return conf, errors.New("Could not create absolute path" + err.Error())
	}

	conf, err = LoadFile(fileFullPath)
	if err != nil {
		return conf, err
	}

	return conf, nil
}

func LoadFile(fileName string) (Config, error) {

	var conf Config

	// See if a TOML file is available in this folder
	tomlBytes, err := utils.ReadIfFileExists(fileName)
	if err != nil {
		return conf, errors.New("You should have a config file at " + colours.Red.Format + fileName + colours.Default.Format + "\nIf you would like to create one, go run " + colours.Blue.Format + "upshift setup config" + colours.Default.Format)
	}

	tomlData := string(tomlBytes)

	_, err = toml.Decode(tomlData, &conf)
	if err != nil {
		return conf, errors.New("Crap, we couldn't really " + colours.Red.Format + "understand" + colours.Default.Format + " what was written in your toml file.\nYou should go and check it again")
	}

	// Add Default Values

	// iPhone="iPhone 6"
	if conf.IOS.TestDevice == "" {
		conf.IOS.TestDevice = "iPhone 6"
	}
	// xcodeVersion="7.3"
	if conf.IOS.Xcode == "" {
		conf.IOS.Xcode = "7.3.1"
	}

	return conf, nil
}
