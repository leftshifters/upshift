package actions

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Xcpretty_Install(t *testing.T) {
	var xcpretty Xcpretty

	// Check if it is installed or not
	installed := xcpretty.IsInstalled()

	// Change env setting, simulate CI
	os.Setenv("GITLAB_CI", "true")

	// Remove Xcpretty
	err := xcpretty.Uninstall()
	assert.Nil(t, err)

	// Install Xcpretty again
	err = xcpretty.Install()
	assert.Nil(t, err)

	// Bring gem back to original state
	if installed == true {
		err := xcpretty.Install()
		assert.Nil(t, err)
	}
}
