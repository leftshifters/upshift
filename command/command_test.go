package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Run(t *testing.T) {
	_, err := Run([]string{}, "")
	assert.Contains(t, err.Error(), "send in something")

	out, _ := Run([]string{"ls"}, "")
	assert.Contains(t, out, "command_test.go")

	out, err = Run([]string{"ls1"}, "")
	assert.Contains(t, err.Error(), "unable to run")
}
