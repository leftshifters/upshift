package actions

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/leftshifters/upshift/config"
	"github.com/stretchr/testify/assert"
)

func Test_Actions_GradleWrapper(t *testing.T) {
	currentWD, _ := os.Getwd()
	os.Chdir(filepath.Join("..", "android-test"))

	status := SetupGradleWrapper()
	assert.Equal(t, 0, status)

	os.Chdir(currentWD)
}

func Test_Actions_PodInstall(t *testing.T) {
	currentWD, _ := os.Getwd()
	os.Chdir(filepath.Join("..", "ios-test-swift"))

	status := SetupPods()
	assert.Equal(t, 0, status)

	os.Chdir(currentWD)
}

func Test_Actions_AndroidBuild(t *testing.T) {
	conf := config.Get()
	currentWD, _ := os.Getwd()
	os.Chdir(filepath.Join("..", "android-test"))
	conf.Settings.CleanBeforeBuild = true

	status := AndroidBuild()
	assert.Equal(t, 0, status)

	os.Chdir(currentWD)
}
