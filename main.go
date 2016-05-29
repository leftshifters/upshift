package main

import (
	"log"
	"os"
	"upshift/config"
	"upshift/tasks"
)

// Main Function
func main() {

	// Load Config
	conf, err := config.Load()
	if err != nil {
		log.Println("We couldn't detect a config file, you should get one", err)
	}

	tasks.Setup(conf)

	os.Exit(0)
}
