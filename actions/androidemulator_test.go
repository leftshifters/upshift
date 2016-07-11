package actions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AndroidEmulator_All(t *testing.T) {
	var emulator AndroidEmulator

	devices, err := emulator.ConnectedDevices()
	t.Log(devices)
	assert.Equal(t, 0, len(devices))
	assert.Nil(t, err)

	err = emulator.CreateAVD()
	assert.Contains(t, err.Error(), "not been implemented")

	avds, err := emulator.AVDsAvailable()
	t.Log(avds)
	assert.Equal(t, 0, len(avds))
	assert.Nil(t, err)

	err = emulator.Launch()
	assert.Nil(t, err)

	err = emulator.LaunchApp("test", "MainActivity")
	assert.Nil(t, err)
}