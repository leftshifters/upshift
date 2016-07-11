package actions

import (
	"errors"
	"strings"
	"upshift/basher"
	"upshift/command"
	"upshift/config"
	"upshift/utils"
)

// Git : Strucut to handle everything related to git
type Git struct {
}

// AreSubmodulesUsed : Find out if the system uses git submodules
func (g *Git) AreSubmodulesUsed() bool {
	return utils.FileExists(".gitmodules")
}

// SubmoduleInit : Initialize submodules
func (g *Git) SubmoduleInit() (int, error) {
	var b basher.Basher
	utils.LogMessage("$ git submodule init")
	return b.RunAndTail("GitSubmoduleInit", []string{".upshift/logs/git-submodule-init.log"}, ".upshift/logs/git-submodule-init.log", []string{}, []string{"error"})
}

// SubmoduleUpdate : Update submodules
func (g *Git) SubmoduleUpdate() (int, error) {
	var b basher.Basher
	utils.LogMessage("$ git submodule update")
	return b.RunAndTail("GitSubmoduleUpdate", []string{".upshift/logs/git-submodule-update.log"}, ".upshift/logs/git-submodule-update.log", []string{}, []string{"error:", "fatal:"})
}

// IsRepo : Find out if you are in a git repo, if yes, throw the output of git status
func (g *Git) IsRepo() bool {
	_, err := command.Run([]string{"git", "status"}, "")
	if err != nil {
		return false
	}
	return true
}

// Branch : find the branch that you are on in the git repo
func (g *Git) Branch() (string, error) {
	// Run git status
	out, err := command.Run([]string{"git", "status"}, "")
	if err != nil {
		return "", err
	}

	// Read the first row of git status which says 'on branch xyz'
	gitStatusOutputRows := strings.Split(out, "\n")
	var firstRow string
	if len(gitStatusOutputRows) > 0 {
		firstRow = gitStatusOutputRows[0]
	} else {
		return "", errors.New("You are probably not in a git repository. Quit messing around.")
	}

	// Alright find the correct branch and show it to the user
	currentBranch := strings.TrimSpace(strings.Replace(firstRow, "On branch ", "", 1))
	return currentBranch, nil
}

// Remote : find the remote they are working with
func (g *Git) Remote() (string, error) {
	// Check how many remotes does the user have
	out, err := command.Run([]string{"git", "remote"}, "")
	if err != nil {
		return "", err
	}

	var currentRemote string
	gitRemoteOutputRows := strings.Split(strings.TrimSpace(out), "\n")
	switch len(gitRemoteOutputRows) {

	case 0:
		return "", errors.New("You have no remotes!")

	case 1:
		// If there's only one remote, use it
		currentRemote = strings.TrimSpace(gitRemoteOutputRows[0])
		return currentRemote, nil

	default:
		// This means that the user has multiple remotes, read the config to see if they have mentioned a remote there
		conf := config.Get()

		if conf.Settings.Remote == "" {
			// They didn't define a remote, throw a tantrum
			return "", errors.New("You have multiple repos in this project. To use one, please add Remote=remoteName to your config.toml. Here are the ones that we see " + strings.TrimSpace(out))
		}

		// Alright, so they have defined a remote, let's check if it exits in our list of remotes
		for _, row := range gitRemoteOutputRows {
			if strings.TrimSpace(conf.Settings.Remote) == strings.TrimSpace(row) {
				currentRemote = row
			}
		}

		// Didn't find their remote in git remotes, tell them so
		if currentRemote == "" {
			return "", errors.New("We can't find " + conf.Settings.Remote + " in this project")
		}

		return currentRemote, nil
	}
}

// Pull : pull a git repo
func (g *Git) Pull(remote string, branch string) (int, error) {
	if remote == "" {
		return 1, errors.New("Please select a remote that you want to pull from")
	}

	if branch == "" {
		return 1, errors.New("Please select a branch that you want to pull from")
	}

	var b basher.Basher
	utils.LogMessage("$ git pull " + remote + " " + branch)
	return b.RunAndTail("GitPull", []string{remote, branch, ".upshift/logs/git-pull.log"}, ".upshift/logs/git-pull.log", []string{}, []string{"fatal:", "error:"})
}
