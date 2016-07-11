package actions

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"upshift/basher"
	"upshift/config"
	"upshift/utils"
)

// SetupProfiles : Setup provisioning profiles
// This will probably be run once in a while
// * SIGH *
func SetupProfiles() int {
	conf := config.Get()
	var b basher.Basher
	email := strings.TrimSpace(conf.Settings.IOSDeveloperAccount)

	if email != "" {
		fmt.Println("Trying to repair your provisioning profiles and installing new and fixed ones for " + email)
		_, err := b.Run("FetchAndRepairProvisioningProfiles", []string{email})
		if err != nil {
			utils.LogError("We couldn't fix and install your provisioning profiles")
			return 1
		}
	}
	return 0
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

// * Pilot *
func uploadBuildToItunes() error {
	// var b basher.Basher
	utils.LogMessage("Upload the IPA on iTunesConnect")

	// Get the username which will need to login
	// Highest priority to local config
	// conf := config.Get()
	// developerAccount := conf.Settings.IOSDeveloperAccount

	// projectScheme := projectSettings["UP_PROJECT_SCHEME"]
	// projectName := projectSettings["PROJECT_NAME"]

	// Add SwitSources if required - AddSwiftSources
	// status, err := b.Run("AddSwiftSources", []string{projectName, projectScheme})
	// fmt.Println("status", status)
	// if err != nil {
	// 	fmt.Println("err", err.Error())
	// 	return errors.New("We could not add SwiftSources to the IPA")
	// }

	// _, err = b.Run("UploadIPAoniTunes", []string{developerAccount, ".upshift/" + projectScheme + ".ipa"})
	// if err != nil {
	// 	return errors.New("We could not upload the IPA on iTunes")
	// }

	fmt.Println("We have successfully uploaded this IPA on iTunes, it's all yours now")
	return nil
}

// * Product *
func createAppOniTunes() error {
	var b basher.Basher
	utils.LogMessage("Create an app on iTunesConnect if it doesn't exist")

	// Get the username which will need to login
	// Highest priority to local config
	conf := config.Get()
	developerAccount := conf.Settings.IOSDeveloperAccount

	// Get the bundle identifier for this project
	// projectBundleIdentifier := projectSettings["PRODUCT_BUNDLE_IDENTIFIER"]

	// Get the name of the project
	// And add your shit to it
	// projectName := projectSettings["PROJECT_NAME"] + " Beta by Upshift"

	_, err := b.Run("CreateAppOnItunes", []string{developerAccount, "projectBundleIdentifier", "projectName"})
	if err != nil {
		return errors.New("We could not create the app on iTunes\n" + err.Error())
	}

	fmt.Println("We have successfully added this app on iTunes, woohoo")
	return nil
}

// * Sigh *
func addProvisioningProfiles() error {
	// var b basher.Basher
	utils.LogMessage("We will now try to find the provisioning profile")

	// Get the username which will need to login
	// Highest priority to local config
	// conf := config.Get()
	// developerAccount := conf.Settings.IOSDeveloperAccount

	// Get the bundle identifier for this project
	// projectBundleIdentifier := projectSettings["PRODUCT_BUNDLE_IDENTIFIER"]

	// _, err := b.Run("FindProvisioningProfile", []string{developerAccount, projectBundleIdentifier})
	// if err != nil {
	// 	return errors.New("We could not find your provisioning profile")
	// }

	fmt.Println("We have successfully added your profiles to this machine, woohoo")
	return nil
}
