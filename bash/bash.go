package bash

import (
	"errors"
	"fmt"
	"github.com/progrium/go-basher"
	"os"
)

var bash *basher.Context

func init() {
	Load()
}

func Load() {
	bash, _ = basher.NewContext("/bin/bash", false)
	if bash.HandleFuncs(os.Args) {
		os.Exit(0)
	}

	bash.CopyEnv()
	bash.Source("main.bash", nil)
}

func Run(command string, params []string) (int, error) {

	if bash == nil {
		Load()
	}

	status, err := bash.Run(command, params)
	if err != nil {
		fmt.Println(err)
		return status, errors.New("We were unable to run the bash task" + err.Error())
	}

	return status, nil
}
