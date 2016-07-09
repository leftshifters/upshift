package actions

import "errors"

// Fastlane : Construct to handle all things related to fastlane
type Fastlane struct{}

// Install : Install fastlane on the machine
func (f *Fastlane) Install() error {
	var g Gems
	err := g.InstallSimple("fastlane")
	if err != nil {
		return errors.New("Could not install fastlane\n" + err.Error())
	}

	return nil
}

// Uninstall : Uninstall fastlane on the machine
func (f *Fastlane) Uninstall() error {
	var g Gems
	err := g.UninstallSimple("fastlane")
	if err != nil {
		return errors.New("Could not uninstall fastlane\n" + err.Error())
	}

	return nil
}
