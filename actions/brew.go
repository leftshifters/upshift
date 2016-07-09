package actions

import (
	"upshift/basher"
	"upshift/command"
)

// Brew : Construct to take care of all brew tasks
type Brew struct {
}

// Install a cask using brew
func (b *Brew) Install(cask string) (int, error) {
	if b.isInstalled(cask) == true {
		return 0, nil
	}

	var basher basher.Basher
	return basher.Run("BrewInstall", []string{cask})
}

// Upgrade a cask using brew
func (b *Brew) Upgrade(cask string) (int, error) {
	if b.isInstalled(cask) == false {
		return b.Install(cask)
	}

	var basher basher.Basher
	return basher.Run("BrewUpgrade", []string{cask})
}

// Uninstall a cask using brew
func (b *Brew) Uninstall(cask string) (int, error) {
	if b.isInstalled(cask) == false {
		return 0, nil
	}

	var basher basher.Basher
	return basher.Run("BrewUninstall", []string{cask})
}

func (b *Brew) isInstalled(cask string) bool {
	// Check if the cask is installed
	_, err := command.Run([]string{cask, "--version"}, "")
	if err == nil {
		// Remove the name of the tool if it is part of the version string
		// version = strings.Replace(version, cask, "", 1)
		// Now trim whatever is left
		// version = strings.TrimSpace(version)
		return true
	}
	return false
}
