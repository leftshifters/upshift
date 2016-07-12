package actions

import (
	"errors"
	"fmt"

	"github.com/leftshifters/upshift/basher"
)

// Pilot : Handle everything related to pilot
type Pilot struct{}

// UploadToITunes : upload the binary to itunes
func (p *Pilot) UploadToITunes(developerAccount string, scheme string, name string) error {
	var b basher.Basher

	// Add SwitSources if required - AddSwiftSources
	_, err := b.Run("AddSwiftSources", []string{name, scheme})
	if err != nil {
		return errors.New("We could not add SwiftSources to the IPA\n" + err.Error())
	}

	_, err = b.Run("UploadIPAoniTunes", []string{developerAccount, ".upshift/" + scheme + ".ipa"})
	if err != nil {
		return errors.New("We could not upload the IPA on iTunes\n" + err.Error())
	}

	fmt.Println("We have successfully uploaded this IPA on iTunes, it's all yours now")
	return nil

}
