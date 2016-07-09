package actions

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Git_AreSubmodulesUsed(t *testing.T) {
	var g Git
	used := g.AreSubmodulesUsed()
	assert.Equal(t, false, used)
}

func Test_Git_SubmoduleInit(t *testing.T) {
	var g Git
	status, err := g.SubmoduleInit()
	assert.Equal(t, 0, status)
	assert.Nil(t, err)
}

func Test_Git_SubmoduleUpdate(t *testing.T) {
	var g Git
	status, err := g.SubmoduleUpdate()
	assert.Equal(t, 0, status)
	assert.Nil(t, err)
}

func Test_Git_IsRepo(t *testing.T) {
	var g Git
	isRepo := g.IsRepo()
	assert.Equal(t, true, isRepo)

	currentWD, _ := os.Getwd()

	os.Chdir(os.Getenv("HOME"))
	isRepo = g.IsRepo()
	assert.Equal(t, false, isRepo)

	os.Chdir(currentWD)
}

func Test_Git_Branch(t *testing.T) {
	var g Git
	branch, err := g.Branch()
	assert.Equal(t, "go", branch)
	assert.Nil(t, err)
}

func Test_Git_Remote(t *testing.T) {
	var g Git
	remote, err := g.Remote()
	assert.Equal(t, "origin", remote)
	assert.Nil(t, err)
}

func Test_Git_Pull(t *testing.T) {
	var g Git
	remote, _ := g.Remote()
	branch, _ := g.Branch()
	status, err := g.Pull(remote, branch)
	assert.Equal(t, 0, status)
	assert.Nil(t, err)
}
