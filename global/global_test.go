package global

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	var c Config
	err := Set(c)
	if err != nil {
		fmt.Println(err.Error())
	}

	c, _ = Get()
	fmt.Println(c)
}
