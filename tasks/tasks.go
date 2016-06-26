package tasks

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"upshift/actions/android"
	"upshift/actions/ios"
	"upshift/actions/setup"
	c "upshift/colours"
	// "upshift/config"
	"upshift/utils"
)

type taskList struct {
	actions []string
}

func init() {

}

func Setup() {

	var job, action string

	if len(os.Args) > 1 {
		job = os.Args[1]
	}

	if len(os.Args) > 2 {
		action = os.Args[2]
	}

	tasks := findTask(job, action)

	fmt.Print("\n")

	var status int
	var skipNext bool

	for i, action := range tasks.actions {
		fmt.Println(c.Blue + c.Bold + "🛢  Starting " + c.Underline + strings.ToUpper(action) + c.Default + c.Light + " " + strconv.Itoa(i+1) + "/" + strconv.Itoa((len(tasks.actions))) + c.Default)
		status, skipNext = loadTask(action)
		fmt.Print("\n")

		if skipNext == true {
			fmt.Println(c.Gray + "➡️  We are stopping because the last step failed with status (" + strconv.Itoa(status) + ")" + c.Default)
			break
		}
	}

	if status > 0 {
		os.Exit(status)
	}

	fmt.Println("🍻  " + c.Yellow + "All done!" + c.Default)
}

func findTask(job string, action string) taskList {
	switch job {
	case "ios", "iOS", "i":
		switch action {
		case "build":
			return taskList{actions: []string{"upgradeScript", "setupXcode", "setupXcpretty", "gitPull", "gitSubmodules", "setupSigh", "setupPods", "installPods", "iosBuild"}}
		case "run":
			return taskList{actions: []string{"upgradeScript", "setupXcode", "setupXcpretty", "gitPull", "gitSubmodules", "setupSigh", "setupPods", "installPods", "iosSimulator", "iosRun"}}
		case "deploy":
			return taskList{actions: []string{"showHelp"}}
		default:
			return taskList{actions: []string{"showHelp"}}
		}
	case "android", "Android", "a":
		switch action {
		// Dont' add Android upgrade or install here. Their update is broken and it keep redownloading packaging
		case "build":
			return taskList{actions: []string{"upgradeScript", "gitPull", "gitSubmodules", "setupGradleW", "androidBuild"}}
		case "run":
			return taskList{actions: []string{"upgradeScript", "gitPull", "gitSubmodules", "setupGradleW", "androidEmulator", "androidRun"}}
		case "deploy":
			return taskList{actions: []string{"showHelp"}}
		case "update":
			return taskList{actions: []string{"setupAndroid"}}
		default:
			return taskList{actions: []string{"showHelp"}}
		}
	case "setup", "Setup", "s":
		switch action {
		// case "clone":
		// 	return taskList{actions: []string{"gitClone"}}
		case "config":
			return taskList{actions: []string{"setupConfig"}}
		case "export.plist":
			return taskList{actions: []string{"setupExportPlist"}}
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
		case "setupGradleW":
			return taskList{actions: []string{"setupGradleW"}}
		case "setupPods":
			return taskList{actions: []string{"setupPods"}}
		case "setupSigh":
			return taskList{actions: []string{"setupSigh"}}
		case "setupXcode":
			return taskList{actions: []string{"setupXcode"}}
		case "setupXcpretty":
			return taskList{actions: []string{"setupXcpretty"}}
		case "setupExportPlist":
			return taskList{actions: []string{"setupExportPlist"}}
		case "setupProfiles":
			return taskList{actions: []string{"setupProfiles"}}
		case "upgradeScript":
			return taskList{actions: []string{"upgradeScript"}}
		case "upgradeAndroid":
			return taskList{actions: []string{"upgradeAndroid"}}
		case "gitPull":
			return taskList{actions: []string{"gitPull"}}
		// case "gitClone":
		// 	return taskList{actions: []string{"gitClone"}}
		case "gitSubmodules":
			return taskList{actions: []string{"gitSubmodules"}}
		case "installPods":
			return taskList{actions: []string{"installPods"}}
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

func loadTask(task string) (int, bool) {
	switch task {
	case "upgradeScript":
		return setup.UpgradeScript()
	case "upgradeAndroid":
		return android.UpgradeAndroid()
	case "showVersion":
		return setup.ShowVersion()
	case "showHelp":
		return setup.ShowHelp()
	case "setupXcode":
		return ios.SetupXcode()
	case "setupXcpretty":
		return setup.SetupXcpretty()
	case "setupPods":
		return setup.SetupPods(false)
	case "setupSigh":
		return setup.SetupSigh(false)
	case "setupGradleW":
		return setup.SetupGradleW()
	case "setupConfig":
		return setup.SetupConfig()
	case "setupProfiles":
		return ios.SetupProfiles()
	// case "setupScript":
	// case "setupSsh":
	case "setupExportPlist":
		return ios.SetupExportPlist()
	case "setupAndroid":
		return android.SetupAndroid()
	case "gitPull":
		return setup.GitPull()
	case "gitSubmodules":
		return setup.GitSubmodules()
	// SKIP case "gitClone":
	case "installPods":
		return setup.InstallPods()
	case "iosBuild":
		return ios.IosBuild()
	// case "iosRun":
	// case "iosSimulator":
	// case "iosDeploy":
	case "androidBuild":
		return android.AndroidBuild()
	// case "androidRun":
	// case "androidEmulator":
	// case "androidDeploy":
	default:
		utils.LogError("It's sad, but we don't know how to " + c.Underline + "handle this effing case" + c.Default + "\nYou should try upshift -v to find out what do we support")
		return 1, true
	}
	return 0, true
}
