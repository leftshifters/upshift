package colours

import (
	"fmt"
	"testing"
)

func TestLogColours(t *testing.T) {
	fmt.Println("This", Red.Format, "word", Default.Format, "should be in red")
	fmt.Println("This", Green.Format, "word", Default.Format, "should be in green")
	fmt.Println("This", Yellow.Format, "word", Default.Format, "should be in yellow")
	fmt.Println("This", Blue.Format, "word", Default.Format, "should be in blue")
	fmt.Println("This", Gray.Format, "word", Default.Format, "should be in gray")
	fmt.Println("This", White.Format, "word", Default.Format, "should be in white")
}
