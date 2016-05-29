package main

import (
	"github.com/progrium/go-basher"
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

func loadBash() {
	bash, _ := basher.NewContext("/bin/bash", false)
	if bash.HandleFuncs(os.Args) {
		os.Exit(0)
	}

	bash.CopyEnv()
	bash.Source("main.bash", nil)

	// status, err := bash.Run("gradleTasks", []string{"123", "234"})
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
