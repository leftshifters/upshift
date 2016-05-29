package setup

import (
	"io/ioutil"
	"log"
	"upshift/utils"
)

func init() {

}

func SetupConfig() (int, bool) {

	configExits := utils.FileExists("./config.toml")
	if configExits == true {
		log.Println("It looks like a config.toml is already here, skipping this step")
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
			log.Println("We could not write the config file, the OS told us this <" + err.Error() + ">")
			return 1, false
		}
	}

	log.Println("We just added a config.toml to this folder!")
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
