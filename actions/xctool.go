package actions

import "errors"

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
