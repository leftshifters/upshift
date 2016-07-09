package actions

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GradleWrapper(t *testing.T) {
	os.Chdir("android-chat")
	status := GradleWrapper()
	assert.Equal(t, 0, status)
}
