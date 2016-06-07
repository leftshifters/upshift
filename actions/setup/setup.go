package setup

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"upshift/bash"
	"upshift/utils"
)

func init() {

}

func UpgradeScript() (int, bool) {
	resp, err := http.Get("https://raw.githubusercontent.com/leftshifters/upshift/master/release")
	if err != nil {
		fmt.Println("We were unable to find out the latest version", err.Error())
		return 1, false
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("We were unable to read the data from the server", err.Error())
		return 1, false
	}

	latestVersion := string(body)
	latestVersion = strings.TrimSpace(latestVersion)

	if latestVersion == utils.GetAppVersion() {
		fmt.Println("You are already at the latest version", utils.GetAppVersion())
		return 0, false
	}

	fmt.Println("We are now going to upgrade you to the latest version of upshift")

	status, err := bash.Run("upgradeScript", []string{})
	if err != nil {
		fmt.Println("We were unable to upgrade you", err.Error())
		return status, false
	}

	return 0, false
}

// function UpgradeVersion {

//   StartAction "UpgradeVersion"

//   # If 'next' is false, exit
//   if [ ${next} == false ]; then
//     ShowPreviousFailed
//     return
//   fi

//   LATEST_VERSION=$(curl -sS https://raw.githubusercontent.com/leftshifters/upshift/master/release 2>&1)
//   LATEST_VERSION_RESULTS=$(grep 'curl:' -c <<< "${LATEST_VERSION}")

//   if [ "${LATEST_VERSION_RESULTS}" -gt 0 ]; then
//     printf "There was a problem in confirming the version number.\nIgnoring the message and moving on\n"
//     printf "${LATEST_VERSION}\n\n"
//   else
//     if [ "${LATEST_VERSION}" == "${version}" ]; then
//       printf "You are using the latest version of upshift v${version}\n\n"
//     else
//       printf "Moving you to the latest version of upshift v${version}\n"
//       VERSION_UPGRADE=$(curl -fsSL https://raw.githubusercontent.com/leftshifters/upshift/master/upshift > upshift.temp && chmod +x upshift.temp && ./upshift.temp install)
//       printf "${VERSION_UPGRADE}\n"
//       rm upshift.temp
//     fi
//   fi

// }

func SetupConfig() (int, bool) {

	configExits := utils.FileExists("./config.toml")
	if configExits == true {
		fmt.Println("It looks like a config.toml is already here, skipping this step")
		return 1, false
	} else {
		// Config does not exist
		// Create a new config.toml in this directory

		sampleToml := `[Application]
Debug = false

[Runner]
RootPassword = "testPassword"

[Build]
GitRepoURL = "testRepo"
GitRepoBranch = "testBranch"
CleanBeforeBuild = false
UninstallOlderBuilds = false

[IOS]
ProjectName = "testProject"
UseWorkspace = false
Scheme = "testScheme"
TestDevice = "iPhone 6"
Xcode = "7.3.1"

[Android]
PackageName = "testPackage"
MainActivityName = "testActivity"`

		tomlBytes := []byte(sampleToml)

		err := ioutil.WriteFile("./config.toml", tomlBytes, 0644)
		if err != nil {
			fmt.Println("We could not write the config file, the OS told us this <" + err.Error() + ">")
			return 1, false
		}
	}

	fmt.Println("We just added a config.toml to this folder!")
	return 0, false
}

// function AddConfig {

//   StartAction "AddConfig"

//   # If 'next' is false, exit
//   if [ ${next} == false ]; then
//     ShowPreviousFailed
//     return
//   fi

//   # Check if config.ci exists in
//   if [ -f "config.ci" ]; then
//     printf "A config file ${redColour}already exists${noColour}\n"
//   else
//     # Write the config file
//     echo debug=false >> config.ci
//     echo alwaysCleanBeforeBuild=true >> config.ci
//     echo alwaysUninstallOlderBuilds=true >> config.ci
//     echo package=\'\' >> config.ci
//     echo mainActivity=\'\' >> config.ci
//     echo gitRepositoryURL=\'\' >> config.ci
//     echo gitRepositoryBranch=\'\' >> config.ci
//     echo masterPassword=\'\' >> config.ci
//     echo projectName=\'\' >> config.ci
//     echo useWorkspace=false >> config.ci
//     echo scheme=\'\' >> config.ci
//     echo iPhone=\'\' >> config.ci
//     echo xcodeVersion=\'\' >> config.ci

//     printf "We just add a very basic confi.ci file in this folder.\n"
//   fi

// }
