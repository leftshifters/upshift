package actions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Brew_isInstalled(t *testing.T) {
	var brew Brew
	installed := brew.isInstalled("xctool")
	assert.Equal(t, true, installed)
}

func Test_Brew_All(t *testing.T) {
	var brew Brew

	status, err := brew.Uninstall("xctool")
	assert.Equal(t, 0, status)
	assert.Nil(t, err)

	status, err = brew.Install("xctool")
	assert.Equal(t, 0, status)
	assert.Nil(t, err)

	status, err = brew.Upgrade("xctool")
	assert.Equal(t, 0, status)
	assert.Nil(t, err)

}
