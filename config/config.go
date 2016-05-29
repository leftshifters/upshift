package config

import (
	"errors"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
)

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

func Load(fileName string) (Config, error) {

	var conf Config

	// See if a TOML file is available in this folder
	if _, err := os.Stat(fileName); err == nil {

		tomlBytes, err := ioutil.ReadFile(fileName)
		if err != nil {
			return conf, errors.New("We couldn't read the file you mentioned")
		}

		tomlData := string(tomlBytes)

		if _, err := toml.Decode(tomlData, &conf); err != nil {
			return conf, err
		}

	} else {
		return conf, errors.New("The config file does not exist")
	}

	return conf, nil
}
