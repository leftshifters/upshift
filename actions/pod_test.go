package actions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Pod_IsInstalled(t *testing.T) {
	var p Pod
	installed := p.IsInstalled()
	assert.Equal(t, true, installed)
}

func Test_Pod_AreUsed(t *testing.T) {
	var p Pod
	used := p.AreUsed()
	assert.Equal(t, false, used)
}

func Test_Pod_Install(t *testing.T) {
	var p Pod
	status, err := p.Install()
	assert.Equal(t, 0, status)
	assert.Nil(t, err)
}

func Test_Pod_RepoUpdate(t *testing.T) {
	var p Pod
	status, err := p.RepoUpdate()
	assert.Equal(t, 0, status)
	assert.Nil(t, err)
}
