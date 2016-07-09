package actions

import "testing"

func Test_IOSSimulator_All(t *testing.T) {
	var simulator IOSSimulator
	simulator.StopSimulator()

	simulator.StartSimulator("iPhone 6 (9.3)")

	// time.Sleep(time.Minute)

	if simulator.IsSimulatorRunning() {
		simulator.StopSimulator()
	}
}
