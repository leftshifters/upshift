package tasks

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/leftshifters/upshift/actions"
	"github.com/leftshifters/upshift/colours"
	"github.com/leftshifters/upshift/utils"
)

type taskList struct {
	actions []string
}

func init() {

}

// Setup : how the actions will be handled
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

	for i, action := range tasks.actions {
		fmt.Println(colours.Blue + colours.Bold + "🛢  Starting " + colours.Underline + strings.ToUpper(action) + colours.Default + colours.Light + " " + strconv.Itoa(i+1) + "/" + strconv.Itoa((len(tasks.actions))) + colours.Default)
		status = loadTask(action)
		fmt.Print("\n")

		if status > 0 {
			fmt.Println(colours.Gray + "➡️  We are stopping because the last step failed with status (" + strconv.Itoa(status) + ")" + colours.Default)
			break
		}
	}

	if status > 0 {
		os.Exit(status)
	}

	fmt.Println("🍻  " + colours.Yellow + "All done!" + colours.Default)
}

func findTask(job string, action string) taskList {
	switch job {
	case "ios", "iOS", "i":
		switch action {
		case "build":
			return taskList{actions: []string{"upgradeScript", "setupXcode", "setupXcpretty", "gitPull", "gitSubmodules", "setupPods", "podsInstall", "iosPrepare", "iosBuild"}}
		case "test":
			return taskList{actions: []string{"upgradeScript", "setupXcode", "setupXctool", "gitPull", "gitSubmodules", "setupPods", "podsInstall", "iosPrepare", "iosTest"}}
		case "run":
			return taskList{actions: []string{"upgradeScript", "setupXcode", "setupXcpretty", "gitPull", "gitSubmodules", "setupFastlane", "setupPods", "podsInstall", "iosPrepare", "iosBuild", "iosRun"}}
		case "deploy":
			return taskList{actions: []string{"upgradeScript", "setupXcode", "setupXcpretty", "gitPull", "gitSubmodules", "setupFastlane", "setupPods", "podsInstall", "iosPrepare", "iosCertificates", "iosProvisioning", "iosArchive", "iosExportIPA", "iosCreateApp", "iosUploadBuild"}}
		default:
			return taskList{actions: []string{"showHelp"}}
		}
	case "android", "Android", "a":
		switch action {
		// Dont' add Android upgrade or install here. Their update is broken and it keep redownloading packaging
		case "build":
			return taskList{actions: []string{"upgradeScript", "gitPull", "gitSubmodules", "setupGradleWrapper", "androidBuild"}}
		case "test":
			return taskList{actions: []string{"upgradeScript", "gitPull", "gitSubmodules", "setupGradleWrapper", "androidBuild"}}
		case "run":
			return taskList{actions: []string{"upgradeScript", "gitPull", "gitSubmodules", "setupGradleWrapper", "androidEmulator", "androidRun"}}
		case "deploy":
			return taskList{actions: []string{"showHelp"}}
		default:
			return taskList{actions: []string{"showHelp"}}
		}
	case "setup", "Setup", "s":
		switch action {
		case "config":
			return taskList{actions: []string{"setupConfig"}}
		default:
			return taskList{actions: []string{"showHelp"}}
		}
	case "-v", "--version", "-version":
		return taskList{actions: []string{"showVersion"}}
	case "action", "act", "do":
		allActions := []string{"androidBuild", "androidDeploy", "androidLoadEmulator", "androidRun", "androidTest", "androidUpgrade", "gitPull", "gitSubmodules", "iosArchive", "iosBuild", "iosCertificates", "iosCreateApp", "iosDeploy", "iosExportIPA", "iosPrepare", "iosProvisioning", "iosRun", "iosSimulator", "iosTest", "podsInstall", "setupConfig", "setupFastlane", "setupGradleWrapper", "setupPods", "setupXcode", "setupXcpretty", "setupXctool", "showHelp", "showVersion", "upgradeScript"}

		for _, item := range allActions {
			if item == action {
				return taskList{actions: []string{item}}
			}
		}
	}
	return taskList{actions: []string{"showHelp"}}
}

func loadTask(task string) int {
	switch task {
	case "androidBuild":
		return actions.AndroidBuild()
	case "androidDeploy":
		return actions.AndroidDeploy()
	case "androidLoadEmulator":
		return actions.AndroidLoadEmulator()
	case "androidRun":
		return actions.AndroidRun()
	case "androidTest":
		return actions.AndroidTest()
	case "androidUpgrade":
		return actions.AndroidUpgrade()
	case "gitPull":
		return actions.GitPull()
	case "gitSubmodules":
		return actions.GitSubmodules()
	case "iosArchive":
		return actions.IosArchive()
	case "iosBuild":
		return actions.IosBuild()
	case "iosCertificates":
		return actions.IosCertificates()
	case "iosCreateApp":
		return actions.IosCreateApp()
	case "iosDeploy":
		return actions.IosDeploy()
	case "iosExportIPA":
		return actions.IosExportIPA()
	case "iosPrepare":
		return actions.IosPrepare()
	case "iosProvisioning":
		return actions.IosProvisioning()
	case "iosRun":
		return actions.IosRun()
	case "iosSimulator":
		return actions.IosSimulator()
	case "iosTest":
		return actions.IosTest()
	case "podsInstall":
		return actions.PodsInstall()
	case "setupConfig":
		return actions.SetupConfig()
	case "setupFastlane":
		return actions.SetupFastlane()
	case "setupGradleWrapper":
		return actions.SetupGradleWrapper()
	case "setupPods":
		return actions.SetupPods()
	case "setupXcode":
		return actions.SetupXcode()
	case "setupXcpretty":
		return actions.SetupXcpretty()
	case "setupXctool":
		return actions.SetupXctool()
	case "showHelp":
		return actions.ShowHelp()
	case "showVersion":
		return actions.ShowVersion()
	case "upgradeScript":
		return actions.UpgradeScript()
	default:
		utils.LogError("It's sad, but we don't know how to " + colours.Underline + "handle this effing case" + colours.Default + "\nYou should try upshift -v to find out what do we support")
		return 1
	}
}
