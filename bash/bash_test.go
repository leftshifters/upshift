package bash

import (
	"fmt"
	"testing"
)

func TestBash(test *testing.T) {
	returnValue := Bash("ls -l")
	fmt.Println("Tested ", returnValue)
}
