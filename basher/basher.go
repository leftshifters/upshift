package basher

import (
	"errors"
	"fmt"
	gobasher "github.com/progrium/go-basher"
	"os"
)

var bash *gobasher.Context

func init() {
	Load()
}

func Load() {
	bash, _ = gobasher.NewContext("/bin/bash", false)
	if bash.HandleFuncs(os.Args) {
		os.Exit(0)
	}

	bash.CopyEnv()
	bash.Source("scripts/main.bash", Asset)
}

func Run(command string, params []string) (int, error) {

	if bash == nil {
		fmt.Println("bash was null, calling again")
		Load()
	}

	status, err := bash.Run(command, params)
	fmt.Println("STATUS")
	fmt.Println(status)
	fmt.Println("ERROR")
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		return status, errors.New("We were unable to run the bash task" + err.Error())
	}

	return status, nil
}
