package actions

import (
	"fmt"
	c "upshift/colours"
	"upshift/config"
)

// ShowVersion : shows the version of upshift on the command line
func ShowVersion() int {
	conf := config.Get()
	fmt.Println(conf.Settings.AppVersion)
	return 0
}

// ShowHelp : Show help, so that the user knows what to do
func ShowHelp() int {
	fmt.Println("\nUPSHIFT(1)               Upshift Commands Manual               UPSHIFT(1)")
	fmt.Println(c.Bold + "\nNAME" + c.Default)
	fmt.Println("\tupshift -- the creative mobile app builder")
	fmt.Println(c.Bold + "\nSYNOPSIS" + c.Default)
	fmt.Println("\tupshift " + c.Underline + "job" + c.Default + " " + c.Underline + "action" + c.Default)
	fmt.Println(c.Bold + "\nDESCRIPTION" + c.Default)
	fmt.Println("\tThis tool helps you run, build, test and deploy your iOS and Android\n\tapps while you dream about the next big thing")

	fmt.Println(c.Bold + "\nOPTIONS (job queues)" + c.Default)
	fmt.Println("\tIt is still not as awesome as we want it to be. But here are the things\n\tthat you can currently do")
	fmt.Println("\tupshift ios build\n\t\tto build your iOS project")
	fmt.Println("\tupshift ios run\n\t\tto run your iOS project in a simulator")
	fmt.Println("\tupshift ios deploy\n\t\tto create an .ipa and deploy it on TestFlight")
	fmt.Println("\tupshift android build\n\t\tto build your Android project")
	fmt.Println("\tupshift android run\n\t\tto run your Android project in a simulator")
	fmt.Println("\tupshift android deploy\n\t\tto create an .apk and upload it to Fabric")
	fmt.Println("\tupshift setup clone\n\t\tto clone a repo defined in config.toml")
	fmt.Println("\tupshift setup config\n\t\tto setup an empty config.toml in your current folder")
	fmt.Println("\tupshift setup export.plist\n\t\tto setup a sample .private/export.plist in your project")
	fmt.Println("\tupshift install\n\t\tto install this binary for the first time")
	fmt.Println("\tupshift -v\n\t\tto view the version number")

	fmt.Println(c.Bold + "\nOPTIONS (specific actions)" + c.Default)
	fmt.Println("\tWe combine actions like these to create the jobs above, you should ideally\n\tbe running jobs not actions")
	fmt.Println("\tupshift action setupSsh -- to setup your ssh keys")
	fmt.Println("\tupshift action setupScript -- to setup this very script")
	fmt.Println("\tupshift action setupGradleW -- to setup gradle on your machine")
	fmt.Println("\tupshift action setupPods -- to setup cocoapods on your machine")
	fmt.Println("\tupshift action setupXcode -- to choose the correct xcode version for the project")
	fmt.Println("\tupshift action setupXcpretty -- to setup xcpretty for build output which doesn't suck")
	fmt.Println("\tupshift action upgradeScript -- to upgrade this script")

	fmt.Println("\tupshift action gitPull -- to pull from code from a repo")
	// fmt.Println("\tupshift action gitClone -- to clone a repo")
	fmt.Println("\tupshift action gitSubmodules -- to setup git modules in the project")

	fmt.Println("\tupshift action iosSimulator -- to start the iOS simulator")
	fmt.Println("\tupshift action iosBuild -- to build an iOS app")
	fmt.Println("\tupshift action iosRun -- to run an iOS app in the simulator")
	fmt.Println("\tupshift action iosDeploy -- to archive and deploy an iOS app")

	fmt.Println("\tupshift action androidEmulator -- to start the android emulator")
	fmt.Println("\tupshift action androidBuild -- to build an android project")
	fmt.Println("\tupshift action androidRun -- to run an android project")
	fmt.Println("\tupshift action androidDeploy -- to deploy an android project")

	fmt.Println(c.Bold + "\nCOMPATIBILITY" + c.Default)
	fmt.Println("\tWe've only tested this on Mac OSX, Linux and Docker. If you're on\n\twindows, you should switch operating systems because nobody can help\n\tyou there.")
	fmt.Println("\nLeftshift Technologies           Made with ❤️  in India                " + c.Underline + "https://leftshift.io\n" + c.Default)

	return 0
}
