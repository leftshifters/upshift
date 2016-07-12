package actions

import (
	"errors"
	"fmt"

	"github.com/leftshifters/upshift/basher"
	"github.com/leftshifters/upshift/utils"
)

// Produce : Construct to handle all things related to produce
type Produce struct {
}

// CreateAppOnITunes : create an app on iTunes if it does not exist
func (p *Produce) CreateAppOnITunes(developerAccount string, bundleIdentifier string, name string) error {
	var b basher.Basher
	utils.LogMessage("Create an app on iTunesConnect if it doesn't exist")

	name = name + " Beta by Upshift"
	_, err := b.Run("CreateAppOnItunes", []string{developerAccount, bundleIdentifier, name})
	if err != nil {
		return errors.New("We could not create the app on iTunes\n" + err.Error())
	}

	fmt.Println("We have successfully added this app on iTunes, woohoo")
	return nil

}
