package actions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Pod_All(t *testing.T) {
	var p Pod

	installed := p.IsInstalled()
	assert.Equal(t, true, installed)

	used := p.AreUsed()
	assert.Equal(t, false, used)

	status, err := p.Install()
	assert.Equal(t, 0, status)
	assert.Nil(t, err)

	status, err = p.RepoUpdate()
	assert.Equal(t, 0, status)
	assert.Nil(t, err)

}
