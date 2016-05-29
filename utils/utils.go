package utils

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

func init() {

}

func IsDocker() (bool, error) {
	// To check if it's docker or not, find out if /proc/1/cgroup has Docker written anywhere
	cGroupFile := "/proc/1/cgroup"

	if _, err := os.Stat(cGroupFile); err == nil {

		cGroupBytes, err := ioutil.ReadFile(cGroupFile)
		if err != nil {
			return false, errors.New("Could not read /proc/1/cgroup")
		}

		// Check if docker is written inside the cGroup file
		cGroupString := string(cGroupBytes)
		return strings.Contains(cGroupString, "docker"), nil

	} else {
		// File not found, ceratinly not docker
		return false, nil
	}
}
