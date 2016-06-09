package ios

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	c "upshift/colours"
	"upshift/command"
	"upshift/config"
	"upshift/utils"
)

var projectSettings map[string]string

func init() {

}

func IosBuild() (int, bool) {
	xcodeBuildSettings()
	projectName := projectSettings["PROJECT_NAME"]
	fullProductName := projectSettings["FULL_PRODUCT_NAME"]
	productBundleId := projectSettings["PRODUCT_BUNDLE_IDENTIFIER"]
	fmt.Println(projectName, fullProductName, productBundleId)

	getProjectPath()
	projectExtension := projectSettings["UP_PROJECT_EXTENSION"]
	fmt.Println(projectName + projectExtension)

	// isSimulatorRunning()

	return 1, true
}

// Is Simulator running
func isSimulatorRunning() bool {
	process, err := command.RunWithoutStdout([]string{"pgrep", "-f", "Simulator.app"}, "")
	if err != nil {
		return false
	}
	return true
}

// # Load Simulator
// OUTPUT=$(ps -aef | grep "Simulator.app" -c)
// if [ "${OUTPUT}" -gt 1 ]; then

// Get Project Path
func getProjectPath() {
	conf, _ := config.Get()

	useWorkspace := conf.IOS.UseWorkspace
	podFileFullPath, _ := filepath.Abs("Podfile")
	podfileExists := utils.FileExists(podFileFullPath)

	if useWorkspace || podfileExists {
		projectSettings["UP_PROJECT_TYPE"] = "workspace"
		projectSettings["UP_PROJECT_EXTENSION"] = ".xcworkspace"
	} else {
		projectSettings["UP_PROJECT_TYPE"] = "project"
		projectSettings["UP_PROJECT_EXTENSION"] = ".xcodeproj"
	}
}

// Read values from xcodebuild settings
func xcodeBuildSettings() error {
	projectSettings = make(map[string]string)
	version, err := command.RunWithoutStdout([]string{"xcodebuild", "-showBuildSettings"}, "")
	if err != nil {
		return errors.New(("We were unable to read xcodebuild settings\n" + err.Error()))
	}

	settingsRows := strings.Split(version, "\n")
	for _, row := range settingsRows {
		settingsKeys := strings.Split(row, "=")
		if len(settingsKeys) == 2 {
			key := strings.TrimSpace(settingsKeys[0])
			value := strings.TrimSpace(settingsKeys[1])
			projectSettings[key] = value
		}
	}
	return nil
}

//
// Choose the correct version of Xcode for the project
// It is usually defined in config.toml
//
func SetupXcode() (int, bool) {
	version, err := command.RunWithoutStdout([]string{"xcodebuild", "-version"}, "")
	if err != nil {
		utils.LogError("We were unable to get the Xcode version " + err.Error())
		return 1, true
	}

	var currentXcodeVersion string
	versionRows := strings.Split(version, "\n")
	for _, row := range versionRows {
		if strings.Contains(row, "Xcode") == true {
			currentXcodeVersion = strings.TrimSpace(strings.Trim(row, "Xcode"))
		}
	}

	fmt.Println("We are currently using Xcode-" + currentXcodeVersion)

	conf, err := config.Get()
	if err != nil {
		fmt.Println("We were unable to load the config file\n", err.Error())

		fmt.Println("You are currently on Xcode-" + currentXcodeVersion + " and the latest Xcode version is " + utils.GetDefaultXcodeVersion() + ". For now, we will continue using the version that you have right now")
		return 0, false
	}

	requiredXcodeVersion := conf.IOS.Xcode

	if requiredXcodeVersion == currentXcodeVersion {
		fmt.Println("You are on the correct version of Xcode")
		return 0, false
	}

	fmt.Println("Alright, so we will try and switch the Xcode version now to " + requiredXcodeVersion)

	if utils.FileExists("/Applications/Xcode-"+requiredXcodeVersion+".app/") == false {
		fmt.Println("It seems you don't have /Applications/Xcode-" + requiredXcodeVersion + ".app/\nWe expect XCode versions to be placed like this\n/Applications/Xcode-7.2.app\n/Applications/Xcode-7.3.app")
		return 1, true
	}

	var RootPassword string
	if utils.IsCI() == true {
		// We are on CI, we need to enter password programatically
		RootPassword, err = utils.GetRootPassword()
		if err != nil {
			utils.LogError(err.Error())
			return 1, true
		}
	}

	out, err := command.RunWithoutStdout([]string{"sudo", "-S", "xcode-select", "-switch", "/Applications/Xcode-" + requiredXcodeVersion + ".app/"}, RootPassword+"\n")
	if err != nil {
		fmt.Println("We couldn't switch Xcodes, you're going to be stuck with this one")
		return 1, true
	}
	fmt.Println(out)
	fmt.Println("We are now on the " + c.Underline + "Xcode-" + requiredXcodeVersion + c.Default)

	return 0, false
}
