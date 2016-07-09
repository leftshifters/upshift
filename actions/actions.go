package actions

import (
	"fmt"
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
