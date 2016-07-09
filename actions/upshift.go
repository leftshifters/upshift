package actions

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"upshift/basher"
	"upshift/config"
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
