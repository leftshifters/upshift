package actions

import (
	"fmt"
	"upshift/config"
	"upshift/utils"
)

// GradleWrapper : Action to install the gradle wrapper
func GradleWrapper() int {
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

// PodInstall : Install pods for this project, if the project uses it
func PodInstall() int {
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
