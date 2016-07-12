package actions

import (
	"errors"
	"strings"

	"github.com/leftshifters/upshift/basher"
	"github.com/leftshifters/upshift/command"
	"github.com/leftshifters/upshift/utils"
)

// Gradle : Keep track of everything related to gradle in upshift
type Gradle struct {
	version          string
	wrapperInstalled bool
	basher           basher.Basher
}

// Install : install gradle on a machine
func (g *Gradle) Install() error {
	var b basher.Basher
	_, err := b.Run("GradleInstall", []string{})
	if err != nil {
		return err
	}
	return nil
}

// Version : Find the version number of gradle installed
func (g *Gradle) Version() error {
	// Run gradle -v to figure out if it is install
	utils.LogMessage("$ gradle -v")
	out, err := command.Run([]string{"gradle", "-v"}, "")
	if err != nil {
		err = g.Install()
		if err != nil {
			return err
		}
	}

	list := utils.CreateList(out, []string{"Build time", "Build number:", "Revision:", "Groovy:", "Ant:", "JVM:", "OS:", "--------------"})
	g.version = strings.TrimSpace(strings.Replace(strings.Join(list[:], ""), "Gradle", "", 1))
	return nil
}

// AddWrapper : Check and confirm if gradlew exists in the project
func (g *Gradle) AddWrapper() (int, error) {
	// Check if gradle is installed
	if g.version == "" {
		err := g.Version()
		if err != nil {
			return 1, err
		}
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

// Task : Execute a gradle task
func (g *Gradle) Task(task string, params []string, logPath string, success string) (int, error) {
	// Check if task exists
	if task == "" {
		return 1, errors.New("Gradle needs a task to run")
	}

	// Check if gradle is installed
	if g.version == "" {
		err := g.Version()
		if err != nil {
			return 1, err
		}
	}

	// Check if gradlew file exists
	if utils.FileExists("./gradlew") == false {
		status, err := g.AddWrapper()
		if err != nil {
			return status, err
		}
	}

	utils.LogMessage("$ ./gradlew " + task + " " + strings.Join(params[:], " "))
	status, err := g.basher.RunAndTail("GradlewTask", []string{task, logPath}, logPath, []string{success}, []string{})
	return status, err
}

// Clean : Execute gradle clean on the Android project
func (g *Gradle) Clean(logPath string) (int, error) {
	return g.Task("clean", []string{}, logPath, "BUILD SUCCESSFUL")
}

// Lint : Execute gradle lint on the Android project
func (g *Gradle) Lint(logPath string) (int, error) {
	return g.Task("lint", []string{}, logPath, "BUILD SUCCESSFUL")
}

// Uninstall : Remove installed versions of the app from connected devices
func (g *Gradle) Uninstall(logPath string) (int, error) {
	return g.Task("uninstallAll", []string{}, logPath, "BUILD SUCCESSFUL")
}

// InstallDebug : Install the debug app on to connected devices
func (g *Gradle) InstallDebug(logPath string) (int, error) {
	return g.Task("installDebug", []string{"--stacktrace"}, logPath, "BUILD SUCCESSFUL")
}

// Assemble : Build an android project
func (g *Gradle) Assemble(logPath string) (int, error) {
	return g.Task("assemble", []string{"--stacktrace"}, logPath, "BUILD SUCCESSFUL")
}
