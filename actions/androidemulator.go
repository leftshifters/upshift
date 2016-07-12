package actions

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/leftshifters/upshift/basher"
	"github.com/leftshifters/upshift/command"
	"github.com/leftshifters/upshift/utils"
)

// AndroidEmulator : Construct to handle all things related to the emulator
type AndroidEmulator struct{}

// Launch : launch the available emulator
func (a *AndroidEmulator) Launch() error {

	// 1. Check if any devices are connected, if yes, use one of those
	devices, err := a.ConnectedDevices()
	if err != nil {
		return err
	}

	// If a device is available, we can do everything there
	if len(devices) > 0 {
		return nil
	}

	// 2. If nothing so far, see if any avds are listed and start the first one
	avds, err := a.AVDsAvailable()
	if err != nil {
		return err
	}

	if len(avds) == 0 {
		// 3. If still nothing, create an avd and launch it
		err = a.CreateAVD()
		if err != nil {
			return err
		}
	}

	var b basher.Basher
	_, err = b.Run("AndroidLaunchEmulator", []string{avds[0], ".upshift/logs/android-emulator.log"})
	if err != nil {
		return err
	}

	return nil
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

// CreateAVD : Create an AVD
func (a *AndroidEmulator) CreateAVD() error {
	// to view a list of available avds you can create, run 'android list targets'
	// look for ones with ABIs
	// android create avd --name "Android_22_AVD" --target android-22

	// check if the ABI is available
	if utils.FileExists(filepath.Join(os.Getenv("ANDROID_HOME"), "system-images", "android-22", "default", "armeabi-v7a", "system.img")) == false {
		var b basher.Basher
		_, err := b.Run("AndroidInstallABI", []string{})
		if err != nil {
			return err
		}
	}

	var b basher.Basher
	_, err := b.Run("AndroidCreateAVD", []string{"Android_22_AVD", "Google Inc.:Google APIs:22"})
	if err != nil {
		return err
	}

	return nil
}

// DeleteAVD : delete an AVD
func (a *AndroidEmulator) DeleteAVD(avd string) error {
	// Check if AVD is available
	avds, err := a.AVDsAvailable()
	if err != nil {
		return err
	}

	var found bool
	for _, item := range avds {
		if avd == item {
			found = true
		}
	}

	if found == false {
		return errors.New("This avd does not exist")
	}

	_, err = command.Run([]string{"android", "delete", "avd", "-n", avd}, "")
	if err != nil {
		return err
	}

	return nil
}
