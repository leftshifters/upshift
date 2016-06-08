package basher

import (
	"errors"
	"fmt"
	gobasher "github.com/progrium/go-basher"
	"os"
	colours "upshift/colours"
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
	if status > 0 || err != nil {
		errorString := "There was a problem running " + colours.Red.Format + command + colours.Default.Format + "."
		if err != nil {
			errorString += " We were stopped by the following error " + err.Error()
		}
		return status, errors.New(errorString)
	}

	return status, nil
}
