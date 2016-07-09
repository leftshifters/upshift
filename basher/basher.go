package basher

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	c "upshift/colours"
	"upshift/utils"

	gobasher "github.com/progrium/go-basher"
)

// Basher : Construct to handle everything related to basher and bash scripts
type Basher struct {
	bash *gobasher.Context
}

// Load : Load the basher binary, so that you can start running scripts
func (b *Basher) Load() {
	b.bash, _ = gobasher.NewContext("/bin/bash", false)
	if b.bash.HandleFuncs(os.Args) {
		os.Exit(0)
	}

	b.bash.CopyEnv()
	b.bash.Source("scripts.bash", Asset)
}

// Run : Run a command on bash
func (b *Basher) Run(command string, params []string) (int, error) {
	utils.SetupFolders()

	if b.bash == nil {
		b.Load()
	}

	status, err := b.bash.Run(command, params)
	if status > 0 || err != nil {
		errorString := "There was a problem running " + c.Red + command + c.Default + "\n"
		if err != nil {
			errorString += "We were stopped by the following error\n" + err.Error()
		}
		return status, errors.New(errorString)
	}

	return status, nil
}

// RunAndTail : Run a command on basher and tail it's output when the command finishes to look for specific words
func (b *Basher) RunAndTail(command string, params []string, logPath string, success []string, failure []string) (int, error) {
	utils.SetupFolders()

	logPath, _ = filepath.Abs(logPath)

	status, err := b.Run(command, params)
	if err != nil {
		return status, errors.New(command + " could not be completed\n" + err.Error())
	}

	tail, err := utils.FileTail(logPath, 500)
	if err != nil {
		return status, errors.New("We could not read the logFile " + logPath + "\n" + err.Error())
	}

	//
	// Look through success and failure cases
	//

	// If any anti success cases are found, return with error
	for _, successItem := range success {
		if strings.Contains(tail, successItem) == false {
			return status, errors.New("The command " + command + " did not run successfully. You need to look this up")
		}
	}

	// If any failure cases are found, return with error
	for _, successItem := range failure {
		if strings.Contains(tail, successItem) == true {
			return status, errors.New("The command " + command + " did not run successfully. You need to look this up")
		}
	}

	return status, nil
}
