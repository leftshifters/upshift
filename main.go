package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/progrium/go-basher"
)

// Define structs for Config
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

type iOSConfig struct {
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

// Read Config from a TOML File

// Main Function
func main() {
	bash, _ := basher.NewContext("/bin/bash", false)
	// bash.ExportFunc("reverse", reverse)
	if bash.HandleFuncs(os.Args) {
		os.Exit(0)
	}

	bash.CopyEnv()
	bash.Source("main.bash", nil)

	status, err := bash.Run("main", os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(status)
}
