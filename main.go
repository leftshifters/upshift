package main

import (
	"os"

	"github.com/leftshifters/upshift/tasks"
	"github.com/leftshifters/upshift/utils"
)

// Main Function
func main() {
	utils.SetupFolders()
	os.Setenv("LANG", "en_US.UTF-8")
	tasks.Setup()
}
