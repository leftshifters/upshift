package actions

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/leftshifters/upshift/basher"
	"github.com/leftshifters/upshift/command"
	"github.com/leftshifters/upshift/config"
)

// Gems structure to install, upgrade, uninstall gems
type Gems struct{}

// InstallSimple a gem with gem name only
func (g *Gems) InstallSimple(gem string) error {
	return g.Install(gem, "")
}

// Install a gem with gem name and scriptName
func (g *Gems) Install(gem string, scriptName string) error {
	// Get the global config
	conf := config.Get()

	// If scriptName is empty, use the gem name
	if scriptName == "" {
		scriptName = gem
	}

	// Check if gem is installed
	installed, _ := g.version(scriptName)
	if installed == true {
		return nil
	}

	// Okay, so the gem was not installed, big deal
	// Let's install the effing thing
	var b basher.Basher
	RootPassword, err := conf.GetRootPassword()
	if conf.IsCI() == true && err != nil {
		return errors.New("Unable to find the RootPassword\n" + err.Error())
	}

	status, err := b.Run("SetupGem", []string{gem, strconv.FormatBool(conf.IsCI()), RootPassword})
	if err != nil {
		return errors.New("Could not install " + gem + " [Error Code <" + strconv.Itoa(status) + ">]\n" + err.Error())
	}

	// Verify if gem was installed
	installed, _ = g.version(scriptName)
	if installed == true {
		// It was successfull installed
		return nil
	}
	return errors.New("The script ran but the gem " + gem + " was not installed on the machine")
}

// UninstallSimple a gem with gem name only
func (g *Gems) UninstallSimple(gem string) error {
	return g.Uninstall(gem, "")
}

// Uninstall : remove a gem from the system
func (g *Gems) Uninstall(gem string, scriptName string) error {
	// If scriptName is not set, use the gem name
	if scriptName == "" {
		scriptName = gem
	}

	// Check if the gem is installed
	installed, _ := g.version(scriptName)
	if installed == false {
		// If it is not installed then return
		return nil
	}

	// Get the global config
	conf := config.Get()

	// So we know that it is installed, go ahead and uninstall it
	// Let's uninstall the effing thing
	var b basher.Basher
	RootPassword, err := conf.GetRootPassword()
	if conf.IsCI() == true && err != nil {
		return errors.New("Unable to find the RootPassword\n" + err.Error())
	}

	status, err := b.Run("UninstallGem", []string{gem, strconv.FormatBool(conf.IsCI()), RootPassword})
	if err != nil {
		return errors.New("Could not uninstall " + gem + " [Error Code <" + strconv.Itoa(status) + ">]\n" + err.Error())
	}

	// Verify if gem was uninstalled
	installed, _ = g.version(scriptName)
	if installed == false {
		// It was successfull uninstalled
		return nil
	}
	return errors.New("The script ran but the gem " + gem + " was not installed on the machine")
}

func (g *Gems) version(gem string) (bool, string) {
	// Check which version of the gem was installed
	ver, err := command.Run([]string{gem, "--version"}, "")
	if err == nil {
		// This means that the gem is installed
		// Get rid of the name and get the version number
		ver = strings.Replace(ver, gem, "", 1)
		// Next trim what's left
		ver = strings.TrimSpace(ver)
		fmt.Println(gem + " v" + ver + " is installed on your machine")
		return true, ver
	}

	// Not installed
	return false, ""
}

// #TODO : Find out how do we know that an update is available and update it
// When an upgrade is available, they say
// CocoaPods 1.0.1 is available.
// To update use: `sudo gem install cocoapods`
// Until we reach version 1.0 the features of CocoaPods can and will change.
// We strongly recommend that you use the latest version at all times.
// if strings.Contains(tailData, "sudo gem install cocoapods") == true {
// This means that an update is available, run cocoapods update
//	status := SetupPods(true)
// if status > 0 {
// return errors.New("We couldn't update to the new version of cocoapods")
// }
// fmt.Println("Updated cocoapods to the latest version")
// }
