package actions

import (
	"errors"
	"fmt"
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

// ConnectedDevices : find which devices are connected
func (a *AndroidEmulator) ConnectedDevices() ([]string, error) {
	out, err := command.Run([]string{"adb", "devices"}, "")
	if err != nil {
		return []string{}, err
	}

	devices := utils.CreateList(out, []string{"List of devices attached", "daemon not running. starting it now on port", "daemon started successfully", "offline"})
	return devices, nil
}

// LaunchApp : start the app in the emulator
func (a *AndroidEmulator) LaunchApp(packageName string, activityName string) error {
	if packageName == "" {
		return errors.New("We need the package name to launch the app")
	}

	if activityName == "" {
		return errors.New("We need the main activity name to launch the app")
	}

	out, err := command.Run([]string{"adb", "shell", "am", "start", "-n", packageName + "/" + packageName + "." + activityName}, "")
	if err != nil {
		return err
	}
	fmt.Println(out)

	return nil
}
