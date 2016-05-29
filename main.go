package main

import (
	"github.com/progrium/go-basher"
	"log"
	"os"
	"path/filepath"
	"upshift/config"
	// lc "upshift/log-colours"
	// "upshift/utils"
)

// Main Function
func main() {

	// Load Config
	fileFullPath, err := filepath.Abs("./config/sample.toml")
	if err != nil {
		log.Println(err)
	}

	_, err = config.Load(fileFullPath)
	if err != nil {
		log.Println(err)
	}

	log.Println(os.Args)

	var job, action, flavour string

	if len(os.Args) > 1 {
		job = os.Args[1]
	}

	if len(os.Args) > 2 {
		action = os.Args[2]
	}

	if len(os.Args) > 3 {
		action = os.Args[3]
	}

	findTask(job, action, flavour)

	os.Exit(0)
}

func findTask(job string, action string, flavour string) {
	switch job {
	case "ios":
		switch action {
		case "build":
		case "run":
		case "deploy":
		default:
		}
	case "android":
		switch action {
		case "build":
		case "run":
		case "deploy":
		default:
		}
	case "setup":
		switch action {
		case "clone":
		case "config":
		default:
		}
	case "install":
	case "-v":
	case "action":
		switch action {
		case "setupSsh":
		case "setupScript":
		case "setupGradle":
		case "setupPods":
		case "setupXcode":
		case "setupXcpretty":
		case "upgradeScript":
		case "gitPull":
		case "gitClone":
		case "gitSubmodules":
		case "iosSimulator":
		case "androidEmulator":
		case "iosBuild":
		case "androidBuild":
		case "iosRun":
		case "androidRun":
		case "iosDeploy":
		case "androidDeploy":
		}
	default:
	}
}

func loadBash() {
	bash, _ := basher.NewContext("/bin/bash", false)
	if bash.HandleFuncs(os.Args) {
		os.Exit(0)
	}

	bash.CopyEnv()
	bash.Source("main.bash", nil)

	// status, err := bash.Run("gradleTasks", []string{"123", "234"})
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
