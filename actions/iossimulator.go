package actions

import (
	"fmt"
	"upshift/basher"
	"upshift/command"
)

// IOSSimulator : Construct to handle everything on the ios simulator
type IOSSimulator struct {
}

// IsSimulatorRunning : find out if the simulator is running
func (i *IOSSimulator) IsSimulatorRunning() bool {
	_, err := command.Run([]string{"pgrep", "-f", "Simulator.app"}, "")
	if err != nil {
		return false
	}
	return true
}

// StartSimulator : start up the iOS simulator
func (i *IOSSimulator) StartSimulator(device string) {
	// If simulator is already running, just ignore
	if i.IsSimulatorRunning() {
		return
	}

	// basher returns an error if status > 0 or if there is an error
	// Whenever we start the simulator, for some reason, the exit code is always 255, though there is no error
	// Hence skipping the error check here
	var b basher.Basher
	b.Run("StartSimulator", []string{device})
}

// StopSimulator : stop the simulator
func (i *IOSSimulator) StopSimulator() {
	// If simulator is not running, just go back
	if i.IsSimulatorRunning() == false {
		return
	}

	// #TODO : If simulator is running, kill it
	fmt.Println("We don't know how to kill the simular yet")
}
