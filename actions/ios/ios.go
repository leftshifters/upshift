package ios

import (
	"fmt"
	"strings"
	c "upshift/colours"
	"upshift/command"
	"upshift/config"
	"upshift/utils"
)

func init() {

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

	if utils.IsCI() == true {
		// We are on CI, we need to enter password programatically

		RootPassword, err := utils.GetRootPassword()
		if err != nil {
			utils.LogError(err.Error())
			return 1, true
		}

		out, err := command.RunWithoutStdout([]string{"sudo", "-S", "xcode-select", "-switch", "/Applications/Xcode-" + requiredXcodeVersion + ".app/"}, RootPassword+"\n")
		if err != nil {
			fmt.Println("We couldn't switch Xcodes, you're going to be stuck with this one")
			return 1, true
		}

		fmt.Println(out)

	} else {
		// We are not on CI, ask the user to enter the password
		out, err := command.RunWithoutStdout([]string{"sudo", "xcode-select", "-switch", "/Applications/Xcode-" + requiredXcodeVersion + ".app/"}, "")
		if err != nil {
			fmt.Println("We couldn't switch Xcodes, you're going to be stuck with this one")
			return 1, true
		}

		fmt.Println("We are now on the " + c.Underline + "Xcode-" + requiredXcodeVersion + c.Default)

		fmt.Println(out)
	}

	return 0, false
}
