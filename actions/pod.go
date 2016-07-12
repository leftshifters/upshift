package actions

import (
	"github.com/leftshifters/upshift/basher"
	"github.com/leftshifters/upshift/utils"
)

// Pod : Handle everything related to install pods
type Pod struct{}

// IsInstalled : Check if cocoapods is installed
func (p *Pod) IsInstalled() bool {
	var c Cocoapods
	err := c.Install()
	if err != nil {
		utils.LogError("Could not install underlying cocoapods\n" + err.Error())
		return false
	}
	return true
}

// AreUsed : Check if this project uses pods
func (p *Pod) AreUsed() bool {
	return utils.FileExists("Podfile")
}

// Install : Install pods which are being used in the system
func (p *Pod) Install() (int, error) {
	var b basher.Basher
	utils.LogMessage("$ pod install")
	return b.RunAndTail("PodInstall", []string{".upshift/logs/pod-install.log"}, ".upshift/logs/pod-install.log", []string{}, []string{})
}

// RepoUpdate : Update the pod repo, not sure who will call it
func (p *Pod) RepoUpdate() (int, error) {
	var b basher.Basher
	utils.LogMessage("$ pod repo update --verbose")
	return b.RunAndTail("PodRepoUpdate", []string{".upshift/logs/pod-repo-update.log"}, ".upshift/logs/pod-repo-update.log", []string{}, []string{"error"})
}

// #TODO : Find out how do we know that an update is available and update it
// When an upgrade is available, they say
// CocoaPods 1.0.1 is available.
// To update use: `sudo gem install cocoapods`
// Until we reach version 1.0 the features of CocoaPods can and will change.
// We strongly recommend that you use the latest version at all times.
// if strings.Contains(tailData, "sudo gem install cocoapods") == true {
// This means that an update is available, run cocoapods update
//	status := SetupPods(true)
// if status > 0 {
// return errors.New("We couldn't update to the new version of cocoapods")
// }
// fmt.Println("Updated cocoapods to the latest version")
// }
