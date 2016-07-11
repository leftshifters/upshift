package actions

import (
	"errors"
	"upshift/basher"
	"upshift/utils"
)

// Xctool : Construct to handle all things related to xctool
type Xctool struct{}

// Install : Install Xctool on the machine
func (x *Xctool) Install() error {
	var b Brew
	_, err := b.Install("xctool")
	if err != nil {
		return errors.New("Could not install xctool\n" + err.Error())
	}

	return nil
}

// Upgrade : Upgrade Xctool on the machine
func (x *Xctool) Upgrade() error {
	var b Brew
	_, err := b.Upgrade("xctool")
	if err != nil {
		return errors.New("Could not upgrade xctool\n" + err.Error())
	}

	return nil
}

// Uninstall : Uninstall xctool on the machine
func (x *Xctool) Uninstall() error {
	var b Brew
	_, err := b.Uninstall("xctool")
	if err != nil {
		return errors.New("Could not uninstall xctool\n" + err.Error())
	}

	return nil
}

// Test : run tests for an iOS project
func (x *Xctool) Test(projecttype string, path string, scheme string, device string) error {
	if projecttype == "" || path == "" || scheme == "" || device == "" {
		return errors.New("We need the Type, Path, Scheme and Device to proceed")
	}

	utils.LogMessage("$ xctool -" + projecttype + " " + path + " -scheme " + scheme + " -sdk iphonesimulator -destination \"platform=iphonesimulator,name=" + device + "\" test")

	var b basher.Basher
	_, err := b.RunAndTail("IOSTest", []string{projecttype, path, scheme, device, ".upshift/logs/xcode-test.log"}, ".upshift/logs/xcode-test.log", []string{"BUILD SUCCEEDED"}, []string{})
	if err != nil {
		return err
	}

	return nil
}
