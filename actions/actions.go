package actions

import (
	"upshift/utils"
)

func GradleWrapper() int {
	var g Gradle
	status, err := g.AddWrapper()
	if err != nil {
		utils.LogError(err.Error())
	}
	return status
}
