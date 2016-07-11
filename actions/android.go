package actions

import (
	// "errors"
	"fmt"
	"path/filepath"
	"strings"
	"upshift/basher"
	"upshift/command"
	"upshift/config"
	"upshift/utils"
)

// UpgradeAndroid : Upgrade the android sdk
func UpgradeAndroid() int {
	var b basher.Basher
	logPath, _ := filepath.Abs(".upshift/logs/android-sdk-upgrade.log")
	_, err := b.Run("AndroidUpgradeSDK", []string{logPath})
	if err != nil {
		utils.LogError("We could not start upgrading android.\n" + err.Error())
		return 1
	}

	return 0
}

// SetupAndroid : Install the android sdk
func SetupAndroid() int {
	var b basher.Basher
	logPath, _ := filepath.Abs(".upshift/logs/android-sdk-upgrade.log")
	_, err := b.Run("AndroidInstallSDK", []string{logPath})
	if err != nil {
		utils.LogError("We could not start upgrading android.\n" + err.Error())
		return 1
	}

	return 0
}

// AndroidBuild : Build the android project
func AndroidBuild() int {

	var b basher.Basher
	var gradle Gradle
	conf := config.Get()
	cleanOnStart := conf.Settings.CleanBeforeBuild

	if cleanOnStart == true {
		fmt.Println("Let's clean the project before starting")
		_, err := gradle.Clean(".upshift/logs/android-clean.log")
		if err != nil {
			utils.LogError("We could not clean your project. It's really dirty\n" + err.Error())
			return 1
		}
	}

	launchEmulator()

	// Check if there are any connected devices
	if len(devicesConnected()) > 0 {
		// Delete older builds if they are installed
		fmt.Println("Removing older builds from connected devices")
		// #TODO : Change to Gradle.Uninstall
		logPath, _ := filepath.Abs(".upshift/logs/android-uninstall.log")
		_, err := b.Run("AndroidUninstall", []string{logPath})
		if err != nil {
			utils.LogError("We could not uninstall the older binaries.\n" + err.Error())
			// Don't return on this, we don't even know why this fails, maybe because a device isn't connected
		}
		// #END
	}

	// #TODO : Replace with Gradle.Lint
	fmt.Println("Before we build, we need to lint")
	utils.LogMessage("$ ./gradlew lint")
	logPath, _ := filepath.Abs(".upshift/logs/android-lint.log")
	_, err := b.Run("AndroidLint", []string{logPath})
	if err != nil {
		utils.LogError("We could not start lintin your project.\n" + err.Error())
		return 1
	}

	tailData, err := utils.FileTail(logPath, 500)
	if err != nil {
		utils.LogError("It seems we couldn't read the output. Here's what happened\n" + err.Error())
		return 1
	}

	if strings.Contains(tailData, "BUILD SUCCESSFUL") == false {
		utils.LogError("Something went wrong while linting, you need to look at this.")
		return 1
	}
	// #END

	// #TODO : Replace with Gradle.Assemble
	fmt.Println("Okay, so lets build Debug and install it on a emulator")
	utils.LogMessage("$ ./gradlew assemble --stacktrace")
	logPath, _ = filepath.Abs(".upshift/logs/android-assemble.log")
	_, err = b.Run("AndroidAssemble", []string{logPath})
	if err != nil {
		utils.LogError("We could not start building your project.\n" + err.Error())
		return 1
	}

	tailData, err = utils.FileTail(logPath, 500)
	if err != nil {
		utils.LogError("It seems we couldn't read the output. Here's what happened\n" + err.Error())
		return 1
	}

	if strings.Contains(tailData, "BUILD SUCCESSFUL") == false {
		utils.LogError("Something went wrong while building, you need to look at this.")
		return 1
	}
	// #END

	return 0
}

func launchEmulator() bool {

	// 1. Check if any devices are connected, if yes, use one of those
	// 2. If nothing so far, see if any avds are listed and start the first one
	// 3. If still nothing, create an avd and launch it
	// - N E V E R - G I V E - U P - (Alright, gave up, not creating avds, too much pain)

	// avdToStart := ""
	// avdToStart = "Nexus_5_API_22"

	devices := devicesConnected()
	if len(devices) > 0 {
		return true
	}

	// avds := avdsAvailable()
	// if len(avds) == 0 {
	// No AVDs found, create one
	// to view a list of available avds you can create, run 'android list targets'
	// we prefer to use the latest one possible, and we only feed in the required fields
	// look for ones with ABIs
	// android create avd --target android-23 --name "Google Inc.:Google APIs:22" -b "google_apis/x86_64"
	// return false
	// }

	// fmt.Println("Time to load up the emulator " + c.Blue + avds[0] + c.Default)
	// logPath, _ := filepath.Abs(".upshift/logs/android-emulator.log")
	// _, err := b.Run("AndroidLaunchEmulator", []string{avds[0], logPath})
	// if err != nil {
	// utils.LogError("We could not start loading up the emulator.\n" + err.Error())
	// return false
	// }

	return true
}

func devicesConnected() []string {
	out, err := command.Run([]string{"adb", "devices"}, "")
	if err != nil {
		fmt.Println("We couldn't start finding devices\n" + err.Error())
		return []string{}
	}

	devices := utils.CreateList(out, []string{"List of devices attached", "daemon not running. starting it now on port", "daemon started successfully", "offline"})

	return devices
}

//     printf "\n\n${greenColour}Super${noColour}! The build was fine.\n"
//     # Check if package is empty
//     if [ "${package}" != "" ];then
//       if [ "${mainActivity}" != "" ]; then
//         printf "Starting activity ${blueColour}${mainActivity}${noColour} in package ${blueColour}${package}${noColour}\n"

//         # Start the activity and package
//         adb shell am start -n ${package}/${package}.${mainActivity}

//         # Tell the user everything is nice and easy
//         printf "\nAlright, the build was ${greenColour}successful${noColour} 🍺\n\n"
//       else
//         # The mainActivity is empty it seems
//         printf "Alright, the build was ${greenColour}successful${noColour}, but there was no ${blueColour}mainActivity${noColour} defined, so couldn't start it automatically 🍺\n\n"
//       fi
//     else
//       # The package is empty it seems
//       printf "Alright, the build was ${greenColour}successful${noColour}, but there was no ${blueColour}package${noColour} defined, so couldn't start it automatically 🍺\n\n"
//     fi
//   fi
// }
