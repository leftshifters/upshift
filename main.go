package main

import (
	"github.com/leftshifters/upshift/tasks"
	"github.com/leftshifters/upshift/utils"
)

// Main Function
func main() {
	utils.SetupFolders()
	tasks.Setup()
}
