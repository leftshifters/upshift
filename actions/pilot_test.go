package actions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Pilot(t *testing.T) {
	var pilot Pilot
	err := pilot.UploadToITunes("ci@leftshift.io", "com.leftshift.name", "Name")
	assert.Nil(t, err)
}
