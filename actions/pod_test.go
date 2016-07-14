package actions

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Pod_All(t *testing.T) {
	var p Pod
	var c Cocoapods

	// Check if it is installed or not
	cocoapodsInstalled := c.IsInstalled()

	// Change env setting, simulate CI
	os.Setenv("GITLAB_CI", "true")

	installed := p.IsInstalled()
	assert.Equal(t, true, installed)

	used := p.AreUsed()
	assert.Equal(t, false, used)

	status, err := p.Install()
	assert.Equal(t, 0, status)
	assert.Nil(t, err)

	status, err = p.RepoUpdate()
	assert.Equal(t, 0, status)
	assert.Nil(t, err)

	// Bring gem back to original state
	if cocoapodsInstalled == true {
		err := c.Install()
		assert.Nil(t, err)
	}

}
