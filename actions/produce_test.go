package actions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Produce(t *testing.T) {
	var produce Produce
	err := produce.CreateAppOnITunes("ci@leftshift.io", "com.leftshift.SwiftDemo", "SwiftDemo")
	assert.Nil(t, err)
}
