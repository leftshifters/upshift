package actions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Xcpretty_Install(t *testing.T) {
	var xcpretty Xcpretty

	// Remove Xcpretty - won't remove, increases build time majorly
	// err := xcpretty.Uninstall()
	// assert.Nil(t, err)

	// Install Xcpretty again
	err := xcpretty.Install()
	assert.Nil(t, err)

}
