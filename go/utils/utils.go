package utils

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

func GetAppVersion() string {
	return "0.7.3"
}

func IsCI() bool {
	isGitlab := os.Getenv("GITLAB_CI")
	if isGitlab == "true" {
		return true
	} else {
		return false
	}
}

// GITLAB_CI=$(printenv GITLAB_CI)

// # Overall, OR all of them to find out is it is running via CI
// CI=${GITLAB_CI}

func IsDocker() (bool, error) {
	// To check if it's docker or not, find out if /proc/1/cgroup has Docker written anywhere
	cGroupFile := "/proc/1/cgroup"

	if _, err := os.Stat(cGroupFile); err == nil {

		cGroupBytes, err := ioutil.ReadFile(cGroupFile)
		if err != nil {
			return false, errors.New("Could not read /proc/1/cgroup " + err.Error())
		}

		// Check if docker is written inside the cGroup file
		cGroupString := string(cGroupBytes)
		return strings.Contains(cGroupString, "docker"), nil

	} else {
		// File not found, ceratinly not docker
		return false, nil
	}
}

func ReadIfFileExists(filePath string) (string, error) {
	if _, err := os.Stat(filePath); err == nil {
		fileBytes, err := ioutil.ReadFile(filePath)
		if err != nil {
			return "", errors.New("Could not read file" + filePath + err.Error())
		}
		return string(fileBytes), nil
	} else {
		return "", errors.New("File does not exist " + filePath + err.Error())
	}
}

func FileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	if err == nil {
		return true
	} else {
		return false
	}
}
