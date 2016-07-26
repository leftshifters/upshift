package actions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Cocoapods_All(t *testing.T) {
	var cocoapods Cocoapods

	// Remove cocoapods - Removing this case because it bumps up the test time substantially
	// err := cocoapods.Uninstall()
	// assert.Nil(t, err)

	// Install cocoapods again
	err := cocoapods.Install()
	assert.Nil(t, err)

}
