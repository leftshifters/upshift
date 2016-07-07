package colours

import (
	"testing"
)

func TestLogColours(t *testing.T) {
	t.Log("This", Red, "word", Default, "should be in red")
	t.Log("This", Green, "word", Default, "should be in green")
	t.Log("This", Yellow, "word", Default, "should be in yellow")
	t.Log("This", Blue, "word", Default, "should be in blue")
	t.Log("This", Gray, "word", Default, "should be in gray")
	t.Log("This", White, "word", Default, "should be in white")
}
