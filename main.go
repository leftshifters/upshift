package main

import (
	log "github.com/Sirupsen/logrus"
	"os"
	bash "upshift/bash"
)

func init() {
	log.SetOutput(os.Stderr)
	log.SetLevel(log.DebugLevel)
}

func main() {
	// output, err := bash.Bash("echo 134 | grep 134 -c")
	output, err := bash.Bash("echo 134 && echo 134| grep 134 -c")
	// output := bash.Bash("xcodebuild -project Deezeno.xcodeproj -scheme deezeno-ios -hideShellScriptEnvironment -sdk iphonesimulator -destination \"platform=iphonesimulator,name=iPhone 6\" -derivedDataPath build | tee \"xcode-build\" | xcpretty")
	// output := bash.Bash("xcodebuild -project Deezeno.xcodeproj -scheme deezeno-ios -hideShellScriptEnvironment -sdk iphonesimulator -destination \"platform=iphonesimulator,name=iPhone 6\" -derivedDataPath build")
	// output := bash.Bash("xcodebuild")

	if err != nil {
		log.Error(err)
		log.Error(err.code)
	}

	log.Info("<", output, ">")

}
