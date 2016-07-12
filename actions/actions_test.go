package actions

import (
	"os"
	"testing"

	"github.com/leftshifters/upshift/config"
	"github.com/stretchr/testify/assert"
)

func Test_Actions_GradleWrapper(t *testing.T) {
	currentWD, _ := os.Getwd()
	os.Chdir("android-chat")

	status := GradleWrapper()
	assert.Equal(t, 0, status)

	os.Chdir(currentWD)
}

func Test_Actions_PodInstall(t *testing.T) {
	currentWD, _ := os.Getwd()
	os.Chdir(os.Getenv("HOME") + "/code/deezeno-ios")
	os.Create(os.Getenv("HOME") + "/code/deezeno-ios/Podfile")

	status := PodInstall()
	assert.Equal(t, 0, status)

	os.Remove(os.Getenv("HOME") + "/code/deezeno-ios/Podfile")
	os.Chdir(currentWD)
}

func Test_Actions_AndroidBuild(t *testing.T) {
	conf := config.Get()
	currentWD, _ := os.Getwd()
	os.Chdir("android-chat")
	conf.Settings.CleanBeforeBuild = true

	status := AndroidBuild()
	assert.Equal(t, 0, status)

	os.Chdir(currentWD)
}
