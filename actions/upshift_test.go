package actions

import (
	"testing"
	"upshift/config"

	"github.com/stretchr/testify/assert"
)

func Test_Upshift_LatestVersion(t *testing.T) {
	var upshift Upshift
	version, err := upshift.LatestVersion()
	assert.Equal(t, "0.8.4", version)
	assert.Nil(t, err)
}

func Test_Upshift_Upgrade(t *testing.T) {
	var upshift Upshift
	t.Log("Upgrading upshift")
	err := upshift.Upgrade("true")
	assert.Nil(t, err)

	conf := config.Get()
	conf.Settings.AppVersion = "1.0"
	t.Log("Upgrading upshift with fake version")
	err = upshift.Upgrade("true")
	assert.Nil(t, err)
}

func Test_Upshift_All(t *testing.T) {
	var upshift Upshift
	upshift.ShowVersion()
	upshift.ShowHelp()
}
