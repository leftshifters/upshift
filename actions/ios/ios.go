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
		fmt.Println("We were unable to load the config file", err.Error())
		return 1, true
	}

	requiredXcodeVersion := conf.IOS.Xcode

	if requiredXcodeVersion == currentXcodeVersion {
		fmt.Println("You are on the correct version of Xcode")
		return 0, false
	}

	fmt.Println("Alright, so we will try and switch the Xcode version now to " + requiredXcodeVersion)

	if utils.FileExists("/Applications/Xcode-"+requiredXcodeVersion+".app/") == false {
		fmt.Println("It seems you don't have /Applications/Xcode-" + requiredXcodeVersion + ".app/")
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

// function XCodeVersion {

//   StartAction "XCodeVersion"

//   XCODE_VERSION=$(xcodebuild -version | grep "Xcode" | tr -d "Xcode ")

//   if [ "${XCODE_VERSION}" != "" ]; then

//     printf "We are currently using XCode ${blueColour}${XCODE_VERSION}${noColour}\n\n"

//     if [ "${xcodeVersion}" != "" ]; then

//       # Alright now we check if the versions match
//       if [ "${XCODE_VERSION}" != "${xcodeVersion}" ]; then
//         # Looks like the xcode version required and available do not match
//         # Check if this system has the XCode version required
//         # TODO : This will vary based on how you setup XCode, find out if there is a better way
//         printf "We expect XCode versions to be placed like this\n/Applications/Xcode-7.2.app\n/Applications/Xcode-7.3.app\n\n"
//         if [ -d "/Applications/Xcode-${xcodeVersion}.app/" ]; then
//           # Looks like this version is available on this machine

//           printf "${blueColour}Switching${noColour} Xcode versions\n\n"
//           if [ "${CI}" == true ]; then
//             # If we are on CI, expect masterPassword
//             if [ "${masterPassword}" != "" ]; then
//               echo -e "${masterPassword}" | sudo xcode-select -switch "/Applications/Xcode-${xcodeVersion}.app/"
//             else
//               ShowError
//               printf "Alright, so we need the sudo password to change the Xcode version.\nWould you be a sweetheart and add it\nto the ${blueColour}masterPassword${noColour} key in the config\n\n"
//               next=false
//               exit 1
//             fi
//           else
//             # If we are NOT on CI, don't expect password
//             sudo xcode-select -switch "/Applications/Xcode-${xcodeVersion}.app/"
//           fi

//           # Maye it's done, check and confirm
//           XCODE_VERSION=$(xcodebuild -version | grep "Xcode" | tr -d "Xcode ")

//           if [ "${XCODE_VERSION}" == "${xcodeVersion}" ]; then
//             printf "We are now using XCode ${blueColour}${XCODE_VERSION}${noColour}\n"
//           else
//             ShowError
//             printf "Something went wrong. Maybe we messed up big time or the password that you entered was wrong.\n\n"
//             next=false
//           fi
//         else
//           ShowError
//           printf "We do not have XCode ${xcodeVersion} on this machine. About time to get it.\n\n"
//           next=false
//         fi
//       fi

//     else
//       printf "Your config shows that you are not picky about the\n${blueColour}xcodeVersion${noColour}. We will use whatever is available!\n\n"
//     fi

//   else
//     ShowError
//     printf "We do not have ${redColour}XCode${noColour} installed on this machine. About time to get it.\n\n"
//     next=false
//   fi
