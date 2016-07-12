package actions

import (
	"errors"
	"fmt"

	"github.com/leftshifters/upshift/basher"
	"github.com/leftshifters/upshift/utils"
)

// Sigh : Handle everything related to sigh
type Sigh struct{}

// FindProvisioning : find and install a provisioning profile
func (s *Sigh) FindProvisioning(developerAccount string, bundleIdentifier string) error {
	var b basher.Basher
	utils.LogMessage("We will now try to find the provisioning profile")

	_, err := b.Run("FindProvisioningProfile", []string{developerAccount, bundleIdentifier})
	if err != nil {
		return errors.New("We could not find your provisioning profile")
	}

	fmt.Println("We have successfully added your profiles to this machine, woohoo")
	return nil
}
