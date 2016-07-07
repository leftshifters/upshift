package basher

import (
	"errors"
	gobasher "github.com/progrium/go-basher"
	"os"
	c "upshift/colours"
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
