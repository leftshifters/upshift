package actions

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/leftshifters/upshift/basher"
	"github.com/leftshifters/upshift/colours"
	"github.com/leftshifters/upshift/command"
	"github.com/leftshifters/upshift/config"
	"github.com/leftshifters/upshift/utils"
)

// Xcodebuild : construct to handle all xcode functions
type Xcodebuild struct {
	Name             string
	BundleIdentifier string
	Type             string
	Extension        string
	XcodeVersion     string
	iPhoneOS         string
	TestDevice       string
	Scheme           string
	DeviceName       string
	Path             string
	UseWorkspace     bool
	UsePods          bool
}

// LoadSettings : load settings from xcodebuild
func (x *Xcodebuild) LoadSettings() error {
	settings, err := command.Run([]string{"xcodebuild", "-showBuildSettings"}, "")
	if err != nil {
		return errors.New("We are unable to read settings for this iOS Project\n" + err.Error())
	}

	// Load settings from xcodebuild -showBuildSettings
	settingsRows := strings.Split(settings, "\n")
	for _, row := range settingsRows {
		settingsKeys := strings.Split(row, "=")
		if len(settingsKeys) == 2 {
			key := strings.TrimSpace(settingsKeys[0])
			value := strings.TrimSpace(settingsKeys[1])
			// fmt.Println("key", key, "value", value)

			switch key {
			case "PROJECT_NAME":
				x.Name = value
			case "PRODUCT_BUNDLE_IDENTIFIER":
				x.BundleIdentifier = value
			default:
			}
		}
	}

	// Load config
	var pod Pod
	conf := config.Get()

	x.UseWorkspace = conf.Settings.IOSUseWorkspace
	x.UsePods = pod.AreUsed()

	x.Type = "project"
	x.Extension = ".xcodeproj"
	if x.UseWorkspace || x.UsePods {
		x.Type = "workspace"
		x.Extension = ".xcworkspace"
	}
	x.Path = x.Name + x.Extension

	// Find the correct simulator
	// From here - https://en.wikipedia.org/wiki/Xcode - Xcode 7.0 - 7.x (since Swift 2.0 support)
	x.XcodeVersion = conf.Settings.IOSXcodeVersion
	x.TestDevice = conf.Settings.IOSTestDevice
	x.Scheme = conf.Settings.IOSScheme

	os93 := "9.3"
	os92 := "9.2"
	os91 := "9.1"
	os90 := "9.0"

	switch conf.Settings.IOSXcodeVersion {
	case "7.3.1":
		x.iPhoneOS = os93
	case "7.3":
		x.iPhoneOS = os93
	case "7.2.1":
		x.iPhoneOS = os92
	case "7.2":
		x.iPhoneOS = os92
	case "7.1.1":
		x.iPhoneOS = os91
	case "7.1":
		x.iPhoneOS = os91
	case "7.0.1":
		x.iPhoneOS = os90
	case "7.0":
		x.iPhoneOS = os90
	default:
		x.iPhoneOS = os93
	}

	x.DeviceName = x.TestDevice + " (" + x.iPhoneOS + ")"

	return nil
}

// FindSchemes : find a scheme which we can use
func (x *Xcodebuild) FindSchemes() error {
	conf := config.Get()

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
				if row == conf.Settings.IOSScheme {
					x.Scheme = row
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
		return errors.New("You have no " + colours.Red + "schemes" + colours.Default + " defined in your project. You need to share them.")
	}

	// 1. If scheme defined in config exists, use that
	// 2. If there is no scheme in config, but there is only one scheme, use that
	// 3. If there is no scheme in config and multiple in xcode, show an error

	// Condition 1 : has already been checked, check the error condition
	// If we didn't find a matching scheme, but your config has marked one, it means we didn't find it
	if conf.Settings.IOSScheme != "" && x.Scheme == "" {
		return errors.New("Your config says we should use " + colours.Red + conf.Settings.IOSScheme + colours.Default + " to build the project. But that scheme is missing!")
	}

	// Condition 2 : Checking
	if len(schemeRows) == 1 {
		// There is only one scheme, you can set this in projectSettings
		fmt.Println("It seems you didn't define a scheme in your config, but Xcode told us about " + colours.Blue + schemeRows[0] + colours.Default + " so we are using that")
		x.Scheme = schemeRows[0]
		return nil
	}

	// Condition 3 : Just throw an error
	return errors.New("You have multiple schemes in your project, hence we can't pick one automatically. Please choose one in your config")
}

// Build : compile an iOS project
func (x *Xcodebuild) Build() error {
	if x.TestDevice == "" || x.Scheme == "" || x.Type == "" || x.Path == "" {
		return errors.New("We need the TestDevice, Scheme, Type, Path to proceed")
	}

	utils.LogMessage("$ xcodebuild -" + x.Type + " " + x.Path + " -scheme " + x.Scheme + " -sdk iphonesimulator -destination \"platform=iphonesimulator,name=" + x.TestDevice + "\" -derivedDataPath .upshift/build")

	var b basher.Basher
	_, err := b.RunAndTail("IOSBuild", []string{x.Type, x.Path, x.Scheme, x.TestDevice, ".upshift/logs/xcode-build.log"}, ".upshift/logs/xcode-build.log", []string{"BUILD SUCCEEDED"}, []string{})
	if err != nil {
		return err
	}

	return nil
}

// Archive : archive a project
func (x *Xcodebuild) Archive() error {
	if x.Name == "" || x.Scheme == "" || x.Type == "" || x.Path == "" {
		return errors.New("We need the Name, Scheme, Type, Path to proceed")
	}

	utils.LogMessage("$ xcodebuild -" + x.Type + " " + x.Path + " -scheme " + x.Scheme + " -derivedDataPath .upshift/build -archivePath .upshift/" + x.Name + ".xcarchive archive")

	var b basher.Basher
	_, err := b.RunAndTail("IOSArchive", []string{x.Type, x.Path, x.Scheme, x.Name, ".upshift/logs/xcode-archive.log"}, ".upshift/logs/xcode-archive.log", []string{"ARCHIVE SUCCEEDED"}, []string{})
	if err != nil {
		return err
	}

	return nil
}

// ExportIPA : export an IPA
func (x *Xcodebuild) ExportIPA() error {
	if x.Name == "" {
		return errors.New("We need a project name to proceed")
	}

	utils.LogMessage("$ xcodebuild -exportArchive -exportOptionsPlist .private/export.plist -archivePath .upshift/" + x.Name + ".xcarchive -exportPath .upshift")

	var b basher.Basher
	_, err := b.RunAndTail("IOSExport", []string{x.Name, ".upshift/logs/xcode-export.log"}, ".upshift/logs/xcode-export.log", []string{"EXPORT SUCCEEDED"}, []string{})
	if err != nil {
		return err
	}

	return nil
}

// SetupExportPlist : create an export.plist file if it doesn't exist
func (x *Xcodebuild) SetupExportPlist() error {
	// Check if the export.plist exists
	if utils.FileExists(".private/export.plist") {
		return nil
	}

	sampleExportPlist := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
        <key>method</key>
        <string>app-store</string>
        <key>uploadSymbols</key>
        <true/>
	<key>uploadBitcode</key>
	<false/>
</dict>
</plist>"`

	exportPlistBytes := []byte(sampleExportPlist)

	err := ioutil.WriteFile(".private/export.plist", exportPlistBytes, 0644)
	if err != nil {
		return errors.New("We could not write the .private/export.plist file\n" + err.Error())
	}

	fmt.Println("We just added a sample file to .private/export.plist!")
	return nil
}

// IncrementBuildNumber : increment the build number.
func (x *Xcodebuild) IncrementBuildNumber() error {
	if x.Name == "" {
		return errors.New("Please find the project name first")
	}

	var b basher.Basher
	_, err := b.Run("IOSIncrementBuildNumber", []string{x.Name})
	if err != nil {
		return errors.New("We couldn't incremenet the build number")
	}

	return nil
}

// SwitchXcode : change the xcode version
func (x *Xcodebuild) SwitchXcode() error {
	currentVersion, err := x.Version()
	if err != nil {
		return err
	}

	if currentVersion == x.XcodeVersion {
		fmt.Println("You are already on Xcode v" + currentVersion)
		return nil
	}

	if utils.FileExists("/Applications/Xcode-"+x.XcodeVersion+".app/") == false {
		return errors.New("It seems you don't have /Applications/Xcode-" + x.XcodeVersion + ".app/\nWe expect XCode versions to be placed like this\n/Applications/Xcode-7.2.app\n/Applications/Xcode-7.3.app")
	}

	conf := config.Get()
	RootPassword, err := conf.GetRootPassword()
	if conf.IsCI() == true && err != nil {
		return err
	}

	inputVia := "-S"
	if conf.IsCI() == false && RootPassword == "" {
		inputVia = "-k"
	}

	out, err := command.Run([]string{"sudo", inputVia, "xcode-select", "--switch", "/Applications/Xcode-" + x.XcodeVersion + ".app/"}, RootPassword+"\n")
	if err != nil {
		return errors.New("We couldn't switch Xcodes, you're going to be stuck with this one")
	}
	fmt.Println(out)
	fmt.Println("We are now on the " + colours.Underline + "Xcode-" + x.XcodeVersion + colours.Default)
	return nil
}

// Version : current xcodebuild version
func (x *Xcodebuild) Version() (string, error) {
	versionData, err := command.Run([]string{"xcodebuild", "-version"}, "")
	if err != nil {
		return "", errors.New("We were unable to get the Xcode version\n" + err.Error())
	}

	var currentXcodeVersion string
	versionRows := strings.Split(versionData, "\n")
	for _, row := range versionRows {
		if strings.Contains(row, "Xcode") == true {
			currentXcodeVersion = strings.TrimSpace(strings.Trim(row, "Xcode"))
		}
	}

	return currentXcodeVersion, nil
}

// Run : run a build in the simulator
func (x *Xcodebuild) Run() error {
	if x.BundleIdentifier == "" {
		return errors.New("We need the bundle identifier to proceed")
	}

	var build string
	debug := ".upshift/build/Build/Products/Debug-iphonesimulator/" + x.Name + ".app"
	release := ".upshift/build/Build/Products/Release-iphonesimulator/" + x.Name + ".app"

	if utils.FileExists(debug) {
		build, _ = filepath.Abs(debug)
	} else if utils.FileExists(release) {
		build, _ = filepath.Abs(release)
	} else {
		return errors.New("You have not build for the simulator yet")
	}

	// If file exists, push it to the simulator
	fmt.Println("Installing the app in the simulator")
	_, err := command.Run([]string{"xcrun", "simctl", "install", "booted", build}, "")
	if err != nil {
		return errors.New("We could not deploy the app to the simulator\n" + err.Error())
	}

	// Start the app
	fmt.Println("Starting the app in the simulator")
	_, err = command.Run([]string{"xcrun", "simctl", "launch", "booted", x.BundleIdentifier}, "")
	if err != nil {
		return errors.New("We could not deploy the app to the simulator\n" + err.Error())
	}

	return nil
}

// InstallCertificates : Find where are the certificates and install them
func (x *Xcodebuild) InstallCertificates() error {
	utils.LogMessage("Checking for apple.cer, distribution.p12 and distribution.cer in .private")

	var appleCert, distributionCert, distributionP12Cert string

	if utils.FileExists(".private/apple.cer") && utils.FileExists(".private/distribution.cer") && utils.FileExists(".private/distribution.p12") {
		// Install certs from .private
		appleCert, _ = filepath.Abs(".private/apple.cer")
		distributionCert, _ = filepath.Abs(".private/distribution.cer")
		distributionP12Cert, _ = filepath.Abs(".private/distribution.p12")
	} else {
		// Since they don't exist in .private, look in machine config
		conf := config.Get()
		base := conf.Settings.IOSCertificatePath

		if base == "" {
			return errors.New("The certificates don't exist in both .private and global conf")
		}

		if utils.FileExists(base+"/apple.cer") && utils.FileExists(base+"/distribution.cer") && utils.FileExists(base+"/distribution.p12") {
			appleCert, _ = filepath.Abs(base + "/apple.cer")
			distributionCert, _ = filepath.Abs(base + "/distribution.cer")
			distributionP12Cert, _ = filepath.Abs(base + "/distribution.p12")
		}
	}

	if appleCert == "" {
		return errors.New("All the certificates don't exist in both .private and global conf")
	}

	outApple, err := command.Run([]string{"security", "import", appleCert, "-k", os.Getenv("HOME") + "/Library/Keychains/login.keychain", "-T", "/usr/bin/codesign", "-T", "/usr/bin/security"}, "")
	fmt.Println(outApple)
	if err != nil && strings.Contains(outApple, "already exists in the keychain") == false {
		return err
	}

	outCert, err := command.Run([]string{"security", "import", distributionCert, "-k", os.Getenv("HOME") + "/Library/Keychains/login.keychain", "-T", "/usr/bin/codesign", "-T", "/usr/bin/security"}, "")
	fmt.Println(outCert)
	if err != nil && strings.Contains(outCert, "already exists in the keychain") == false {
		return err
	}

	outP12, err := command.Run([]string{"security", "import", distributionP12Cert, "-k", os.Getenv("HOME") + "/Library/Keychains/login.keychain", "-T", "/usr/bin/codesign", "-T", "/usr/bin/security", "-P", ""}, "")
	fmt.Println(outP12)
	if err != nil && strings.Contains(outP12, "already exists in the keychain") == false {
		return err
	}

	return nil
}
