package main

import (
	"github.com/progrium/go-basher"
	"log"
	"os"
)

// Main Function
func main() {
	bash, _ := basher.NewContext("/bin/bash", false)
	bash.ExportFunc("reverse", reverse)
	if bash.HandleFuncs(os.Args) {
		os.Exit(0)
	}

	bash.CopyEnv()
	bash.Source("main.bash", nil)

	status, err := bash.Run("main", os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(status)
}

func reverse(str []string) {

}
