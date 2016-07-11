package actions

import (
	"errors"
	"fmt"
	"upshift/basher"
	"upshift/config"
	"upshift/utils"
)

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
