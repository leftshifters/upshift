package actions

import "errors"

// Cocoapods : Construct to handle all things related to fastlane
type Cocoapods struct{}

// Install : Install cocoapods on the machine
func (c *Cocoapods) Install() error {
	var g Gems
	err := g.Install("cocoapods", "pod")
	if err != nil {
		return errors.New("Could not install cocoapods\n" + err.Error())
	}

	return nil
}

// Uninstall : Uninstall cocoapods on the machine
func (c *Cocoapods) Uninstall() error {
	var g Gems
	err := g.Uninstall("cocoapods", "pod")
	if err != nil {
		return errors.New("Could not uninstall cocoapods\n" + err.Error())
	}

	return nil
}

// IsInstalled : check if cocoapods is installed
func (c *Cocoapods) IsInstalled() bool {
	var g Gems

	installed, _ := g.version("pod")
	return installed
}
