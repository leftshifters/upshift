package actions

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Sigh_All(t *testing.T) {
	var sigh Sigh
	os.Chdir(filepath.Join("..", "ios-test-swift"))

	err := sigh.FindProvisioning("ci@leftshift.io", "com.leftshift.deezeno")
	assert.Nil(t, err)
}
