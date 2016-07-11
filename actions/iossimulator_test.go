package actions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IOSSimulator_All(t *testing.T) {
	var simulator IOSSimulator
	simulator.StopSimulator()

	simulator.StartSimulator("iPhone 6 (9.3)")

	// time.Sleep(time.Minute)

	if simulator.IsSimulatorRunning() {
		simulator.StopSimulator()
	}
}

func Test_IOSSimulator_FindDevice(t *testing.T) {
	var simulator IOSSimulator

	found := simulator.FindDevice("iPhone 6 (9.3)")
	assert.Equal(t, true, found)

	found = simulator.FindDevice("iPhone 6")
	assert.Equal(t, false, found)
}
