package actions

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_Gradle_Version(t *testing.T) {
	var g Gradle
	os.Chdir("android-chat")
	g.Version()
	assert.NotEmpty(t, g.version)
}

func Test_Gradle_AddWrapper(t *testing.T) {
	var g Gradle
	os.Chdir("android-chat")
	g.Version()
	g.AddWrapper()
	assert.Equal(t, true, g.wrapperInstalled)
}

func Test_Gradle_Clean(t *testing.T) {
	var g Gradle
	os.Chdir("android-chat")
	g.Version()
	g.AddWrapper()
	g.Clean("clean.log")
}
