package config

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
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

func loadConfig() Config {

	var conf Config

	// See if a TOML file is available in this folder
	if _, err := os.Stat("sample.toml"); err == nil {
		// path/to/whatever exists
		log.Println("File exits")

		tomlBytes, err := ioutil.ReadFile("sample.toml")
		if err != nil {
			log.Println("Unable to read sample.toml")
		}

		tomlData := string(tomlBytes)

		if _, err := toml.Decode(tomlData, &conf); err != nil {
			log.Println(err.Error())
		}

	} else {
		log.Println("File does not exist")
	}

	return conf
}
