package logColours

import (
	"log"
	"testing"
)

func TestLogColours(t *testing.T) {
	log.Println("This", Red.Format, "word", Default.Format, "should be in red")
	log.Println("This", Green.Format, "word", Default.Format, "should be in green")
	log.Println("This", Yellow.Format, "word", Default.Format, "should be in yellow")
	log.Println("This", Blue.Format, "word", Default.Format, "should be in blue")
	log.Println("This", Gray.Format, "word", Default.Format, "should be in gray")
	log.Println("This", White.Format, "word", Default.Format, "should be in white")
}
