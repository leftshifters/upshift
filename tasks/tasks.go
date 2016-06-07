package tasks

import (
	"fmt"
	"os"
	// "upshift/actions/android"
	"strconv"
	"strings"
	"upshift/actions/ios"
	"upshift/actions/setup"
	colours "upshift/colours"
	"upshift/config"
	"upshift/utils"
)

type taskList struct {
	actions []string
}

func init() {

}

func Setup(conf config.Config) {

	var job, action string

	if len(os.Args) > 1 {
		job = os.Args[1]
	}

	if len(os.Args) > 2 {
		action = os.Args[2]
	}

	tasks := findTask(job, action)

	for i, action := range tasks.actions {
		fmt.Println(colours.Blue.Format + colours.Bold.Format + "🛢  Starting " + colours.Underline.Format + strings.ToUpper(action) + colours.Default.Format + colours.Light.Format + " " + strconv.Itoa(i+1) + "/" + strconv.Itoa((len(tasks.actions))) + colours.Default.Format)
		loadTask(i+1, len(tasks.actions), action)
	}
}

func findTask(job string, action string) taskList {
	switch job {
	case "ios", "iOS", "i":
		switch action {
		case "build":
			return taskList{actions: []string{"upgradeScript", "setupXcode", "setupXcpretty", "gitPull", "gitSubmodules", "setupPods", "iosBuild"}}
		case "run":
			return taskList{actions: []string{"upgradeScript", "setupXcode", "setupXcpretty", "gitPull", "gitSubmodules", "setupPods", "iosSimulator", "iosRun"}}
		case "deploy":
			return taskList{actions: []string{"showHelp"}}
		default:
			return taskList{actions: []string{"showHelp"}}
		}
	case "android", "Android", "a":
		switch action {
		case "build":
			return taskList{actions: []string{"upgradeScript", "gitPull", "gitSubmodules", "setupGradle", "androidBuild"}}
		case "run":
			return taskList{actions: []string{"upgradeScript", "gitPull", "gitSubmodules", "setupGradle", "androidEmulator", "androidRun"}}
		case "deploy":
			return taskList{actions: []string{"showHelp"}}
		default:
			return taskList{actions: []string{"showHelp"}}
		}
	case "setup", "Setup", "s":
		switch action {
		case "clone":
			return taskList{actions: []string{"gitClone"}}
		case "config":
			return taskList{actions: []string{"setupConfig"}}
		default:
			return taskList{actions: []string{"showHelp"}}
		}
	case "install":
		return taskList{actions: []string{"setupScript"}}
	case "-v", "--version", "-version":
		return taskList{actions: []string{"showVersion"}}
	case "action", "act":
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
		return taskList{actions: []string{"showHelp"}}
	}
	return taskList{actions: []string{"showHelp"}}
}

func loadTask(count int, total int, task string) bool {
	switch task {
	case "upgradeScript":
		setup.UpgradeScript()
	// case "showVersion":
	case "showHelp":
		setup.ShowHelp()
	case "setupXcode":
		ios.SetupXcode()
	// case "setupXcpretty":
	// case "setupPods":
	// case "setupGradle":
	case "setupConfig":
		setup.SetupConfig()
	// case "setupScript":
	// case "setupSsh":
	// case "gitPull":
	// case "gitSubmodules":
	// case "gitClone":
	// case "iosBuild":
	// case "iosRun":
	// case "iosSimulator":
	// case "iosDeploy":
	// case "androidBuild":
	// case "androidRun":
	// case "androidEmulator":
	// case "androidDeploy":
	default:
		utils.LogError("It's sad, but we don't know how to " + colours.Underline.Format + "handle this effing case" + colours.Default.Format + "\nYou should try upshift -v to find out what do we support")
	}
	return true
}
