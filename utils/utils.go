package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	colours "upshift/colours"
)

// Get the App Version
func GetAppVersion() string {
	return "0.7.3"
}

// Log Information, this shows up in blue
func LogInfo(message string) {
	fmt.Println("🔰  " + colours.Green.Format + colours.Bold.Format + "Maybe should " + colours.Underline.Format + "know this" + colours.Default.Format)
	fmt.Println(message + "\n")
}

// Log an error, show them this shit in color, red most probably
func LogError(message string) {
	fmt.Println("☎️  " + colours.Red.Format + colours.Bold.Format + "Shit! Something broke" + colours.Default.Format)
	fmt.Println(message + "\n")
}

// Check if the currect script is running in a CI
func IsCI() bool {
	// Inspiration
	// GITLAB_CI=$(printenv GITLAB_CI)
	// # Overall, OR all of them to find out is it is running via CI
	// CI=${GITLAB_CI}

	// Get GITLAB_CI from the environment
	isGitlab := os.Getenv("GITLAB_CI")
	if isGitlab == "true" {
		return true
	} else {
		return false
	}
}

// Check if the current script is running in a docker container
func IsDocker() bool {
	// To check if it's docker or not, find out if /proc/1/cgroup has Docker written anywhere
	// We don't need to return an error on this, just a true of false
	cGroupFile := "/proc/1/cgroup"

	if FileExists(cGroupFile) == false {
		// File not found, ceratinly not docker or a linux machine
		return false
	}

	// Read the file, and then check if the work docker is written inside it
	// We can read it directly because we know the file exits
	cGroupBytes, _ := ReadIfFileExists(cGroupFile)
	cGroupString := string(cGroupBytes)
	return strings.Contains(cGroupString, "docker")

}

// Read a file if it exists
func ReadIfFileExists(filePath string) (string, error) {
	// Check if file exits, if it doesn't just return an error
	if FileExists(filePath) == false {
		return "", errors.New("File does not exist " + filePath)
	}

	// If file exists, go ahead and read the shit out of it
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", errors.New("Could not read file" + filePath + err.Error())
	}

	return string(fileBytes), nil
}

// Simply to check if a file exists or not
func FileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	if err == nil {
		return true
	} else {
		return false
	}
}
