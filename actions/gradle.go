package actions

import (
	"errors"
	"strings"
	"upshift/basher"
	"upshift/command"
	"upshift/utils"
)

type Gradle struct {
	version          string
	wrapperInstalled bool
	basher           basher.Basher
}

func (g *Gradle) Version() error {
	// Run gradle -v to figure out if it is install
	out, err := command.Run([]string{"gradle", "-v"}, "")
	if err != nil {
		return errors.New("Gradle is not installed, you can download it from http://gradle.com.")
	}

	list := utils.CreateList(out, []string{"Build time", "Build number:", "Revision:", "Groovy:", "Ant:", "JVM:", "OS:", "--------------"})
	g.version = strings.TrimSpace(strings.Replace(strings.Join(list[:], ""), "Gradle", "", 1))
	return nil
}

func (g *Gradle) AddWrapper() (int, error) {
	// Check if gradle is installed
	if g.version == "" {
		g.wrapperInstalled = false
		return 1, errors.New("Gradle version not known. Please check version first")
	}

	// Check if gradlew file exists
	if utils.FileExists("./gradlew") == true {
		g.wrapperInstalled = true
		return 0, nil
	}

	// So, gradle is installed, just need to install wrapper [SetupGradleW]
	// I won't touch anything to do with gradle and pipes with a ten foot pole, so this goes to basher
	utils.LogMessage("$ gradle wrapper")
	status, err := g.basher.Run("GradleWrapper", []string{})
	if err != nil {
		g.wrapperInstalled = false
		return status, errors.New("We couldn't initialise gradle wrapper\n" + err.Error())
	}

	g.wrapperInstalled = true
	return status, nil
}

func (g *Gradle) Clean(logPath string) (int, error) {
	// Check if gradle is installed
	if g.version == "" {
		return 1, errors.New("Gradle version not known. Please check version first")
	}

	// Check if gradlew file exists
	if utils.FileExists("./gradlew") == false {
		return 1, errors.New("Gradle wrapper is not installed")
	}

	status, err := g.basher.RunAndTail("GradlewClean", []string{logPath}, logPath, "BUILD SUCCESSFUL")
	return status, err
}
