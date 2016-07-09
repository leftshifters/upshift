package actions

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"upshift/basher"
	c "upshift/colours"
	"upshift/config"
	"upshift/utils"
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

// UpgradeScript : Call this function to download the latest version of the binary
// And update the user to the latest version.
// It does nothing if the user is on the latest version
func UpgradeScript() int {
	conf := config.Get()

	resp, err := http.Get("https://raw.githubusercontent.com/leftshifters/upshift/master/release")
	if err != nil {
		fmt.Println("We couldn't connect to the internet :(", err.Error())
		return 1
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("We were unable to figure out what is the latest version on the server, next time maybe", err.Error())
		return 1
	}

	latestVersion := string(body)
	latestVersion = strings.TrimSpace(latestVersion)

	if latestVersion == conf.Settings.AppVersion {
		fmt.Println("Your powers (and version) are already at the top. You're running v", conf.Settings.AppVersion)
		return 0
	}

	fmt.Println("Get ready to feel the power at your fingertips")

	var b basher.Basher
	status, err := b.Run("UpgradeScript", []string{})
	if err != nil {
		utils.LogError("Your fingertips will suck for some more time, we couldn't upgrade you because of this - \n" + err.Error())
		return status
	}

	fmt.Println("You are now awesome. The new version of awesomeness is v", conf.Settings.AppVersion)
	return 0
}

// SetupConfig : When a new project doesn't have config, they call this one to create one for them
func SetupConfig() int {

	configExits := utils.FileExists("./config.toml")
	if configExits == true {
		fmt.Println("It looks like a config.toml is already here, skipping this step")
		return 1
	}
	// Config does not exist
	// Create a new config.toml in this directory

	sampleToml := `[Application]
Debug = false

[Runner]
RootPassword = "testPassword"

[Build]
GitRepoURL = "testRepo"
GitRepoRemote = "origin"
CleanBeforeBuild = false
UninstallOlderBuilds = false

[IOS]
ProjectName = "testProject"
UseWorkspace = false
Scheme = "testScheme"
TestDevice = "iPhone 6"
Xcode = "7.3.1"
DeveloperAccount = ""

[Android]
PackageName = "testPackage"
MainActivityName = "testActivity"`

	tomlBytes := []byte(sampleToml)

	err := ioutil.WriteFile("./config.toml", tomlBytes, 0644)
	if err != nil {
		utils.LogError("We could not write the config file, the OS told us this <" + err.Error() + ">")
		return 1
	}

	fmt.Println("We just added a config.toml to this folder!")
	return 0
}
