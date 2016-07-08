package actions

import (
	"upshift/utils"
)

// GradleWrapper : Action to install the gradle wrapper
func GradleWrapper() int {
	var g Gradle

	// Check if gradle exists
	err := g.Version()
	if err != nil {
		utils.LogError(err.Error())
		return 1
	}

	// Add the wrapper
	status, err := g.AddWrapper()
	if err != nil {
		utils.LogError(err.Error())
	}
	return status
}
