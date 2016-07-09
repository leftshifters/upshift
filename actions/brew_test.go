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

	t.Log("Uninstalling")
	status, err := brew.Uninstall("xctool")
	assert.Equal(t, 0, status)
	assert.Nil(t, err)

	t.Log("Installing")
	status, err = brew.Install("xctool")
	assert.Equal(t, 0, status)
	assert.Nil(t, err)

	t.Log("Upgrading")
	status, err = brew.Upgrade("xctool")
	assert.Equal(t, 0, status)
	assert.Nil(t, err)

}
