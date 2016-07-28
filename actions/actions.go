package actions

import (
	"fmt"

	"github.com/leftshifters/upshift/config"
	"github.com/leftshifters/upshift/utils"
)

// AndroidBuild : Build an android project
func AndroidBuild() int {
	var gradle Gradle
	var emulator AndroidEmulator
	conf := config.Get()

	// Clean builds if required
	if conf.Settings.CleanBeforeBuild {
		// Clean the build
		_, err := gradle.Clean(".upshift/logs/android-clean.log")
		if err != nil {
			utils.LogError("We could not clean your project. It's really dirty\n" + err.Error())
			return 1
		}
	}

	// Launch a device or emulator
	err := emulator.Launch()
	if err != nil {
		utils.LogError(err.Error())
		return 1
	}

	// Uninstall apps if required
	if conf.Settings.UninstallOlderBuilds {
		_, err = gradle.Uninstall(".upshift/logs/android-unintall.log")
		if err != nil {
			utils.LogError(err.Error())
			return 1
		}
	}

	// Run Lint
	_, err = gradle.Lint(".upshift/logs/android-lint.log")
	if err != nil {
		utils.LogError(err.Error())
		return 1
	}

	// Gradle Assemble
	_, err = gradle.Assemble(".upshift/logs/android-assemble.log")
	if err != nil {
		utils.LogError(err.Error())
		return 1
	}

	return 0
}

// AndroidDeploy : Deploy an app to crashlytics or PlayStore
func AndroidDeploy() int {
	return 1
}

// AndroidLoadEmulator : Start up the android emulator
func AndroidLoadEmulator() int {
	var emulator AndroidEmulator

	if emulator.IsEmulatorRunning() {
		utils.LogInfo("The android emulator is already running")
		return 0
	}

	devices, err := emulator.ConnectedDevices()
	if len(devices) > 0 {
		utils.LogInfo("Some android devices or emulators are already connected to this machine")
		return 0
	}

	err = emulator.Launch()
	if err != nil {
		utils.LogError(err.Error())
		return 1
	}

	return 0
}

// AndroidRun : Run an app on the android emulator or device
func AndroidRun() int {
	return 1
}

// AndroidTest : Test an android app
func AndroidTest() int {
	return 1
}

// AndroidUpgrade : Upgrade the android sdk
func AndroidUpgrade() int {
	return 1
}

// GitPull : Pull the current git repo
func GitPull() int {
	var git Git

	remote, err := git.Remote()
	if err != nil {
		utils.LogError(err.Error())
		return 1
	}

	branch, err := git.Branch()
	if err != nil {
		utils.LogError(err.Error())
		return 1
	}

	status, err := git.Pull(remote, branch)
	if err != nil {
		utils.LogError(err.Error())
		return 1
	}

	return status
}

// GitSubmodules : Setup git modules for this project
func GitSubmodules() int {
	var git Git

	if git.AreSubmodulesUsed() == false {
		utils.LogInfo("This project does not use submodules, skipping")
		return 0
	}

	status, err := git.SubmoduleInit()
	if err != nil {
		utils.LogError(err.Error())
		return 1
	}

	status, err = git.SubmoduleUpdate()
	if err != nil {
		utils.LogError(err.Error())
		return 1
	}

	return status
}

// IosArchive : Archive the current project
func IosArchive() int {
	return 1
}

// IosBuild : Build the current project
func IosBuild() int {
	return 1
}

// IosCertificates : Install the certificates for this project
func IosCertificates() int {
	return 1
}

// IosCreateApp : Create an app on iTunes
func IosCreateApp() int {
	return 1
}

// IosDeploy : Deploy an app to iTunes
func IosDeploy() int {
	return 1
}

// IosExportIPA : Export an IPA
func IosExportIPA() int {
	return 1
}

// IosPrepare : prepare the current project by reading xcodebuild
func IosPrepare() int {
	return 1
}

// IosProvisioning : Setup provisioning profile for iOS
func IosProvisioning() int {
	return 1
}

// IosRun : Run an app on the simulator
func IosRun() int {
	return 1
}

// IosSimulator : Start up the ios simulator
func IosSimulator() int {
	return 1
}

// IosTest : Test an iOS app
func IosTest() int {
	return 1
}

// SetupConfig : Setup the config.toml for this project
func SetupConfig() int {
	return 1
}

// SetupFastlane : setup the fastlane tools
func SetupFastlane() int {
	var fastlane Fastlane

	err := fastlane.Install()
	if err != nil {
		utils.LogError(err.Error())
		return 1
	}

	return 0
}

// SetupGradleWrapper : Action to install the gradle wrapper
func SetupGradleWrapper() int {
	var gradle Gradle

	// Check if gradle exists
	err := gradle.Version()
	if err != nil {
		utils.LogError(err.Error())
		return 1
	}

	// Add the wrapper
	status, err := gradle.AddWrapper()
	if err != nil {
		utils.LogError(err.Error())
	}
	return status
}

// SetupPods : Install pods for this project, if the project uses it
func SetupPods() int {
	var pod Pod

	// Check if this project uses cocoapods
	if pod.AreUsed() == false {
		fmt.Println("It looks like this project doesn't use pods")
		return 0
	}

	// Install pod
	status, err := pod.Install()
	if err != nil {
		utils.LogError("Could not install pods\n" + err.Error())
		return status
	}

	fmt.Println("We were able to successfully setup cocoapods, moving on")
	return 0
}

// SetupXcode : SetupXcode for this project
func SetupXcode() int {
	var xcodebuild Xcodebuild

	err := xcodebuild.LoadSettings()
	if err != nil {
		utils.LogError(err.Error())
		return 1
	}

	err = xcodebuild.SwitchXcode()
	if err != nil {
		utils.LogError(err.Error())
		return 1
	}

	return 0
}

// SetupXcpretty : Setup xcpretty
func SetupXcpretty() int {
	return 1
}

// SetupXctool : setup xctool
func SetupXctool() int {
	return 1
}

// ShowHelp : show help text to the user
func ShowHelp() int {
	return 1
}

// ShowVersion : Show the version of the upshift binary
func ShowVersion() int {
	var upshift Upshift
	upshift.ShowVersion()
	return 0
}

// UpgradeScript : check if you are on the latest version of upshift
func UpgradeScript() int {
	var upshift Upshift

	err := upshift.Upgrade("false")
	if err != nil {
		utils.LogError(err.Error())
	}

	return 0
}
