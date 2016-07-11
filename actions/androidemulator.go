package actions

import (
	"upshift/command"
	"upshift/utils"
)

// AndroidEmulator : Construct to handle all things related to the emulator
type AndroidEmulator struct{}

// Launch : launch the available emulator
func (a *AndroidEmulator) Launch() {

}

// AVDsAvailable : find which AVDs are available
func (a *AndroidEmulator) AVDsAvailable() ([]string, error) {
	out, err := command.Run([]string{"emulator", "-list-avds"}, "")
	if err != nil {
		return []string{}, err
	}

	avds := utils.CreateList(out, []string{})
	return avds, nil
}
