package setup

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"upshift/basher"
	c "upshift/colours"
	"upshift/command"
	"upshift/config"
	g "upshift/global"
	"upshift/utils"
)

//
// Init function, doesn't do shit
//
func init() {

}

//
// Show the version of the current app
//
func ShowVersion() (int, bool) {
	fmt.Println(utils.GetAppVersion())
	return 0, false
}

//
// Setup Gradle Wrapper
//
func SetupGradleW() (int, bool) {
	// Run gradle -v to figure out if it is install
	_, err := command.RunWithoutStdout([]string{"gradle", "-v"}, "")
	if err != nil {
		utils.LogError("Gradle itself is not installed, can't install wrapper.")
		return 1, true
	}

	gradlewPath, _ := filepath.Abs("./gradlew")
	gradlewExist := utils.FileExists(gradlewPath)

	if gradlewExist == true {
		fmt.Println("You already have gradle wrapper installed, moving on.")
		return 0, false
	}

	// So, gradle is installed, just need to install wrapper [SetupGradleW]
	// I won't touch anything to do with gradle and pipes with a ten foot pole, so this goes to basher
	utils.LogMessage("$ gradle wraper")
	status, err := basher.Run("SetupGradleW", []string{})
	if err != nil {
		utils.LogError("We couldn't initialise gradle wrapper\n" + err.Error())
		return status, true
	}

	fmt.Println("Gradle wrapper has been successfully setup")
	return 0, false
}

//
// Install cocoapods
//
func InstallPods() (int, bool) {
	podsPath, _ := filepath.Abs("Podfile")
	podsExist := utils.FileExists(podsPath)

	if podsExist == false {
		fmt.Println("It looks like this project doesn't use pods")
		return 0, false
	}

	// So pods exist, damn
	fmt.Println("It looks like this project uses pods, let me try and set them up")

	err := runPodRepoUpdate()
	if err != nil {
		utils.LogError(err.Error())
		return 1, true
	}

	utils.LogMessage("$ pod install")
	podInstallLogFullPath, _ := filepath.Abs(".upshift/logs/pod-install.log")
	status, err := basher.Run("PodInstall", []string{podInstallLogFullPath})
	if err != nil {
		utils.LogError("We couldn't initialise pods\n" + err.Error())
		return status, true
	}

	// Read the last 500 bytes from the whole message, we just want to what happened at the end
	tailData, err := utils.ReadTailIfFileExists(podInstallLogFullPath, 500)
	if err != nil {
		utils.LogError("It seems we couldn't read the output. Here's what happened\n" + err.Error())
		return status, true
	}

	if strings.Contains(tailData, "fatal") == true || strings.Contains(tailData, "error") == true || strings.Contains(tailData, "Invalid") == true {
		utils.LogError("Something went wrong with installing pods, you need to look at this.")
		return 1, true
	}

	fmt.Println("We were able to successfully setup cocoapods, moving on")
	return 0, false

}

func runPodRepoUpdate() error {

	// Check if it's been a while since we ran `pod repo update` and if yes, run it
	globalConf, _ := g.Get()
	difference := int32(time.Now().Unix()) - globalConf.AndroidSDKUpdatedTime

	// Run this if last update was never done or more than a month ago
	if globalConf.AndroidSDKUpdatedTime == 0 || difference > 3600*24*30 {
		// This means we have never come here.
		utils.LogMessage("$ pod repo update --verbose")
		podInstallLogFullPath, _ := filepath.Abs(".upshift/logs/pod-repo-update.log")
		_, err := basher.Run("PodRepoUpdate", []string{podInstallLogFullPath})
		if err != nil {
			return err
		}

		// Read the last 500 bytes from the whole message, we just want to what happened at the end
		tailData, err := utils.ReadTailIfFileExists(podInstallLogFullPath, 500)
		if err != nil {
			return err
		}

		if strings.Contains(tailData, "error") == true {
			return err
		}

		// When an upgrade is available, they say
		// CocoaPods 1.0.1 is available.
		// To update use: `sudo gem install cocoapods`
		// Until we reach version 1.0 the features of CocoaPods can and will change.
		// We strongly recommend that you use the latest version at all times.
		if strings.Contains(tailData, "sudo gem install cocoapods") == true {
			// This means that an update is available, run cocoapods update
			status, _ := SetupPods(true)
			if status > 0 {
				return errors.New("We couldn't update to the new version of cocoapods")
			}
			fmt.Println("Updated cocoapods to the latest version")
		}

		globalConf.AndroidSDKUpdatedTime = int32(time.Now().Unix())
		g.Set(globalConf)
	}
	return nil
}

//
// Initialize submodules in a project
//
func GitSubmodules() (int, bool) {
	submodulePath, _ := filepath.Abs(".gitmodules")
	submodulesExist := utils.FileExists(submodulePath)

	if submodulesExist == false {
		fmt.Println("It looks like this project doesn't use submodules")
		return 0, false
	}

	// So git submodules exist
	fmt.Println("It looks like this project uses git submodules, let me try and set them up")

	utils.LogMessage("$ git submodule init")
	initLogFileFullPath, _ := filepath.Abs(".upshift/logs/git-submodule-init.log")
	status, err := basher.Run("GitSubmoduleInit", []string{initLogFileFullPath})
	if err != nil {
		utils.LogError("We couldn't initialise submodules\n" + err.Error())
		return status, true
	}

	utils.LogMessage("$ git submodule update")
	updateLogFileFullPath, _ := filepath.Abs(".upshift/logs/git-submodule-update.log")
	status, err = basher.Run("GitSubmoduleUpdate", []string{updateLogFileFullPath})
	if err != nil {
		utils.LogError("We couldn't update submodules\n" + err.Error())
		return status, true
	}

	// Read the last 500 bytes from the whole message, we just want to what happened at the end
	tailData, err := utils.ReadTailIfFileExists(updateLogFileFullPath, 500)
	if err != nil {
		utils.LogError("It seems we couldn't read the output. Here's what happened\n" + err.Error())
		return status, true
	}

	if strings.Contains(tailData, "fatal:") == true || strings.Contains(tailData, "error:") == true {
		utils.LogError("Something went wrong with submodule update, you need to look at this.")
		return 1, true
	}

	fmt.Println("We were able to successfully setup submodules, moving on")
	return 0, false
}

//
// Do a git pull on the project based on the defined remote and the branch the user is currently on
//
func GitPull() (int, bool) {

	// If you are running on a CI, we don't need to worry about this, just skip and take up the next thing
	if utils.IsCI() == true {
		fmt.Println("It seems you're running this on a CI, so we are going to skip the git pull, it's the CI's job to give me the latest code")
		return 0, false
	}

	// Find out which repo and branch are they on
	out, err := command.RunWithoutStdout([]string{"git", "status"}, "")
	if err != nil {
		utils.LogError("Either this is not a git repository, or you don't even have git installed.")
		return 1, true
	}

	// Read the first row of git sttus which says 'on branch xyz'
	gitStatusOutputRows := strings.Split(out, "\n")
	var firstRow string
	if len(gitStatusOutputRows) > 0 {
		firstRow = gitStatusOutputRows[0]
	} else {
		utils.LogError("You are probably not in a git repository. Quit messing around.")
		return 1, true
	}

	// Alright find the correct branch and show it to the user
	currentBranch := strings.TrimSpace(strings.Replace(firstRow, "On branch ", "", 1))
	fmt.Println("We suspect that you are on branch " + c.Blue + currentBranch + c.Default)

	// Check if the user has one or more remotes
	// If they have one, just use it
	// If they have more
	// 		Show them the remotes, and ask them to specify it in config.toml

	// Find out which repo and branch are they on
	out, err = command.RunWithoutStdout([]string{"git", "remote"}, "")
	if err != nil {
		utils.LogError("Either this is not a git repository, or you don't even have git installed.")
		return 1, true
	}

	var currentRemote string
	gitRemoteOutputRows := strings.Split(strings.TrimSpace(out), "\n")
	switch len(gitRemoteOutputRows) {
	case 0:
		utils.LogError("Um, you have no remotes, I really don't know what to do. I'm going to kill myself")
		return 1, true
	case 1:
		currentRemote = strings.TrimSpace(gitRemoteOutputRows[0])
		fmt.Println("And we suspect that you are using the " + c.Blue + currentRemote + c.Default + " remote")
	default:
		// This means that the user has multiple remotes, read the config to see if they have mentioned a remote there
		conf, _ := config.Get()

		if conf.Build.GitRepoRemote != "" {
			// Alright, so they have defined a remote, let's check if it exits in our list of remotes
			for _, row := range gitRemoteOutputRows {
				if strings.TrimSpace(conf.Build.GitRepoRemote) == strings.TrimSpace(row) {
					currentRemote = row
				}
			}

			// Didn't find their remote in git remotes, tell them so
			if currentRemote == "" {
				utils.LogError("Here's a strange problem. Your config says you want to use\nthe " + conf.Build.GitRepoRemote + " remote, but sadly we " + c.Underline + "couldn't find that remote" + c.Default + " for\nthis repo. All we found was " + strings.Join(gitRemoteOutputRows, ", "))
				return 1, true
			} else {
				fmt.Println("And your config tells me you want to read from the remote " + c.Blue + currentRemote + c.Default)
			}
		} else {
			// They haven't defined a remote in config, screw them
			utils.LogError("You have more than one remote. In your config.toml you need to specify which remote to pull from\nThe following remotes are avilable " + strings.TrimSpace(out))
			return 1, true
		}
	}

	utils.LogMessage("$ git pull " + currentRemote + " " + currentBranch)
	// Now we have both, a remote and a branch, let's pull
	logFileFullPath, _ := filepath.Abs(".upshift/logs/git-pull.log")
	status, err := basher.Run("GitPull", []string{currentRemote, currentBranch, logFileFullPath})
	// status, err := basher.Run("RunSingleCommand", []string{"$(git pull " + currentRemote + " " + currentBranch + " 2>&1 | tee \"" + logFileFullPath + "\")"})
	if err != nil {
		utils.LogError("We couldn't pull the branch " + currentBranch + " from the remote " + currentRemote + " - \n" + err.Error())
		return status, true
	}

	// Read the last 500 bytes from the whole message, we just want to what happened at the end
	tailData, err := utils.ReadTailIfFileExists(logFileFullPath, 500)
	if err != nil {
		utils.LogError("It seems we couldn't read the output. Here's what happened\n" + err.Error())
		return status, true
	}

	if strings.Contains(tailData, "fatal:") == true || strings.Contains(tailData, "error:") == true {
		utils.LogError("Something went wrong with the pull, you need to look at this.")
		return 1, true
	}

	fmt.Println("We were able to pull the latest code, awesome")
	return 0, false
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
	fmt.Println("\tupshift setup export.plist\n\t\tto setup a sample .private/export.plist in your project\n")
	fmt.Println("\tupshift install\n\t\tto install this binary for the first time\n")
	fmt.Println("\tupshift -v\n\t\tto view the version number\n")

	fmt.Println(c.Bold + "\nOPTIONS (specific actions)" + c.Default)
	fmt.Println("\tWe combine actions like these to create the jobs above, you should ideally\n\tbe running jobs not actions\n")
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

	return 0, false
}

//
// If fastlane.tools if is not setup, we go ahead and do it
//
func SetupFastlane(force bool) (int, bool) {
	return SetupGem("fastlane", "fastlane", force)
}

//
// If cocoapods is not setup, we go ahead and do it
//
func SetupPods(force bool) (int, bool) {
	return SetupGem("cocoapods", "pod", force)
}

//
// If Xcpretty is not setup, we go ahead and do it
// It formats the output from xcode so that you can make sense of what is going wrong
//
func SetupXcpretty() (int, bool) {
	return SetupGem("xcpretty", "xcpretty", false)
}

//
// General script to setup a gem
//
func SetupGem(gem string, gemName string, force bool) (int, bool) {
	// Check which version of the gem was installed
	// force will be true when you want to force running sudo install gem irrespective of if it is installed or not
	version, err := command.RunWithoutStdout([]string{gemName, "--version"}, "")
	if force == false {
		if err == nil {
			// Remove the name of the gem if it is part of the version string
			version = strings.Replace(version, gemName, "", 1)
			// Now trim whatever is left
			version = strings.TrimSpace(version)
			fmt.Println(gem + " is pretty much setup on this system. You are on version " + version)
			return 0, false
		}
	}

	// Check if the command was not found
	var errorString string
	if err != nil {
		errorString = err.Error()
	}

	if strings.Contains(errorString, "executable file not found") == true || force == true {
		// Alright, so the gem was not found, go ahead and install it
		fmt.Println("Latest " + gem + " was not found, installing it")

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

		status, err := basher.Run("SetupGem", []string{gem, strconv.FormatBool(utils.IsCI()), RootPassword})
		if err != nil {
			utils.LogError("We couldn't install " + gem + "\n" + err.Error())
			return status, false
		}

		fmt.Println(gem + " has been setup on your machine. Have fun.")
		return 0, false
	}

	fmt.Println("There was a problem installing " + gem + "\n" + err.Error())
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
			return 1, false
		}
	}

	fmt.Println("We just added a config.toml to this folder!")
	return 0, false
}
