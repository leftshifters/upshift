# Upshift
The upshift utility helps you clone, build and test your iOS and Android projects.

# Installation
You can install upshift by running the following command on your mac
```
curl -fsSL https://raw.githubusercontent.com/leftshifters/upshift/master/upshift > upshift && upshift install
```

# Available Commands
The following options are currently available

### upshift android build
This command pulls the latest code, checks and install submodules and finally builds the Android project.

### upshift android emulator
This command starts up the emulator on the system, pulls the latest code, installs the git submodules and installs the app on the emulator.

### upshift ios build
This command checks if you are building on the correct Xcode version, pulls the code, installs submodules, installs pods, build the latest iOS project and deploys it on the iOS simulator

### upshift setup clone
This command helps you clone a new repository, install submodules and installs pods.

### upshift install
This command installs this script on your machine

### upshift -v
This command gets you the latest version number

# Compatibility
This software has only been tried and tested on Mac OSX 10.11.4. If you're planning to use this on windows, your're on your own!
