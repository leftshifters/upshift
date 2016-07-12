package actions

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/leftshifters/upshift/basher"
	"github.com/leftshifters/upshift/colours"
	"github.com/leftshifters/upshift/config"
)

// Upshift : Construct to handle all upshift related things
type Upshift struct {
}

// Upgrade : upgrade the upshift binary to the latest one
func (u *Upshift) Upgrade(beta string) error {
	conf := config.Get()

	// Get Latest version
	version, err := u.LatestVersion()
	if err != nil {
		return err
	}

	if version == conf.Settings.AppVersion {
		return nil
	}

	var b basher.Basher
	_, err = b.Run("UpgradeScript", []string{beta})
	if err != nil {
		return err
	}

	fmt.Println("The new version of awesomeness is v", version)
	return nil
}

// LatestVersion : get the latest version of upshift from the server
func (u *Upshift) LatestVersion() (string, error) {
	versionURL := "https://raw.githubusercontent.com/leftshifters/upshift/master/release"

	// Get the latest version from the server
	resp, err := http.Get(versionURL)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}

	// Read the response received
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Trim and send the version back
	return strings.TrimSpace(string(body)), nil
}

// ShowVersion : shows the version of upshift on the command line
func (u *Upshift) ShowVersion() {
	conf := config.Get()
	fmt.Println(conf.Settings.AppVersion)
}

// ShowHelp : Show help, so that the user knows what to do
func (u *Upshift) ShowHelp() {
	fmt.Println("\nUPSHIFT(1)               Upshift Commands Manual               UPSHIFT(1)")
	fmt.Println(colours.Bold + "\nNAME" + colours.Default)
	fmt.Println("\tupshift -- the creative mobile app builder")
	fmt.Println(colours.Bold + "\nSYNOPSIS" + colours.Default)
	fmt.Println("\tupshift " + colours.Underline + "job" + colours.Default + " " + colours.Underline + "action" + colours.Default)
	fmt.Println(colours.Bold + "\nDESCRIPTION" + colours.Default)
	fmt.Println("\tThis tool helps you run, build, test and deploy your iOS and Android\n\tapps while you dream about the next big thing")

	fmt.Println(colours.Bold + "\nOPTIONS (job queues)" + colours.Default)
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

	fmt.Println(colours.Bold + "\nOPTIONS (specific actions)" + colours.Default)
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

	fmt.Println(colours.Bold + "\nCOMPATIBILITY" + colours.Default)
	fmt.Println("\tWe've only tested this on Mac OSX, Linux and Docker. If you're on\n\twindows, you should switch operating systems because nobody can help\n\tyou there.")
	fmt.Println("\nLeftshift Technologies           Made with ❤️  in India                " + colours.Underline + "https://leftshift.io\n" + colours.Default)
}
