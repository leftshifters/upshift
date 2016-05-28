package bash

import (
	"testing"
)

func TestBash(test *testing.T) {
	type testCase struct {
		command, output string
	}
	var tests = []testCase{{"pwd", "/Users/sudhanshuraheja/gocode/src/upshift/bash"}}

	for _, t := range tests {
		if t.output == Bash(t.command) {
			test.Errorf("The output of command <%s> should be <%s>", t.command, t.output)
		}
	}
}
