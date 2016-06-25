# Upshift
The upshift utility helps you clone, build and test your iOS and Android projects.

# Installation
You can install upshift by running the following command on your mac
```
curl -fsSL https://raw.githubusercontent.com/leftshifters/upshift/master/install.sh > install.sh && chmod +x install.sh && ./install.sh && rm ./install.sh
```

# Docker
You can find the upshift docker image for building android at [https://github.com/leftshifters/upshift-docker-android](https://github.com/leftshifters/upshift-docker-android)

# Available Commands
The following options are currently available

### upshift android build
This command pulls the latest code, checks and install submodules and finally builds the Android project.

### upshift ios build
This command checks if you are building on the correct Xcode version, pulls the code, installs submodules, installs pods, build the latest iOS project and deploys it on the iOS simulator

### upshift setup config
This command will create an empty config file in the current folder that you are in. Config files are setup in the main folder and are called config.ci . The variable defined in this file gets more priority than the ones defined inside upshift

### upshift action androidBuild
### upshift action iosBuild
### upshift action installPods
### upshift action gitSubmodules
### upshift action gitPull
### upshift action setupExportPlist
### upshift action setupConfig
### upshift action setupGradleW
### upshift action setupPods
### upshift action setupXcpretty
### upshift action setupXcode
### upshift action showHelp
### upshift action upgradeScript

These commands will allow you to run each specific action separately

### upshift -v
This command gets you the latest version number

# Compatibility
This software has only been tried and tested on Mac OSX 10.11.4. If you're planning to use this on windows, your're on your own!
