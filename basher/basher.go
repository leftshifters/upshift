package basher

import (
	"errors"
	gobasher "github.com/progrium/go-basher"
	"os"
	"path/filepath"
	"strings"
	c "upshift/colours"
	"upshift/utils"
)

type Basher struct {
	bash *gobasher.Context
}

func (b *Basher) Load() {
	b.bash, _ = gobasher.NewContext("/bin/bash", false)
	if b.bash.HandleFuncs(os.Args) {
		os.Exit(0)
	}

	b.bash.CopyEnv()
	b.bash.Source("scripts.bash", Asset)
}

func (b *Basher) Run(command string, params []string) (int, error) {

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

func (b *Basher) RunAndTail(command string, params []string, logPath string, success string) error {

	logPath, _ = filepath.Abs(logPath)

	_, err := b.Run(command, params)
	if err != nil {
		return errors.New(command + " could not be completed\n" + err.Error())
	}

	tail, err := utils.FileTail(logPath, 500)
	if err != nil {
		return errors.New("We could not read the logFile " + logPath + "\n" + err.Error())
	}

	if strings.Contains(tail, success) == false {
		return errors.New("The command " + command + " did not run successfully. You need to look this up")
	}

	return nil
}
