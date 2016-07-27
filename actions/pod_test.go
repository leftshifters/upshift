package actions

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Pod_All(t *testing.T) {
	var p Pod

	os.Chdir(filepath.Join("..", "ios-test-swift"))

	installed := p.IsInstalled()
	assert.Equal(t, true, installed)

	used := p.AreUsed()
	assert.Equal(t, true, used)

	status, err := p.Install()
	assert.Equal(t, 0, status)
	assert.Nil(t, err)

	status, err = p.RepoUpdate()
	assert.Equal(t, 0, status)
	assert.Nil(t, err)

}
