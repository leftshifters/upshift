package actions

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/leftshifters/upshift/basher"
	"github.com/leftshifters/upshift/colours"
	"github.com/leftshifters/upshift/command"
	"github.com/leftshifters/upshift/utils"
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

// FindDevice : find if the suggested device is available on this machine
func (i *IOSSimulator) FindDevice(device string) bool {
	instrumentsDump, err := command.Run([]string{"instruments", "-s", "devices"}, "")
	if err != nil {
		return false
	}

	instrumentRows := strings.Split(instrumentsDump, "\n")
	uuidRegexp, _ := regexp.Compile("\\[(.*?)\\]")
	var instruments []string

	for _, instrument := range instrumentRows {
		uuid := uuidRegexp.FindString(instrument)

		isSimulator := strings.Contains(instrument, "(Simulator)")
		instrument = strings.Replace(instrument, "(Simulator)", "", 1)

		simulatorString := "DEVICE   "
		if isSimulator == true {
			simulatorString = "SIMULATOR"
		}

		if uuid != "" {
			instrument = strings.TrimSpace(strings.Replace(instrument, uuid, "", 1))
			instruments = append(instruments, colours.Gray+simulatorString+colours.Default+" "+colours.Bold+colours.Green+instrument+colours.Default+" "+uuid)
		}

		if instrument == device {
			return true
		}
	}

	// If device is not available, show all devices that are available
	utils.LogError("Your device " + colours.Red + device + colours.Default + " was not found\nThe following devices are available")
	for _, item := range instruments {
		fmt.Println(item)
	}
	return false
}
