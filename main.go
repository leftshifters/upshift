package main

import (
	"fmt"
	colours "upshift/colours"
	"upshift/config"
	"upshift/tasks"
)

// Main Function
func main() {

	// Load Config
	conf, err := config.Load()
	if err != nil {
		fmt.Println("We couldn't detect a config file, you should get one", err)
	}
	fmt.Println(colours.Red.Format + "Config:" + colours.Default.Format)
	fmt.Println(conf)

	tasks.Setup(conf)
}
