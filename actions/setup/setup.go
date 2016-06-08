package setup

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"upshift/basher"
	c "upshift/colours"
	"upshift/command"
	"upshift/utils"
)

//
// Init function, doesn't do shit
//
func init() {

}

//
// Show help, so that the user knows what to do
//
func ShowHelp() (int, bool) {
	fmt.Println("\nUPSHIFT(1)               Upshift Commands Manual               UPSHIFT(1)")
	fmt.Println(c.Bold + "\nNAME" + c.Default)
	fmt.Println("\tupshift -- the creative mobile app builder")
	fmt.Println(c.Bold + "\nSYNOPSIS" + c.Default)
	fmt.Println("\tupshift " + c.Underline + "job" + c.Default + " " + c.Underline + "action" + c.Default)
	fmt.Println(c.Bold + "\nDESCRIPTION" + c.Default)
	fmt.Println("\tThis tool helps you run, build, test and deploy your iOS and Android\n\tapps while you dream about the next big thing")

	fmt.Println(c.Bold + "\nOPTIONS (job queues)" + c.Default)
	fmt.Println("\tIt is still not as awesome as we want it to be. But here are the things\n\tthat you can currently do\n")
	fmt.Println("\tupshift ios build\n\t\tto build your iOS project\n")
	fmt.Println("\tupshift ios run\n\t\tto run your iOS project in a simulator\n")
	fmt.Println("\tupshift ios deploy\n\t\tto create an .ipa and deploy it on TestFlight\n")
	fmt.Println("\tupshift android build\n\t\tto build your Android project\n")
	fmt.Println("\tupshift android run\n\t\tto run your Android project in a simulator\n")
	fmt.Println("\tupshift android deploy\n\t\tto create an .apk and upload it to Fabric\n")
	fmt.Println("\tupshift setup clone\n\t\tto clone a repo defined in config.toml\n")
	fmt.Println("\tupshift setup config\n\t\tto setup an empty config.toml in your current folder\n")
	fmt.Println("\tupshift install\n\t\tto install this binary for the first time\n")
	fmt.Println("\tupshift -v\n\t\tto view the version number\n")

	fmt.Println(c.Bold + "\nOPTIONS (specific actions)" + c.Default)
	fmt.Println("\tWe combine actions like these to create the jobs above, you should ideally\n\tbe running jobs not actions\n")
	fmt.Println("\tupshift action setupSsh -- to setup your ssh keys")
	fmt.Println("\tupshift action setupScript -- to setup this very script")
	fmt.Println("\tupshift action setupGradle -- to setup gradle on your machine")
	fmt.Println("\tupshift action setupPods -- to setup cocoapods on your machine")
	fmt.Println("\tupshift action setupXcode -- to choose the correct xcode version for the project")
	fmt.Println("\tupshift action setupXcpretty -- to setup xcpretty for build output which doesn't suck")
	fmt.Println("\tupshift action upgradeScript -- to upgrade this script")

	fmt.Println("\tupshift action gitPull -- to pull from code from a repo")
	fmt.Println("\tupshift action gitClone -- to clone a repo")
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

	return 0, false
}

func SetupXcpretty() (int, bool) {
	// Check which version of Xcpretty was installed
	version, err := command.RunWithoutStdout([]string{"xcpretty", "--version"}, "")
	if err == nil {
		fmt.Println("Xcpretty is pretty much setup on this system. You are on version " + version)
		return 0, false
	}

	// Check if the command was not found
	if strings.Contains(err.Error(), "executable file not found") == true {
		// Alright, so Xcpretty was not found, go ahead and install it
		fmt.Println("Xcpretty was not found, installing it")

		var RootPassword string

		if utils.IsCI() == true {
			// We are on CI, we need to enter password programatically
			RootPassword, err = utils.GetRootPassword()
			if err != nil {
				utils.LogError(err.Error())
				return 1, true
			}
		}
		// else we are not on CI, ask the user to enter the password

		status, err := basher.Run("SetupXcpretty", []string{strconv.FormatBool(utils.IsCI()), RootPassword})
		if err != nil {
			utils.LogError("We couldn't install xcpretty, you will get shitty outout from Xcode - \n" + err.Error())
			return status, false
		}

		fmt.Println("Xcpretty has been setup on your machine. Have fun.")
		return 0, false
	}

	fmt.Println("There was a problem installing xcpretty - " + err.Error())
	return 1, true
}

//
// Call this function to download the latest version of the binary
// And update the user to the latest version.
// It does nothing if the user is on the latest version
//
func UpgradeScript() (int, bool) {
	resp, err := http.Get("https://raw.githubusercontent.com/leftshifters/upshift/master/release")
	if err != nil {
		fmt.Println("We couldn't connect to the internet :(", err.Error())
		return 1, false
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("We were unable to figure out what is the latest version on the server, next time maybe", err.Error())
		return 1, false
	}

	latestVersion := string(body)
	latestVersion = strings.TrimSpace(latestVersion)

	if latestVersion == utils.GetAppVersion() {
		fmt.Println("Your powers (and version) are already at the top. You're running v", utils.GetAppVersion())
		return 0, false
	}

	fmt.Println("Get ready to feel the power at your fingertips")

	status, err := basher.Run("UpgradeScript", []string{})
	if err != nil {
		utils.LogError("Your fingertips will suck for some more time, we couldn't upgrade you because of this - \n" + err.Error())
		return status, false
	}

	fmt.Println("You are now awesome. The new version of awesomeness is v", utils.GetAppVersion())
	return 0, false
}

//
// When a new project doesn't have config, they call this one to create one for them
//
func SetupConfig() (int, bool) {

	configExits := utils.FileExists("./config.toml")
	if configExits == true {
		fmt.Println("It looks like a config.toml is already here, skipping this step")
		return 1, false
	} else {
		// Config does not exist
		// Create a new config.toml in this directory

		sampleToml := `[Application]
Debug = false

[Runner]
RootPassword = "testPassword"

[Build]
GitRepoURL = "testRepo"
GitRepoBranch = "testBranch"
CleanBeforeBuild = false
UninstallOlderBuilds = false

[IOS]
ProjectName = "testProject"
UseWorkspace = false
Scheme = "testScheme"
TestDevice = "iPhone 6"
Xcode = "7.3.1"

[Android]
PackageName = "testPackage"
MainActivityName = "testActivity"`

		tomlBytes := []byte(sampleToml)

		err := ioutil.WriteFile("./config.toml", tomlBytes, 0644)
		if err != nil {
			utils.LogError("We could not write the config file, the OS told us this <" + err.Error() + ">")
			return 1, false
		}
	}

	fmt.Println("We just added a config.toml to this folder!")
	return 0, false
}
