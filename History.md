
0.8.2 / 2016-06-25
==================

  * Make upshift executable and changed instructions on home screen
  * Fixed the install script [32m0s]
  * Merged with 0.8.1 [2m0s]
  * Merge branch 'release0.8.0'
  * Fixed issue with finding simulator path after building an archive [1h42m0s]

0.8.1 / 2016-06-17
==================

  * Removed upshift.sh [19m0s]
  * Removed android actions because google sucks [16m0s]
  * Setup a new action for upgrade android, Removed unwated logs [57m0s]
  * Run a pod repo update if it has never been run before of if has been over a month. If pod repo update says you need to update cocoapods, it updates it too. [12m0s]
  * Added a way to save defaults for upshift as a whole
  * Added action to update android sdk
  * created a makefile for cross compilation [23m0s]
  * Follow through redirects in install.sh [7m0s]
  * Fixed type in simulator loading [26m0s]

0.8.0 / 2016-06-10
==================

  * Bumped up version number to 0.8.0 [3m0s]
  * show version is now working [7m0s]
  * Updated README [4m0s]
  * Changed upgrade script to install install.sh [47m0s]
  * Picking up clean from config, ignoring offline devices on android [18m0s]
  * Added build for Android, launch emulator, find connected devices, find avds available [2h59m0s]
  * Setup to get started with Android, wrote all bash scripts first for android [1h13m0s]
  * Added the option to deploy the build to the simulator [15m0s]
  * Fixed typo with export ipa [1m0s]
  * Added a final goodbye emoji [9m0s]
  * Loading defaults properly if no config file is present [6m0s]
  * Add provisioning profiles before exporting IPA, fixed issue with extra new line in remotes [14m0s]
  * Added a step to export export.plist automatically [12m0s]
  * Added a check to see if exports.plist exists before you try to export IPA [9m0s]
  * Add option to export IPA, added provisioning profiles [47m0s]
  * Created action for archive iOS [21m0s]
  * Finished action for iOS build, it helps compile, fine available schemes, find if simulator devices are available, start simulator [2h49m0s]
  * Reading xcodeBuildSettings, getting the projectPath and checking if the simulator is open [1h16m0s]
  * Added action for setting up gradle wrapper [18m0s]
  * Decided to ignore git clone for now [14m0s]
  * Created activity to run pod install [28m0s]
  * Added action to install cocoapods, setup gem adding action [52m0s]
  * added comments for git pull and submodules [3m0s]
  * Added action for setting up submodules, updated logfile for git pull [22m0s]
  * Removed inspiration from git pull [1m0s]
  * Removed bash functions from ios.go [5m0s]
  * Added Git Pull action, Not switching Xcode if no config present, Automatically find the branches and remotes, Added a LogMessage for green logs, Added a function to read the last few bytes of a file [2h12m0s]
  * Started writing gitPull, found out the branch you are on, next up is finding the remote that you are on [45m0s]
  * Removed empty space from the output [15m0s]
  * SetupXCPretty is now working [4m0s]
  * Create bindata with SetupXcpretty [2m0s]
  * Temp commit to move bash files to the correct place [1h26m0s]
  * Undetected actions now exit with status 1 [2m0s]
  * Fixed issues with setupXcode [11m0s]
  * Added logs and fixed flow for reaching config [25m0s]
  * Added comments and removed inspiration from upshift in setup.go [7m0s]
  * Looked up setup config again [9m0s]
  * Handled skipNext and status return values [35m0s]
  * Changed the structure for color and updated it throughout the project [19m0s]
  * Handled good and bad cases of basher [22m0s]
  * Added bindata to package bash files, Setup basher to mix go and bash, Better messaging, Removed unwanted bash scripts [1h22m0s]
  * Added super awesome help, Added more loveable emojis [1h9m0s]
  * Added Bold, Light and Underline to Colours, Removed flavors in tasks, Added better formatting for messages, Added emoticons to match the output colours, Added show help for all cases where nothing matches [39m0s]
  * improved messaging for reading toml files [38m0s]
  * Added fmt instead of log and fixed up the utilities [2h41m0s]
  * Fixed issue with os