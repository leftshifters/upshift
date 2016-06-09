package main

import (
	"upshift/config"
	"upshift/tasks"
	"upshift/utils"
)

// Main Function
func main() {

	// Load Config
	_, err := config.Load()
	if err != nil {
		utils.LogInfo(err.Error())
	}

	tasks.Setup()
}
