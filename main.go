package main

import (
	"upshift/tasks"
	"upshift/utils"
)

// Main Function
func main() {
	utils.SetupFolders()
	tasks.Setup()
}
