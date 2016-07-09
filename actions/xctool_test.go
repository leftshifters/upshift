package actions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Xctool_All(t *testing.T) {
	var xctool Xctool

	// Remove xctool
	err := xctool.Uninstall()
	assert.Nil(t, err)

	// Install xctool again
	err = xctool.Install()
	assert.Nil(t, err)

	// Upgrade xctool
	err = xctool.Upgrade()
	assert.Nil(t, err)
}
