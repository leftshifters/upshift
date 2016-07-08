package basher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Run(t *testing.T) {
	var b Basher
	status, _ := b.Run("TestScript", []string{"Sleeping for 1 second"})
	assert.Equal(t, 0, status)

	status, err := b.Run("TestScriptWhichDoesntExist", []string{"This message will never show up"})
	assert.Contains(t, err.Error(), "problem running")
	assert.Equal(t, 127, status)
}

func Test_RunAndTail(t *testing.T) {
	var b Basher
	status, err := b.RunAndTail("TestScript2", []string{"Sleeping for 1 second", "tests.log"}, "tests.log", "basher_test.go")
	assert.Equal(t, 0, status)
	assert.Nil(t, err)
}
