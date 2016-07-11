package actions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Android_All(t *testing.T) {
	var android Android

	err := android.SetupAndroid("android-setup.log")
	assert.Nil(t, err)

	err = android.UpgradeAndroid("android-upgrade.log")
	assert.Nil(t, err)
}
