package main

import (
	// "fmt"
	// colours "upshift/colours"
	"upshift/config"
	"upshift/tasks"
	"upshift/utils"
)

// Main Function
func main() {

	// Load Config
	conf, err := config.Load()
	if err != nil {
		utils.LogInfo(err.Error())
	}
	// fmt.Println(colours.Red.Format + "Config:" + colours.Default.Format)
	// fmt.Println(conf)

	tasks.Setup(conf)
}
