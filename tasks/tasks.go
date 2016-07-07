package tasks

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"upshift/actions"
	c "upshift/colours"
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

	for i, action := range tasks.actions {
		fmt.Println(c.Blue + c.Bold + "🛢  Starting " + c.Underline + strings.ToUpper(action) + c.Default + c.Light + " " + strconv.Itoa(i+1) + "/" + strconv.Itoa((len(tasks.actions))) + c.Default)
		status, _ = loadTask(action)
		fmt.Print("\n")

		if status > 0 {
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
			return taskList{actions: []string{"upgradeScript", "setupXcode", "setupXcpretty", "gitPull", "gitSubmodules", "setupPods", "installPods", "iosPrepare", "iosBuild"}}
		case "test":
			return taskList{actions: []string{"upgradeScript", "setupXcode", "setupXctool", "gitPull", "gitSubmodules", "setupPods", "installPods", "iosPrepare", "iosTest"}}
		case "run":
			return taskList{actions: []string{"upgradeScript", "setupXcode", "setupXcpretty", "gitPull", "gitSubmodules", "setupFastlane", "setupPods", "installPods", "iosPrepare", "iosBuild", "iosDeploySimulator"}}
		case "deploy":
			return taskList{actions: []string{"upgradeScript", "setupXcode", "setupXcpretty", "gitPull", "gitSubmodules", "setupFastlane", "setupPods", "installPods", "iosPrepare", "iosCertificates", "iosProvisioning", "iosArchive", "iosExportIPA", "iosCreateApp", "iosUploadBuild"}}
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
		case "setupFastlane":
			return taskList{actions: []string{"setupFastlane"}}
		case "setupXcode":
			return taskList{actions: []string{"setupXcode"}}
		case "setupXcpretty":
			return taskList{actions: []string{"setupXcpretty"}}
		case "setupXctool":
			return taskList{actions: []string{"setupXctool"}}
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
		case "iosPrepare":
			return taskList{actions: []string{"iosPrepare"}}
		case "iosTest":
			return taskList{actions: []string{"iosTest"}}
		case "iosDeploySimulator":
			return taskList{actions: []string{"iosDeploySimulator"}}
		case "iosProvisioning":
			return taskList{actions: []string{"iosProvisioning"}}
		case "iosCertificates":
			return taskList{actions: []string{"iosCertificates"}}
		case "iosArchive":
			return taskList{actions: []string{"iosArchive"}}
		case "iosExportIPA":
			return taskList{actions: []string{"iosExportIPA"}}
		case "iosCreateApp":
			return taskList{actions: []string{"iosCreateApp"}}
		case "iosUploadBuild":
			return taskList{actions: []string{"iosUploadBuild"}}
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
		return actions.UpgradeScript()
	case "upgradeAndroid":
		return actions.UpgradeAndroid()
	case "showVersion":
		return actions.ShowVersion()
	case "showHelp":
		return actions.ShowHelp()
	case "setupXcode":
		return actions.SetupXcode()
	case "setupXcpretty":
		return actions.SetupXcpretty()
	case "setupXctool":
		return actions.SetupXctool()
	case "setupPods":
		return actions.SetupPods(false)
	case "setupFastlane":
		return actions.SetupFastlane(false)
	case "setupGradleW":
		return actions.SetupGradleW()
	case "setupConfig":
		return actions.SetupConfig()
	case "setupProfiles":
		return actions.SetupProfiles()
	// case "setupScript":
	// case "setupSsh":
	case "setupExportPlist":
		return actions.SetupExportPlist()
	case "setupAndroid":
		return actions.SetupAndroid()
	case "gitPull":
		return actions.GitPull()
	case "gitSubmodules":
		return actions.GitSubmodules()
	// SKIP case "gitClone":
	case "installPods":
		return actions.InstallPods()
	case "iosBuild":
		return actions.IosBuild()
	case "iosPrepare":
		return actions.IosPrepare()
	case "iosTest":
		return actions.IosTest()
	case "iosDeploySimulator":
		return actions.IosDeploySimulator()
	case "iosProvisioning":
		return actions.IosProvisioning()
	case "iosCertificates":
		return actions.IosCertificates()
	case "iosArchive":
		return actions.IosArchive()
	case "iosExportIPA":
		return actions.IosExportIPA()
	case "iosCreateApp":
		return actions.IosCreateApp()
	case "iosUploadBuild":
		return actions.IosUploadBuild()
	// case "iosRun":
	// case "iosSimulator":
	// case "iosDeploy":
	case "androidBuild":
		return actions.AndroidBuild()
	// case "androidRun":
	// case "androidEmulator":
	// case "androidDeploy":
	default:
		utils.LogError("It's sad, but we don't know how to " + c.Underline + "handle this effing case" + c.Default + "\nYou should try upshift -v to find out what do we support")
		return 1, true
	}
	return 0, true
}
