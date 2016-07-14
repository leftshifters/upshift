package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Run(t *testing.T) {
	_, err := Run([]string{}, "")
	assert.Contains(t, err.Error(), "send in something")

	out, _ := Run([]string{"ls"}, "")
	assert.Contains(t, out, "command_test.go")

	_, err = Run([]string{"ls1"}, "")
	assert.Contains(t, err.Error(), "unable to run")
}
