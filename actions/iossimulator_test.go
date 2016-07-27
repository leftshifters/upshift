package actions

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_IOSSimulator_All(t *testing.T) {
	var simulator IOSSimulator
	simulator.StopSimulator()

	simulator.StartSimulator("iPhone 6 (9.3)")

	fmt.Println("Sleeping for 15 seconds while the simulator loads up")
	time.Sleep(time.Second * 15)

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
