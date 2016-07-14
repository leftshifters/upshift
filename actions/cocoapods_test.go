package actions

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Cocoapods_All(t *testing.T) {
	var cocoapods Cocoapods

	// Change env setting, simulate CI
	os.Setenv("GITLAB_CI", "true")

	// Remove cocoapods
	err := cocoapods.Uninstall()
	assert.Nil(t, err)

	// Install cocoapods again
	err = cocoapods.Install()
	assert.Nil(t, err)

	// Reset CI flag
	os.Unsetenv("GITLAB_CI")

}
