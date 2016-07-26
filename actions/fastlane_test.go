package actions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Fastlane_Install(t *testing.T) {
	var fastlane Fastlane

	// Remove fastlane - not removing, it bumps up test time substantially
	// err := fastlane.Uninstall()
	// assert.Nil(t, err)

	// Install fastlane again
	err := fastlane.Install()
	assert.Nil(t, err)
}
