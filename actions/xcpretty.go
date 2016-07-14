package actions

import "errors"

// Xcpretty : Construct to handle all things related to xcpretty
type Xcpretty struct{}

// Install : Install xcpretty on the machine
func (x *Xcpretty) Install() error {
	var g Gems
	err := g.InstallSimple("xcpretty")
	if err != nil {
		return errors.New("Could not install xcpretty\n" + err.Error())
	}

	return nil
}

// Uninstall : Uninstall xcpretty on the machine
func (x *Xcpretty) Uninstall() error {
	var g Gems
	err := g.UninstallSimple("xcpretty")
	if err != nil {
		return errors.New("Could not uninstall xcpretty\n" + err.Error())
	}

	return nil
}

// IsInstalled : check if xcpretty is installed or not
func (x *Xcpretty) IsInstalled() bool {
	var g Gems
	found, _ := g.version("xcpretty")
	return found
}
