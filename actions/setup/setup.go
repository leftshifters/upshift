package setup

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"upshift/basher"
	colours "upshift/colours"
	"upshift/utils"
)

func init() {

}

func ShowHelp() (int, bool) {
	fmt.Println("\nUPSHIFT(1)               Upshift Commands Manual               UPSHIFT(1)")
	fmt.Println(colours.Bold.Format + "\nNAME" + colours.Default.Format)
	fmt.Println("\tupshift -- the creative mobile app builder")
	fmt.Println(colours.Bold.Format + "\nSYNOPSIS" + colours.Default.Format)
	fmt.Println("\tupshift " + colours.Underline.Format + "job" + colours.Default.Format + " " + colours.Underline.Format + "action" + colours.Default.Format)
	fmt.Println(colours.Bold.Format + "\nDESCRIPTION" + colours.Default.Format)
	fmt.Println("\tThis tool helps you run, build, test and deploy your iOS and Android\n\tapps while you dream about the next big thing")

	fmt.Println(colours.Bold.Format + "\nOPTIONS (job queues)" + colours.Default.Format)
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

	fmt.Println(colours.Bold.Format + "\nOPTIONS (specific actions)" + colours.Default.Format)
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

	fmt.Println(colours.Bold.Format + "\nCOMPATIBILITY" + colours.Default.Format)
	fmt.Println("\tWe've only tested this on Mac OSX, Linux and Docker. If you're on\n\twindows, you should switch operating systems because nobody can help\n\tyou there.")
	fmt.Println("\nLeftshift Technologies           Made with ❤️  in India                " + colours.Underline.Format + "https://leftshift.io\n" + colours.Default.Format)

	return 0, false
}

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

// function UpgradeVersion {

//   StartAction "UpgradeVersion"

//   # If 'next' is false, exit
//   if [ ${next} == false ]; then
//     ShowPreviousFailed
//     return
//   fi

//   LATEST_VERSION=$(curl -sS https://raw.githubusercontent.com/leftshifters/upshift/master/release 2>&1)
//   LATEST_VERSION_RESULTS=$(grep 'curl:' -c <<< "${LATEST_VERSION}")

//   if [ "${LATEST_VERSION_RESULTS}" -gt 0 ]; then
//     printf "There was a problem in confirming the version number.\nIgnoring the message and moving on\n"
//     printf "${LATEST_VERSION}\n\n"
//   else
//     if [ "${LATEST_VERSION}" == "${version}" ]; then
//       printf "You are using the latest version of upshift v${version}\n\n"
//     else
//       printf "Moving you to the latest version of upshift v${version}\n"
//       VERSION_UPGRADE=$(curl -fsSL https://raw.githubusercontent.com/leftshifters/upshift/master/upshift > upshift.temp && chmod +x upshift.temp && ./upshift.temp install)
//       printf "${VERSION_UPGRADE}\n"
//       rm upshift.temp
//     fi
//   fi

// }

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
			fmt.Println("We could not write the config file, the OS told us this <" + err.Error() + ">")
			return 1, false
		}
	}

	fmt.Println("We just added a config.toml to this folder!")
	return 0, false
}

// function AddConfig {

//   StartAction "AddConfig"

//   # If 'next' is false, exit
//   if [ ${next} == false ]; then
//     ShowPreviousFailed
//     return
//   fi

//   # Check if config.ci exists in
//   if [ -f "config.ci" ]; then
//     printf "A config file ${redColour}already exists${noColour}\n"
//   else
//     # Write the config file
//     echo debug=false >> config.ci
//     echo alwaysCleanBeforeBuild=true >> config.ci
//     echo alwaysUninstallOlderBuilds=true >> config.ci
//     echo package=\'\' >> config.ci
//     echo mainActivity=\'\' >> config.ci
//     echo gitRepositoryURL=\'\' >> config.ci
//     echo gitRepositoryBranch=\'\' >> config.ci
//     echo masterPassword=\'\' >> config.ci
//     echo projectName=\'\' >> config.ci
//     echo useWorkspace=false >> config.ci
//     echo scheme=\'\' >> config.ci
//     echo iPhone=\'\' >> config.ci
//     echo xcodeVersion=\'\' >> config.ci

//     printf "We just add a very basic confi.ci file in this folder.\n"
//   fi

// }
