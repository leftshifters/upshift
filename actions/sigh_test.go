package actions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Sigh_All(t *testing.T) {
	var sigh Sigh

	err := sigh.FindProvisioning("ci@leftshift.io", "com.leftshift.deezeno")
	assert.Nil(t, err)
}
