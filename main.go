package main

import (
	"github.com/progrium/go-basher"
	"log"
	"os"
	"path/filepath"
	"upshift/config"
	lc "upshift/log-colours"
	"upshift/utils"
)

// Main Function
func main() {
	bash, _ := basher.NewContext("/bin/bash", false)
	bash.ExportFunc("reverse", reverse)
	if bash.HandleFuncs(os.Args) {
		os.Exit(0)
	}

	// Setup Config
	fileFullPath, err := filepath.Abs("./config/sample.toml")
	if err != nil {
		log.Println(err)
	}

	conf, err := config.Load(fileFullPath)
	if err != nil {
		log.Println(err)
	}
	log.Println("Logging is", conf.Application.Debug)

	bash.CopyEnv()
	bash.Source("main.bash", nil)

	// TODO : Add xterm if TERM is empty

	log.Println("Getting", lc.Blue.Format, "started", lc.Default.Format, "again.")

	// TODO : Read from env masterPasswordFromEnv
	// TODO : Read GITLAB_CI from env

	inDocker, _ := utils.IsDocker()
	log.Println("Are we in docker? ", inDocker)

	status, err := bash.Run("main", os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(status)
}

func reverse(str []string) {

}
