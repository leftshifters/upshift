package actions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Gem_Install(t *testing.T) {
	var g Gems

	gem := "xcpretty"

	// Uninstall gem first
	err := g.UninstallSimple(gem)
	assert.Nil(t, err)

	// Install gem
	err = g.InstallSimple(gem)
	assert.Nil(t, err)

}
