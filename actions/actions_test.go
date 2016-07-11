package actions

import (
	"os"
	"testing"

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
	status := PodInstall()
	assert.Equal(t, 0, status)
	os.Chdir(currentWD)
}

func Test_Actions_AndroidBuild(t *testing.T) {
	currentWD, _ := os.Getwd()
	os.Chdir(os.Getenv("HOME") + "android-chat")
	status := AndroidBuild()
	assert.Equal(t, 0, status)
	os.Chdir(currentWD)
}
