package actions

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_GradleWrapper(t *testing.T) {
	os.Chdir("android-chat")
	status := GradleWrapper()
	assert.Equal(t, 0, status)
}
