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
