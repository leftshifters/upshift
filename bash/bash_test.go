package bash

import (
	log "github.com/Sirupsen/logrus"
	"testing"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func showResult(success bool, name string) {
	if success == true {
		log.Info("✅  " + name)
	} else {
		log.Error("🍎  " + name)
	}
}

func TestBash(test *testing.T) {

	var success bool

	// _, err := Bash("")
	// if err.Error() != "You need to enter something for this to work" {
	// 	test.Fail()
	// 	success = false
	// } else {
	// 	success = true
	// }
	// showResult(success, "Checking for empty params")

	// output, _ := Bash("echo 123")
	output, err := Bash("xcodebuild -showBuildSettings")
	// output, err := Bash("sleep 3")
	if output != "123" {
		test.Fail()
		success = false
	} else {
		success = true
	}
	showResult(success, "Checking for simple echo commands")

	// output, _ := Bash("echo 134 | grep 13 -c | wc -l")
	// if output != "1" {
	// 	test.Fail()
	// 	success = false
	// } else {
	// 	success = true
	// }
	// showResult(success, "Checking with a pipe function")

	// output, err = Bash("echo 123 && echo 234 | grep 234 -c")
	// if err.Error() != "Found an unsupported action in the shell script" {
	// 	test.Fail()
	// 	success = false
	// } else {
	// 	success = true
	// }
	// showResult(success, "Checking with && in the sequence")

	// output, err := Bash("xcodebuild -showBuildSettings | tee xcode-build.log | xcpretty")
	// output, err := Bash("xcodebuild -showBuildSettings")

	if err != nil {
		log.Info("TEST ERROR ", err.Error())
	}
	log.Info("TEST OUTPUT ", output)

}
