package actions

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
	"upshift/basher"
	c "upshift/colours"
	"upshift/command"
	"upshift/config"
	"upshift/utils"
)

var projectSettings map[string]string

func IosUploadBuild() (int, bool) {
	err := uploadBuildToItunes()
	if err != nil {
		utils.LogError(err.Error())
		return 1, true
	}

	return 0, false
}

func IosCreateApp() (int, bool) {
	err := createAppOniTunes()
	if err != nil {
		utils.LogError(err.Error())
		return 1, true
	}

	return 0, false
}

func IosExportIPA() (int, bool) {
	// Try the build now
	projectName := projectSettings["PROJECT_NAME"]

	err := exportIPAForIOS(projectName)
	if err != nil {
		utils.LogError(err.Error())
		return 1, true
	}

	return 0, false
}

func IosArchive() (int, bool) {
	projectType := projectSettings["UP_PROJECT_TYPE"]
	projectName := projectSettings["PROJECT_NAME"]
	projectExtension := projectSettings["UP_PROJECT_EXTENSION"]
	projectPath := projectName + projectExtension
	projectScheme := projectSettings["UP_PROJECT_SCHEME"]

	err := archiveForIOS(projectType, projectPath, projectScheme, projectName)
	if err != nil {
		utils.LogError(err.Error())
		return 1, true
	}

	return 0, false
}

func IosCertificates() (int, bool) {
	err := installCertificates()
	if err != nil {
		utils.LogError(err.Error())
		return 1, true
	}

	return 0, false
}

func IosProvisioning() (int, bool) {
	err := addProvisioningProfiles()
	if err != nil {
		utils.LogError(err.Error())
		return 1, true
	}

	return 0, false
}

func IosDeploySimulator() (int, bool) {
	projectName := projectSettings["PROJECT_NAME"]
	projectBundleIdentifier := projectSettings["PRODUCT_BUNDLE_IDENTIFIER"]

	err := deployToSimulator(projectName, projectBundleIdentifier)
	if err != nil {
		utils.LogError(err.Error())
		return 1, true
	}

	return 0, false
}

func IosTest() (int, bool) {

	projectType := projectSettings["UP_PROJECT_TYPE"]
	projectName := projectSettings["PROJECT_NAME"]
	projectExtension := projectSettings["UP_PROJECT_EXTENSION"]
	projectPath := projectName + projectExtension
	projectScheme := projectSettings["UP_PROJECT_SCHEME"]
	projectDevice := projectSettings["UP_SIMULATOR_IPHONE"]

	err := testForIOS(projectType, projectPath, projectScheme, projectDevice)
	if err != nil {
		utils.LogError(err.Error())
		return 1, true
	}

	return 0, false
}

func IosBuild() (int, bool) {

	// Try the build now
	projectType := projectSettings["UP_PROJECT_TYPE"]
	projectName := projectSettings["PROJECT_NAME"]
	projectExtension := projectSettings["UP_PROJECT_EXTENSION"]
	projectPath := projectName + projectExtension
	projectScheme := projectSettings["UP_PROJECT_SCHEME"]
	projectDevice := projectSettings["UP_SIMULATOR_IPHONE"]

	err := compileForIOS(projectType, projectPath, projectScheme, projectDevice)
	if err != nil {
		utils.LogError(err.Error())
		return 1, true
	}

	return 0, false
}

func IosPrepare() (int, bool) {
	fmt.Println("We will now try and load your xcode settings")
	err := xcodeBuildSettings() // Gets PROJECT_NAME, FULL_PRODUCT_NAME, PRODUCT_BUNDLE_IDENTIFIER amongst others in projectSettings
	if err != nil {
		utils.LogError("We could not read your xcode settings, it this an iOS repo?\n" + err.Error())
		return 1, true
	}

	getProjectPath()             // Gets UP_PROJECT_EXTENSION, UP_PROJECT_TYPE in projectSettings
	findXcodeAndOSForSimulator() // Gets UP_XCODE_VERSION, UP_SIMULATOR_IPHONE, UP_SIMULATOR_IPHONEOS in projectSettings

	// Create the device name for testing
	deviceForSimulator := projectSettings["UP_SIMULATOR_IPHONE"] + " (" + projectSettings["UP_SIMULATOR_IPHONEOS"] + ")"

	// Check if this device is available on this machine
	fmt.Println("Checking if this machine can run the simulator for " + deviceForSimulator)
	if findIfDeviceIsAvailable(deviceForSimulator) == true {
		// If simulator is running, don't do anything, if not start it
		if isSimulatorRunning() == true {
			fmt.Println("Oh, it seems your simulator is already running, we can't figure out which version it is though")
		} else {
			fmt.Println("It sure is available, lets start up the simulator")
			startSimulator(deviceForSimulator)
			fmt.Println("You should see a simulator loading in the background")
		}
	} else {
		// Device not found, can't start the simulator
		return 1, true
	}

	// Either simulator should be running by now or should have been started

	// Find available schemes
	err = findAvailableSchemes()
	if err != nil {
		utils.LogError(err.Error())
		return 1, true
	}

	// Increment the build number
	err = incrementBuildNumber()
	if err != nil {
		utils.LogError(err.Error())
		return 1, true
	}

	return 0, false
}

func incrementBuildNumber() error {
	var b basher.Basher
	projectName := projectSettings["PROJECT_NAME"]
	_, err := b.Run("IOSIncrementBuildNumber", []string{projectName})
	if err != nil {
		return errors.New("We couldn't incremenet the build number")
	}
	return nil
}

func deployToSimulator(projectName string, projectBundleIdentifier string) error {
	builtFile := ""
	debugFile, _ := filepath.Abs(".upshift/build/Build/Products/Debug-iphonesimulator/" + projectName + ".app")
	releaseFile, _ := filepath.Abs(".upshift/build/Build/Products/Release-iphonesimulator/" + projectName + ".app")
	fileExits := utils.FileExists(debugFile)
	if fileExits == false {
		fileExits = utils.FileExists(releaseFile)
		if fileExits == false {
			return errors.New("It seems you haven't build for the simulator yet. Please try " + c.Red + "upshift ios build" + c.Default + " first")
		} else {
			// we found the release file
			builtFile = releaseFile
		}
	} else {
		// we found the debug file
		builtFile = debugFile
	}

	// If file exists, push it to the simulator
	fmt.Println("Installing the app in the simulator")
	_, err := command.Run([]string{"xcrun", "simctl", "install", "booted", builtFile}, "")
	if err != nil {
		return errors.New("We could not deploy the app to the simulator\n" + err.Error())
	}

	// Start the app
	fmt.Println("Starting the app in the simulator")
	_, err = command.Run([]string{"xcrun", "simctl", "launch", "booted", projectBundleIdentifier}, "")
	if err != nil {
		return errors.New("We could not deploy the app to the simulator\n" + err.Error())
	}

	return nil
}

// xcrun simctl launch booted "${productBundleIdentifier}"

func SetupExportPlist() (int, bool) {
	configExits := utils.FileExists(".private/export.plist")
	if configExits == true {
		fmt.Println("It looks like .private/export.plist is already here, skipping this step")
		return 1, false
	} else {
		// export.Plist does not exist
		// Create a new export.plist in this directory in .private

		sampleExportPlist := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
        <key>method</key>
        <string>ad-hoc</string>
        <key>uploadSymbols</key>
        <true/>
	<key>uploadBitcode</key>
	<true/>
</dict>
</plist>"`

		exportPlistBytes := []byte(sampleExportPlist)

		err := ioutil.WriteFile(".private/export.plist", exportPlistBytes, 0644)
		if err != nil {
			utils.LogError("We could not write the .private/export.plist file\n" + err.Error())
			return 1, true
		}
	}

	fmt.Println("We just added a sample file to .private/export.plist!")
	return 0, false

}

//
// Setup provisioning profiles
// This will probably be run once in a while
//
func SetupProfiles() (int, bool) {

	var b basher.Basher
	developerAccounts, err := getDeveloperAccounts()
	if err != nil {
		utils.LogError(err.Error())
		return 1, false
	}

	for _, email := range developerAccounts {
		email = strings.TrimSpace(email)

		if email != "" {
			fmt.Println("Trying to repair your provisioning profiles and installing new and fixed ones for " + email)
			_, err := b.Run("FetchAndRepairProvisioningProfiles", []string{email})
			if err != nil {
				utils.LogError("We couldn't fix and install your provisioning profiles")
				return 1, false
			}
		}
	}
	return 0, true
}

func getDeveloperAccounts() ([]string, error) {

	var developerAccounts []string

	conf := config.Get()

	if conf.Settings.IOSDeveloperAccount == "" {
		return developerAccounts, errors.New("You should define the emails of the developer account in the global config")
	}

	developerAccounts = strings.SplitN(conf.Settings.IOSDeveloperAccount, ",", -1)
	return developerAccounts, nil
}

func installCertificates() error {
	var b basher.Basher
	utils.LogMessage("Checking for apple.cer, distribution.p12 and distribution.cer in .private")

	basePath, _ := filepath.Abs(".private")

	// First check if the certificates are added to .private folder
	appleCer, _ := filepath.Abs(".private/apple.cer")
	distributionCer, _ := filepath.Abs(".private/distribution.cer")
	distributionP12, _ := filepath.Abs(".private/distribution.p12")

	// If they exist, install them
	appleCerExists := utils.FileExists(appleCer)
	distributionCerExists := utils.FileExists(distributionCer)
	distributionP12Exists := utils.FileExists(distributionP12)

	if !(appleCerExists && distributionCerExists && distributionP12Exists) {
		// If they don't exist, check global config to see if they have them
		conf := config.Get()
		globalBasePath := conf.Settings.IOSCertificatePath

		if globalBasePath == "" {
			return errors.New("The certificates don't exist in both .private and global conf")
		}

		basePath = globalBasePath

		appleCer, _ = filepath.Abs(globalBasePath + "/apple.cer")
		distributionCer, _ = filepath.Abs(globalBasePath + "/distribution.cer")
		distributionP12, _ = filepath.Abs(globalBasePath + "/distribution.p12")

		// If they exist, install them
		appleCerExists = utils.FileExists(appleCer)
		distributionCerExists = utils.FileExists(distributionCer)
		distributionP12Exists = utils.FileExists(distributionP12)

		if !(appleCerExists && distributionCerExists && distributionP12Exists) {
			return errors.New("All the certificates don't exist in both .private and global conf")
		}
		// If they don't, crash and burn
	}

	b.Run("InstallCertificates", []string{basePath})
	// #TODO Ignore the error here. Even if it says that certificates are already installed, it gets treated like an error
	// if err != nil {
	// 	return errors.New("The certicates could not be installed!")
	// }

	fmt.Println("The certificates were successfully installed")
	return nil
}

func uploadBuildToItunes() error {
	var b basher.Basher
	utils.LogMessage("Upload the IPA on iTunesConnect")

	// Get the username which will need to login
	// Highest priority to local config
	conf := config.Get()
	developerAccount := conf.Settings.IOSDeveloperAccount

	if developerAccount == "" {
		// Second priority to machine config
		developerAccounts, err := getDeveloperAccounts()
		if err != nil {
			return errors.New(err.Error())
		}

		// If there is only one in machine config, then we can use it, if there are more you need to add it to local config
		if len(developerAccounts) > 1 {
			return errors.New("There are too many developer accounts in the machine config, we either need one in machine config or one in local config")
		}

		if len(developerAccounts) == 1 {
			developerAccount = developerAccounts[0]
		}
	}

	projectScheme := projectSettings["UP_PROJECT_SCHEME"]
	projectName := projectSettings["PROJECT_NAME"]

	// Add SwitSources if required - AddSwiftSources
	status, err := b.Run("AddSwiftSources", []string{projectName, projectScheme})
	fmt.Println("status", status)
	if err != nil {
		fmt.Println("err", err.Error())
		return errors.New("We could not add SwiftSources to the IPA")
	}

	status, err = b.Run("UploadIPAoniTunes", []string{developerAccount, ".upshift/" + projectScheme + ".ipa"})
	if err != nil {
		return errors.New("We could not upload the IPA on iTunes")
	}

	fmt.Println("We have successfully uploaded this IPA on iTunes, it's all yours now")
	return nil
}

func createAppOniTunes() error {
	var b basher.Basher
	utils.LogMessage("Create an app on iTunesConnect if it doesn't exist")

	// Get the username which will need to login
	// Highest priority to local config
	conf := config.Get()
	developerAccount := conf.Settings.IOSDeveloperAccount

	if developerAccount == "" {
		// Second priority to machine config
		developerAccounts, err := getDeveloperAccounts()
		if err != nil {
			return errors.New(err.Error())
		}

		// If there is only one in machine config, then we can use it, if there are more you need to add it to local config
		if len(developerAccounts) > 1 {
			return errors.New("There are too many developer accounts in the machine config, we either need one in machine config or one in local config")
		}

		if len(developerAccounts) == 1 {
			developerAccount = developerAccounts[0]
		}
	}

	// Get the bundle identifier for this project
	projectBundleIdentifier := projectSettings["PRODUCT_BUNDLE_IDENTIFIER"]

	// Get the name of the project
	// And add your shit to it
	projectName := projectSettings["PROJECT_NAME"] + " Beta by Upshift"

	_, err := b.Run("CreateAppOnItunes", []string{developerAccount, projectBundleIdentifier, projectName})
	if err != nil {
		return errors.New("We could not create the app on iTunes\n" + err.Error())
	}

	fmt.Println("We have successfully added this app on iTunes, woohoo")
	return nil
}

func addProvisioningProfiles() error {
	var b basher.Basher
	utils.LogMessage("We will now try to find the provisioning profile")

	// Get the username which will need to login
	// Highest priority to local config
	conf := config.Get()
	developerAccount := conf.Settings.IOSDeveloperAccount

	if developerAccount == "" {
		// Second priority to machine config
		developerAccounts, err := getDeveloperAccounts()
		if err != nil {
			return errors.New(err.Error())
		}

		// If there is only one in machine config, then we can use it, if there are more you need to add it to local config
		if len(developerAccounts) > 1 {
			return errors.New("There are too many developer accounts in the machine config, we either need one in machine config or one in local config")
		}

		if len(developerAccounts) == 1 {
			developerAccount = developerAccounts[0]
		}
	}

	// Get the bundle identifier for this project
	projectBundleIdentifier := projectSettings["PRODUCT_BUNDLE_IDENTIFIER"]

	_, err := b.Run("FindProvisioningProfile", []string{developerAccount, projectBundleIdentifier})
	if err != nil {
		return errors.New("We could not find your provisioning profile")
	}

	fmt.Println("We have successfully added your profiles to this machine, woohoo")
	return nil
}

func exportIPAForIOS(projectName string) error {
	var b basher.Basher
	// Check if .private/export.plist exists, we can't do shit without it
	exportPlistPath, _ := filepath.Abs(".private/export.plist")
	exportPlistExists := utils.FileExists(exportPlistPath)

	if exportPlistExists == false {
		// If export.plist doesn't exist, create it
		status, next := SetupExportPlist()
		if status != 0 || next == true {
			return errors.New("We could not add an export.plist to your .private folder")
		}
	}

	// Fire the export IPA bash script
	utils.LogMessage("$ xcodebuild -exportArchive -exportOptionsPlist .private/export.plist -archivePath .upshift/" + projectName + ".xcarchive -exportPath .upshift")
	logPath, _ := filepath.Abs(".upshift/logs/xcode-export.log")
	_, err := b.Run("ExportIOS", []string{projectName, logPath})
	if err != nil {
		return errors.New("We could not export IPA\n" + err.Error())
	}

	// Read the last 500 bytes from the whole message, we just want to see what happened at the end
	tailData, err := utils.FileTail(logPath, 500)
	if err != nil {
		return errors.New("It seems we couldn't read the output. Here's what happened\n" + err.Error())
	}

	if strings.Contains(tailData, "EXPORT SUCCEEDED") == false {
		return errors.New("Something went wrong while exporting the IPA, you need to look at this.")
	}

	fmt.Println("We were able to export an IPA successfully, awesome")
	return nil
}

func archiveForIOS(projectType string, projectPath string, scheme string, projectName string) error {
	var b basher.Basher
	utils.LogMessage("$ xcodebuild -" + projectType + " " + projectPath + " -scheme " + scheme + " -derivedDataPath .upshift/build -archivePath .upshift/" + projectName + ".xcarchive archive")
	logPath, _ := filepath.Abs(".upshift/logs/xcode-archive.log")
	_, err := b.Run("ArchiveIOS", []string{projectType, projectPath, scheme, projectName, logPath})
	if err != nil {
		return errors.New("We could not archive for iOS\n" + err.Error())
	}

	// Read the last 500 bytes from the whole message, we just want to what happened at the end
	tailData, err := utils.FileTail(logPath, 500)
	if err != nil {
		return errors.New("It seems we couldn't read the output. Here's what happened\n" + err.Error())
	}

	if strings.Contains(tailData, "ARCHIVE SUCCEEDED") == false {
		return errors.New("Something went wrong with the archive, you need to look at this.")
	}

	fmt.Println("We were able to archive successfully, awesome")
	return nil
}

func testForIOS(projectType string, projectPath string, scheme string, device string) error {
	var b basher.Basher
	utils.LogMessage("$ xctool -" + projectType + " " + projectPath + " -scheme " + scheme + " -sdk iphonesimulator -destination \"platform=iphonesimulator,name=" + device + "\" test")
	logPath, _ := filepath.Abs(".upshift/logs/xcode-test.log")
	_, err := b.Run("TestIOS", []string{projectType, projectPath, scheme, device, logPath})
	if err != nil {
		return errors.New("We could not run tests for iOS\n" + err.Error())
	}

	// Read the last 500 bytes from the whole message, we just want to what happened at the end
	tailData, err := utils.FileTail(logPath, 500)
	if err != nil {
		return errors.New("It seems we couldn't read the output. Here's what happened\n" + err.Error())
	}

	if strings.Contains(tailData, "BUILD SUCCEEDED") == false {
		return errors.New("Something went wrong with the build, you need to look at this.")
	}

	fmt.Println("We were able to run all tests successfully, awesome")
	return nil
}

func compileForIOS(projectType string, projectPath string, scheme string, device string) error {
	var b basher.Basher
	utils.LogMessage("$ xcodebuild -" + projectType + " " + projectPath + " -scheme " + scheme + " -sdk iphonesimulator -destination \"platform=iphonesimulator,name=" + device + "\" -derivedDataPath .upshift/build")
	logPath, _ := filepath.Abs(".upshift/logs/xcode-build.log")
	_, err := b.Run("CompileIOS", []string{projectType, projectPath, scheme, device, logPath})
	if err != nil {
		return errors.New("We could not compile for iOS\n" + err.Error())
	}

	// Read the last 500 bytes from the whole message, we just want to what happened at the end
	tailData, err := utils.FileTail(logPath, 500)
	if err != nil {
		return errors.New("It seems we couldn't read the output. Here's what happened\n" + err.Error())
	}

	if strings.Contains(tailData, "BUILD SUCCEEDED") == false {
		return errors.New("Something went wrong with the build, you need to look at this.")
	}

	fmt.Println("We were able to build successfully, awesome")
	return nil
}

func findAvailableSchemes() error {
	// 1. If scheme defined in config exists, use that
	// 2. If there is no scheme in config, but there is only one scheme, use that
	// 3. If there is no scheme in config and multiple in xcode, show an error

	// If a scheme is available in config, get that
	conf := config.Get()
	confScheme := conf.Settings.IOSScheme

	// Get the list of schemes from xcode
	listDump, err := command.Run([]string{"xcodebuild", "-list"}, "")
	if err != nil {
		return errors.New("We could not get a list from xcode\n" + err.Error())
	}

	listRows := strings.Split(listDump, "\n")
	schemeRows := []string{}
	readingSchemesFlag := false
	for _, row := range listRows {
		// If we have already found "    Schemes:", lets add non null values
		if readingSchemesFlag == true {
			// Remove extra spaces
			row = strings.TrimSpace(row)
			if row != "" {
				// Add if non-null
				schemeRows = append(schemeRows, row)
				// Since we are looping, check if the scheme available from conf is available in the sytem
				// if it is, just return and say it is
				if row == confScheme {
					projectSettings["UP_PROJECT_SCHEME"] = row
					return nil
				}
			}
		}
		// After you see the row which says "    Schemes:", start adding items to list of schemes available
		if strings.Contains(row, "Schemes:") {
			readingSchemesFlag = true
		}
	}

	// Throw an error if there are no schemes available
	if len(schemeRows) == 0 {
		return errors.New("You have no " + c.Red + "schemes" + c.Default + " defined in your project. I'm going home.")
	}

	// Condition 1 : has already been checked, check the error condition
	if confScheme != "" {
		return errors.New("Your config says we should use " + c.Red + confScheme + c.Default + " to build the project. But that scheme is missing!")
	}

	// Condition 2 : Checking
	if len(schemeRows) == 1 {
		// There is only one scheme, you can set this in projectSettings
		fmt.Println("It seems you didn't define a scheme, but Xcode told us about " + c.Blue + schemeRows[0] + c.Default + " so we are using that")
		projectSettings["UP_PROJECT_SCHEME"] = schemeRows[0]
		return nil
	}

	// Condition 3 : Just throw an error
	return errors.New("You have multiple schemes in your project, hence we can't pick one automatically. Please choose one in your config")
}

func startSimulator(device string) {
	// basher returns an error if status > 0 or if there is an error
	// Whenever we start the simulator, for some reason, the exit code is always 255, though there is no error
	// Hence skipping the error check here
	var b basher.Basher
	b.Run("StartSimulator", []string{device})
}

func findIfDeviceIsAvailable(device string) bool {
	instrumentsDump, err := command.Run([]string{"instruments", "-s", "devices"}, "")
	if err != nil {
		return false
	}

	instrumentRows := strings.Split(instrumentsDump, "\n")
	uuidRegexp, _ := regexp.Compile("\\[(.*?)\\]")
	var instruments []string

	for _, instrument := range instrumentRows {
		uuid := uuidRegexp.FindString(instrument)

		isSimulator := strings.Contains(instrument, "(Simulator)")
		instrument = strings.Replace(instrument, "(Simulator)", "", 1)

		simulatorString := "DEVICE   "
		if isSimulator == true {
			simulatorString = "SIMULATOR"
		}

		if uuid != "" {
			instrument = strings.TrimSpace(strings.Replace(instrument, uuid, "", 1))
			instruments = append(instruments, c.Gray+simulatorString+c.Default+" "+c.Bold+c.Green+instrument+c.Default+" "+uuid)
		}

		if instrument == device {
			return true
		}
	}

	utils.LogError("Your device " + c.Red + device + c.Default + " was not found\nThe following devices are available")
	for _, item := range instruments {
		fmt.Println(item)
	}
	return false
}

// Start the simulator
func findXcodeAndOSForSimulator() {
	conf := config.Get()
	var iPhoneOS string

	// Find the correct simulator
	// From here - https://en.wikipedia.org/wiki/Xcode - Xcode 7.0 - 7.x (since Swift 2.0 support)
	switch conf.Settings.IOSXcodeVersion {
	case "7.3.1":
		iPhoneOS = "9.3"
	case "7.3":
		iPhoneOS = "9.3"
	case "7.2.1":
		iPhoneOS = "9.2"
	case "7.2":
		iPhoneOS = "9.2"
	case "7.1.1":
		iPhoneOS = "9.1"
	case "7.1":
		iPhoneOS = "9.1"
	case "7.0.1":
		iPhoneOS = "9.0"
	case "7.0":
		iPhoneOS = "9.0"
	default:
		iPhoneOS = "9.3"
	}

	projectSettings["UP_XCODE_VERSION"] = conf.Settings.IOSXcodeVersion
	projectSettings["UP_SIMULATOR_IPHONE"] = conf.Settings.IOSTestDevice
	projectSettings["UP_SIMULATOR_IPHONEOS"] = iPhoneOS

	// This is sort of the wrong place, but we want scheme for later too
	projectSettings["UP_PROJECT_SCHEME"] = conf.Settings.IOSScheme
}

// Is Simulator running
func isSimulatorRunning() bool {
	_, err := command.Run([]string{"pgrep", "-f", "Simulator.app"}, "")
	if err != nil {
		return false
	}
	return true
}

// Get Project Path
func getProjectPath() {
	conf := config.Get()

	useWorkspace := conf.Settings.IOSUseWorkspace
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
	version, err := command.Run([]string{"xcodebuild", "-showBuildSettings"}, "")
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
	version, err := command.Run([]string{"xcodebuild", "-version"}, "")
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

	conf := config.Get()
	if err != nil {
		fmt.Println("We were unable to load the config file\n", err.Error())

		fmt.Println("You are currently on Xcode-" + currentXcodeVersion + " and the latest Xcode version is " + conf.Settings.IOSXcodeVersion + ". For now, we will continue using the version that you have right now")
		return 0, false
	}

	requiredXcodeVersion := conf.Settings.IOSXcodeVersion

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
	if conf.IsCI() == true {
		// We are on CI, we need to enter password programatically
		RootPassword, err = conf.GetRootPassword()
		if err != nil {
			utils.LogError(err.Error())
			return 1, true
		}
	}

	out, err := command.Run([]string{"sudo", "-S", "xcode-select", "-switch", "/Applications/Xcode-" + requiredXcodeVersion + ".app/"}, RootPassword+"\n")
	if err != nil {
		fmt.Println("We couldn't switch Xcodes, you're going to be stuck with this one")
		return 1, true
	}
	fmt.Println(out)
	fmt.Println("We are now on the " + c.Underline + "Xcode-" + requiredXcodeVersion + c.Default)

	return 0, false
}
