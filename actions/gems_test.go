package actions

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Gem_Install(t *testing.T) {
	var g Gems

	gem := "xcpretty"

	// Check if it is installed or not
	installed, _ := g.version(gem)

	// Change env setting, simulate CI
	os.Setenv("GITLAB_CI", "true")

	// Uninstall gem first
	err := g.UninstallSimple(gem)
	assert.Nil(t, err)

	// Install gem
	err = g.InstallSimple(gem)
	assert.Nil(t, err)

	// Reset CI flag
	os.Unsetenv("GITLAB_CI")

	// Bring gem back to original state
	if installed == true {
		err := g.InstallSimple(gem)
		assert.Nil(t, err)
	}
}
