package tasks

import (
	"log"
	"os"
	// "upshift/actions/android"
	// "upshift/actions/ios"
	"upshift/actions/setup"
	"upshift/config"
)

type taskList struct {
	actions []string
}

func init() {

}

func Setup(conf config.Config) {

	var job, action, flavour string

	if len(os.Args) > 1 {
		job = os.Args[1]
	}

	if len(os.Args) > 2 {
		action = os.Args[2]
	}

	if len(os.Args) > 3 {
		flavour = os.Args[3]
	}

	tasks := findTask(job, action, flavour)
	// log.Println(tasks)

	for i, action := range tasks.actions {
		loadTask(i+1, len(tasks.actions), action)
	}
}

func findTask(job string, action string, flavour string) taskList {
	switch job {
	case "ios":
		switch action {
		case "build":
			return taskList{actions: []string{"upgradeScript", "setupXcode", "setupXcpretty", "gitPull", "gitSubmodules", "setupPods", "iosBuild"}}
		case "run":
			return taskList{actions: []string{"upgradeScript", "setupXcode", "setupXcpretty", "gitPull", "gitSubmodules", "setupPods", "iosSimulator", "iosRun"}}
		case "deploy":
			return taskList{actions: []string{}}
		default:
			return taskList{actions: []string{}}
		}
	case "android":
		switch action {
		case "build":
			return taskList{actions: []string{"upgradeScript", "gitPull", "gitSubmodules", "setupGradle", "androidBuild"}}
		case "run":
			return taskList{actions: []string{"upgradeScript", "gitPull", "gitSubmodules", "setupGradle", "androidEmulator", "androidRun"}}
		case "deploy":
			return taskList{actions: []string{}}
		default:
			return taskList{actions: []string{}}
		}
	case "setup":
		switch action {
		case "clone":
			return taskList{actions: []string{"gitClone"}}
		case "config":
			return taskList{actions: []string{"setupConfig"}}
		default:
			return taskList{actions: []string{}}
		}
	case "install":
		return taskList{actions: []string{"setupScript"}}
	case "-v":
		return taskList{actions: []string{"showVersion"}}
	case "action":
		switch action {
		case "setupSsh":
			return taskList{actions: []string{"setupSsh"}}
		case "setupScript":
			return taskList{actions: []string{"setupScript"}}
		case "setupGradle":
			return taskList{actions: []string{"setupGradle"}}
		case "setupPods":
			return taskList{actions: []string{"setupPods"}}
		case "setupXcode":
			return taskList{actions: []string{"setupXcode"}}
		case "setupXcpretty":
			return taskList{actions: []string{"setupXcpretty"}}
		case "upgradeScript":
			return taskList{actions: []string{"upgradeScript"}}
		case "gitPull":
			return taskList{actions: []string{"gitPull"}}
		case "gitClone":
			return taskList{actions: []string{"gitClone"}}
		case "gitSubmodules":
			return taskList{actions: []string{"gitSubmodules"}}
		case "iosSimulator":
			return taskList{actions: []string{"iosSimulator"}}
		case "androidEmulator":
			return taskList{actions: []string{"androidEmulator"}}
		case "iosBuild":
			return taskList{actions: []string{"iosBuild"}}
		case "androidBuild":
			return taskList{actions: []string{"androidBuild"}}
		case "iosRun":
			return taskList{actions: []string{"iosRun"}}
		case "androidRun":
			return taskList{actions: []string{"androidRun"}}
		case "iosDeploy":
			return taskList{actions: []string{"iosDeploy"}}
		case "androidDeploy":
			return taskList{actions: []string{"androidDeploy"}}
		}
	default:
		return taskList{actions: []string{}}
	}
	return taskList{actions: []string{}}
}

func loadTask(count int, total int, task string) bool {
	log.Println(count, "/", total, task)
	switch task {
	case "upgradeScript":
		setup.UpgradeScript()
	case "showVersion":
	case "setupXcode":
	case "setupXcpretty":
	case "setupPods":
	case "setupGradle":
	case "setupConfig":
		setup.SetupConfig()
	case "setupScript":
	case "setupSsh":
	case "gitPull":
	case "gitSubmodules":
	case "gitClone":
	case "iosBuild":
	case "iosRun":
	case "iosSimulator":
	case "iosDeploy":
	case "androidBuild":
	case "androidRun":
	case "androidEmulator":
	case "androidDeploy":
	default:
	}
	return true
}
