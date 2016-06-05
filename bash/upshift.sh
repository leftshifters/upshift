#!/usr/bin/env bash

# 1. Read the configuration files
debug=false
alwaysCleanBeforeBuild=true
alwaysUninstallOlderBuilds=true
package=""
mainActivity=""
gitRepositoryURL=""
gitRepositoryBranch=""
masterPassword=""
projectName=""
useWorkspace=false
scheme=""
iPhone="iPhone 6"
xcodeVersion="7.3"

# 2. Load up config from config file
if [ -f "./config.ci" ]; then
  # shellcheck disable=SC1091
  source ./config.ci
fi

#2.1 Load default values for variables which might have been null in config.ci
if [ "${xcodeVersion}" == "" ]; then
  xcodeVersion="7.3"
fi

if [ "${iPhone}" == "" ]; then
  iPhone="iPhone 6"
fi

# 3. Dump commands to the screen, only if one wants to debug
if [ "${debug}" == true ];then
  set -v
fi

# 4. Make sure things look good. Here are some font and color adjustments

# To avoid getting the "tput: No value for $TERM and no -T specified" error, export TERM in the script
termInEnv=$(printenv TERM)

if [ "${termInEnv}" == "" ]; then
  export TERM="xterm"
fi

redColour='\033[0;31m'
greenColour='\033[0;32m'
yellowColour='\033[0;33m'
blueColour='\033[0;34m'
grayColour='\033[0;90m'
whiteColour='\033[0;97m'
noColour='\033[0m'

grayBgColor='\033[0;100m'

boldStyle=$(tput bold)
normalStyle=$(tput sgr0)
underlineStyle=$(tput smul)
noUnderlineStyle=$(tput rmul)


# 5. Exit script on error
# set -e
# (Maye not)

# 6. Setup Global Variables
next=true
platform=$1
job=$2

# 7. Application Version
version="0.7.3"

# 8. Set time functions
startTime=$(date +%s)
endTime=""

# 9. Use the env variables to override the user defined ones
# Currently only reading masterPassword

# Test: export masterPasswordFromEnv="1231"
masterPasswordFromEnv=$(printenv masterPasswordFromEnv)

if [ "${masterPasswordFromEnv}" != "" ]; then
  masterPassword=${masterPasswordFromEnv}
  printf "\nSet the value of masterPassword from the environment\n\n"
fi

# 10. Find out if this is running via CI
# We could have multiple, so one variable for each type of service
GITLAB_CI=$(printenv GITLAB_CI)

# Overall, OR all of them to find out is it is running via CI
CI=${GITLAB_CI}

# 11. Find out if you are in docker or not
DOCKER_COUNT=$(cat /proc/1/cgroup | grep "docker" -c)

DOCKER=false
if [ "${DOCKER_COUNT}" -gt 0 ]; then
  DOCKER=true
fi



# Setup Internal Functions
# ⚡ ⚙

function StartScript {
  printf ""
  printf "\n${boldStyle}◀ BOOTING UP THE ENGINES ►${normalStyle}\n"
}

function EndScript {
  endTime=$(date +%s)
  totalTime=$((endTime - startTime))
  printf "This job took ${blueColour}${totalTime} seconds${noColour}\n"

  if [ "${next}" == false ]; then
    exit 1
  fi

}

function ShowError {
  printf "${boldStyle}ERROR ${normalStyle}\n"
}

function ShowPreviousFailed {
  printf "${redColour}Skipping${noColour} action because the previous actions failed\n"
}

function ShowCIFailed {
  printf "${redColour}Skipping${noColour} action because this is running via CI\n"
}

function StartAction {
  printf "\n${blueColour}◀ ACTION: $1 ►${noColour}\n\n"
}





# Setup Actions

##
## setup-ssh
##

function SetupSSH {

  StartAction "SetupSSH"

  # Details about the script came from here
  # https://help.github.com/articles/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent/

  # If 'next' is false, exit
  if [ ${next} == false ]; then
    ShowPreviousFailed
    return
  fi

  # Check if email has been defined by the user
  if [ "${emailForSSHKey}" != "" ]; then
    # Check if an id_rsa already exists at the defualt location
    if [ ! -f ~/.ssh/id_rsa ]; then
      printf "File does not exist at ~/.ssh/id_rsa"
      echo -e '\n' | ssh-keygen -t rsa -b 4096 -C "${emailForSSHKey}"

      # Show the created keys on the screen
      ID_RSA=$(<~/.ssh/id_rsa)
      ID_RSA_PUB=$(<~/.ssh/id_rsa.pub)

      printf "${boldStyle}id_rsa${normalStyle}\n"
      printf "${ID_RSA}"
      printf "\n\n${boldStyle}id_rsa.pub${normalStyle}\n"
      printf "${ID_RSA_PUB}"

      printf "All done 🍺\n"

    else
      ShowError
      printf "Can't do this, looks like an id_rsa already exists at ~/.ssh/id_rsa, get rid of that first\n"
      next=false
    fi
  else
    ShowError
    printf "Dude, you need to add the <${redColour}emailForSSHKey${noColour}> parameter to get this to work\nYou can get a sample config by running upshift setup config\n"
    next=false
  fi
}

##
## install-on-android
##

function InstallOnAndroid {

  StartAction "InstallOnAndroid"

  # If 'next' is false, exit
  if [ ${next} == false ]; then
    ShowPreviousFailed
    return
  fi

  # Make a TIMESHTAMP for the log file
  TIMESHTAMP=$(date +%Y%m%d%H%M%S)

  # If alwaysCleanBeforeBuild then run clean
  if [ "${alwaysCleanBeforeBuild}" == true ]; then
    printf "Time to ${greenColour}clean up${noColour}\n\n"
    ./gradlew clean 2>&1 | tee "gradle-clean-${TIMESHTAMP}.log"
  else
    printf "Skipping :clean:, you have ${blueColour}alwaysCleanBeforeBuild${noColour} turned off\n"
  fi

  # Uninstall older builds if the setting so desires
  if [ "${alwaysUninstallOlderBuilds}" == true ]; then
    printf "Time to ${greenColour}uninstall${noColour} older builds\n\n"
    ./gradlew uninstallAll 2>&1 | tee "gradle-uninstall-${TIMESHTAMP}.log"

  else
    printf "Skipping :uninstallAll:, you have ${blueColour}alwaysUninstallOlderBuilds${noColour} turned off\n"
  fi

  # Now time to build again
  printf "\nTime to run ${greenColour}installDebug${noColour} on this thing\n\n"
  ./gradlew installDebug --stacktrace 2>&1 | tee "gradle-install-${TIMESHTAMP}.log"

  # Get the logged results and try to make some sense out of it
  BUILD_RESULTS=$(<"gradle-install-${TIMESHTAMP}.log")

  # When you build via Gradle, it seems it always sends BUILD SUCCESSFUL in the results
  # This could mess up if in some build configuration, there are two messages, one is BUILD SUCCESSFUL and one otherwise
  BUILD_SUCCESSFUL=$(echo "${BUILD_RESULTS}" | grep "BUILD SUCCESSFUL" -c)

  # If the build was successful, let 'em know
  if [ "$BUILD_SUCCESSFUL" != "1" ]; then
    ShowError
    printf "Damn, it looks like something went ${redColour}wrong${noColour}. You should check this up.\n\n"
    next=false
  else
    printf "\n\n${greenColour}Super${noColour}! The build was fine.\n"
    # Check if package is empty
    if [ "${package}" != "" ];then
      if [ "${mainActivity}" != "" ]; then
        printf "Starting activity ${blueColour}${mainActivity}${noColour} in package ${blueColour}${package}${noColour}\n"

        # Start the activity and package
        adb shell am start -n ${package}/${package}.${mainActivity}

        # Tell the user everything is nice and easy
        printf "\nAlright, the build was ${greenColour}successful${noColour} 🍺\n\n"
      else
        # The mainActivity is empty it seems
        printf "Alright, the build was ${greenColour}successful${noColour}, but there was no ${blueColour}mainActivity${noColour} defined, so couldn't start it automatically 🍺\n\n"
      fi
    else
      # The package is empty it seems
      printf "Alright, the build was ${greenColour}successful${noColour}, but there was no ${blueColour}package${noColour} defined, so couldn't start it automatically 🍺\n\n"
    fi
  fi
}

##
## git-pull
##

function GitPull {

  StartAction "GitPull"

  # If this is coming via CI, don't pull
  if [ "${CI}" == true ]; then
    ShowCIFailed
    return
  fi

  # If 'next' is false, exit
  if [ ${next} == false ]; then
    ShowPreviousFailed
    return
  fi

  # Make a TIMESHTAMP for log file
  TIMESHTAMP=$(date +%Y%m%d%H%M%S)

  # Check if the branch name is defined
  if [ "${gitRepositoryBranch}" != "" ]; then

    printf "Alright, let's ${greenColour}pull${noColour} the ${gitRepositoryBranch} branch for this repo\n\n"

    # Alright, let's pull
    git pull origin ${gitRepositoryBranch} 2>&1 | tee "git-pull-${TIMESHTAMP}.log"

    # Load up the results to see if there were any errors
    PULL_RESULTS=$(<"git-pull-${TIMESHTAMP}.log")
    PULL_RESULTS_FATAL=$(grep "fatal:" -c <<< "${PULL_RESULTS}")
    PULL_RESULTS_ERROR=$(grep "error:" -c <<< "${PULL_RESULTS}")
    # If there was a fatal error, tell the user there's something wrong
    if [ "${PULL_RESULTS_FATAL}" -gt "0" ] || [ "${PULL_RESULTS_ERROR}" -gt "0" ]; then
      ShowError
      printf "Something went wrong with the pull, you should look this up\n\n"
      next=false
    else
      # All done
      printf "\nAll done ${greenColour}baby${noColour}! 🍺.\n\n"
    fi

  else
    # The user hasn't added the required keys
    ShowError
    printf "Dude, you need to add the ${blueColour}gitRepositoryBranch${noColour} for this to work\nYou can get a sample config by running upshift setup config\n\n"
    next=false
  fi
}

##
## git-clone
##

function GitClone {

  StartAction "GitClone"

  # If 'next' is false, exit
  if [ ${next} == false ]; then
    ShowPreviousFailed
    return
  fi

  # Make a TIMESHTAMP for log file
  TIMESHTAMP=$(date +%Y%m%d%H%M%S)

  # Check if the repo URL is defined
  if [ "${gitRepositoryURL}" != "" ]; then
    # Check if the branch name is defined
    if [ "${gitRepositoryBranch}" != "" ]; then

      # Check if this is already a git repository

      if [ -f ".git/config" ]; then

        printf "You have a .git/config, I believe you already have a cloned repo here!\n"

      else

        printf "Alright, let's ${greenColour}clone${noColour} the ${blueColour}${gitRepositoryBranch}${noColour} branch for the ${blueColour}${gitRepositoryURL}${noColour} repo\n\n"

        # Alright, let's pull
        # But you can't pull into an empty directly, now since you already have bitrise.yml and .bitrise.secrets.yml in your directory,
        #   you will need to clone into another folder, and move stuff back here
        #   we can't do what the rest of the world tries to do - which is git init, add remote,
        #   because we want to ensure we do depth=1 and not download the whole repo, which can be painful at times
        git clone -b "${gitRepositoryBranch}" "${gitRepositoryURL}" tmp --progress --depth 1  2>&1 | tee "git-clone-${TIMESHTAMP}.log"
        mv tmp/* . 2>&1 | tee "git-clone-${TIMESHTAMP}.log"
        mv tmp/.* . 2>/dev/null | tee "git-clone-${TIMESHTAMP}.log"
        rm -rf tmp/ 2>&1 | tee "git-clone-${TIMESHTAMP}.log"

        # Load up the results to see if there were any errors
        CLONE_RESULTS=$(<"git-clone-${TIMESHTAMP}.log")
        # If there was a fatal error, tell the user there's something wrong
        if [ "$(printf "${CLONE_RESULTS}" | grep "fatal" -c)" -gt "0" ] || [ "$(printf "${CLONE_RESULTS}" | grep "error" -c)" -gt "0" ]; then
          ShowError
          printf "Something failed while we were cloning, you should look this up\n\n"
          next=false
        else
          # All done
          printf "\nAll done ${greenColour}baby${noColour}! 🍺.\n\n"
        fi

      fi

    else
      # The user hasn't added the required keys
      ShowError
      printf "Dude, you need to add the ${blueColour}gitRepositoryBranch${noColour} for this to work\nYou can get a sample config by running upshift setup config\n\n"
      next=false
    fi
  else
    # The user hasn't added the required keys
    ShowError
    printf "Dude, you need to add the ${blueColour}gitRepositoryURL${noColour} for this to work\nYou can get a sample config by running upshift setup config\n\n"
    next=false
  fi
}

##
## start-emulator
##

function StartEmulator {

  StartAction "StartEmulator"

  # If 'next' is false, exit
  if [ ${next} == false ]; then
    ShowPreviousFailed
    return
  fi

  # Check if Boot Animation is still on
  # https://devmaze.wordpress.com/2011/12/12/starting-and-stopping-android-emulators/
  # adb shell getprop init.svc.bootanim
  # We don't really care about this right now

  # Check if a process which calls itself the emulator is running
  OUTPUT=$(ps -aef | grep emulator | grep "sdk/tools" -c)
  # If 0 processes are called emulator, it means we need to load up one
  if [ "$OUTPUT" == "0" ]; then

    if [ "${emulatorName}" != "" ]; then

      EMULATOR_RESULTS=$(nohup "$ANDROID_HOME/tools/emulator" -avd "${emulatorName}" 2>emulator.log 1>emulator.log &)
      # TODO : This is a big #HACK, only errors are returned in the first two seconds, I suck and I don't know a way out
      # TODO : Another potential problem, we redirect both 1,2 in reset mode (>), the file could get overwritten
      sleep 2
      EMULATOR_RESULTS=$(<emulator.log)

      # Check if there was a PANIC [to test this, put in the wrong emulator name]
      PANIC_COUNT=$(echo "${EMULATOR_RESULTS}" | grep "PANIC" -c)

      # If there was a panic, throw it at the user, they messed up, not us
      if [ "${PANIC_COUNT}" -gt 0 ]; then
        ShowError
        printf "The emulator won't load up. Maybe the ${redColour}emulatorName${noColour} key isn't correct\n"
        printf "Here's what Mr. Log says:\n\n"
        printf "${EMULATOR_RESULTS}\n\n"
        next=false
      else
        # Seems like there is no panic, let's check for errors
        ERROR_COUNT=$(echo "${EMULATOR_RESULTS}" | grep "ERROR" -c)

        if [ "${ERROR_COUNT}" -gt 0 ]; then
          printf "${EMULATOR_RESULTS}\n\n"
          next=false
        else
          printf "All set baby, ${greenColour}starting${noColour} to load up the emulator\n"
          # ADB gives this option to wait for the device till it comes up, let's just depend on it,
          # this is really mess with us when there is a problem with the emulator fails to load because of it's own bugs
          adb wait-for-device

          # Now that the device is available, the worst is over
          # Check if the emulator has finished botting, if not, sleedp for 2 seconds and try this again, this is our final trigger
          printf "Seems like the device is now ${greenColour}available${noColour}, we are getting close\n"
          SCREEN_LOADING=$(adb shell getprop sys.boot_completed | tr -d '\r')
          while [ "$SCREEN_LOADING" != "1" ]; do
            sleep 4
            printf "Check if the emulator has finished booting, why is this thing so ${blueColour}damn${noColour} slow?\n"
            SCREEN_LOADING=$(adb shell getprop sys.boot_completed | tr -d '\r')
          done

          # Alright, everything is now done. This is just used to unlock the device.
          printf "Almost ${greenColour}done${noColour}, touch the device once\n"
          adb shell input keyevent 82
          printf "${greenColour}Super!${noColour} The emulator is now running. You're one lucky person 🍺\n"
        fi
      fi

    else
      ShowError
      printf "Dude, you need to define ${blueColour}emulatorName${noColour} for this to work\nYou can get a sample config by running upshift setup config\n\n"
      next=false
    fi
  else
    printf "${greenColour}Dude${noColour}, Looks like the emulator is already running!\n\n"
  fi
}

##
## android-devices
##
function AndroidDevices {

  StartAction "AndroidDevices"

  # If 'next' is false, exit
  if [ ${next} == false ]; then
    ShowPreviousFailed
    return
  fi

  # Touch device to that they can get unlocked, otherwise ADB just ignores them
  # Just make sure you hide the STDERR because well we don't care too much about it
  TOUCH_DEVICE=$(adb -d shell input keyevent 82 2>&1)

  # Find out how many devices are available, and are not emulators
  DEVICES_FOUND=$(adb devices | grep -v "List of devices attached" | grep -v emulator -c)
  # For some reason we get one added to the value, so let's just substract it
  let DEVICES_FOUND=DEVICES_FOUND-1;

  # Show appropriate error message
  if [ "${DEVICES_FOUND}" -gt 0 ]; then
    printf "${DEVICES_FOUND} physical device(s) were found 🍺\n\n"
  else
    ShowError
    printf "No physical devices were found\n\n"
    next=false
  fi
}

##
## assemble-android
##
function AssembleAndroid {

  StartAction "AssembleAndroid"

  # If 'next' is false, exit
  if [ ${next} == false ]; then
    ShowPreviousFailed
    return
  fi

  # Make a TIMESHTAMP for log file
  TIMESHTAMP=$(date +%Y%m%d%H%M%S)

  # If alwaysCleanBeforeBuild then run clean
  if [ "${alwaysCleanBeforeBuild}" == true ]; then
    printf "Time to ${greenColour}clean up${noColour}\n\n"
    ./gradlew clean 2>&1 | tee "gradle-clean-${TIMESHTAMP}.log"
  else
    printf "Skipping :clean:, you have ${blueColour}alwaysCleanBeforeBuild${noColour} turned off\n"
  fi

  # Now time to build again
  printf "\nTime to run ${greenColour}assemble${noColour} on this thing\n\n"
  ./gradlew assemble --stacktrace 2>&1 | tee "gradle-assemble-${TIMESHTAMP}.log"

  # Get the logged results and try to make some sense out of it
  BUILD_RESULTS=$(<"gradle-assemble-${TIMESHTAMP}.log")

  # When you build via Gradle, it seems it always sends BUILD SUCCESSFUL in the results
  # This could mess up if in some build configuration, there are two messages, one is BUILD SUCCESSFUL and one otherwise
  BUILD_SUCCESSFUL=$(echo "$BUILD_RESULTS" | grep "BUILD SUCCESSFUL" -c)

  # If the build was successful, let 'em know
  if [ "$BUILD_SUCCESSFUL" != "1" ]; then
    ShowError
    printf "Damn, the build was not ${redColour}successful${noColour}. You should check this up.\n\n"
    next=false
  else
    printf "\nAlright, the build was ${greenColour}successful${noColour} 🍺\n\n"
  fi
}

##
## Setup Submodule if they exist
##

function GitSubmodules {

  StartAction "GitSubmodules"

  # If 'next' is false, exit
  if [ ${next} == false ]; then
    ShowPreviousFailed
    return
  fi

  # Make a TIMESHTAMP for log file
  TIMESHTAMP=$(date +%Y%m%d%H%M%S)

  # Check if .gitmodules exists
  if [ -f ".gitmodules" ]; then
    # If the file exists, we need to run init and update and catch errors
    git submodule init 2>&1 | tee "git-submodule-init-${TIMESHTAMP}.log"
    git submodule update 2>&1 | tee "git-submodule-update-${TIMESHTAMP}.log"

    SUBMODULE_RESULTS=$(<"git-submodule-update-${TIMESHTAMP}.log")

    if [ $(echo "${SUBMODULE_RESULTS}" | grep "fatal:" -c) -gt 0 ] || [ $(echo "${SUBMODULE_RESULTS}" | grep "error:" -c) -gt 0 ]; then
      ShowError
      printf "Damn, initializing submodules was ${redColour}not successful${noColour}. You should check this up.\n\n"
      next=false
    else
      printf "\nSubmodules are now ${greenColour}setup${noColour}, one less thing to think about! 🍺\n\n"
    fi
    # Else Quietly ignore
  else
    printf "\nIt looks like this project doesn't use ${greenColour}submodules${noColour}.\n\n"
  fi
}

##
## Install pods if they exist
##

function SetupPods {

  StartAction "SetupPods"

  # If 'next' is false, exit
  if [ ${next} == false ]; then
    ShowPreviousFailed
    return
  fi

  # Make a TIMESHTAMP for log file
  TIMESHTAMP=$(date +%Y%m%d%H%M%S)

  # Check if Podfile exits
  if [ -f "Podfile" ]; then
    # Check if cocoapods is installed
    POD_VERSION=$(pod --version 2>&1)
    POD_INSTALLED=$(grep 'command not found' -c <<< "${POD_VERSION}")

    if [ "${POD_INSTALLED}" -gt 0 ]; then
      # Cocoapods is not installed, let's install it first

      if [ "${CI}" == true ]; then
        # If we are on CI, we need password
        if [ "${masterPassword}" == "" ]; then
          ShowError
          printf "Alright, so it seems we need to install cocoapods and that requires\nadmin permissions. You need to add  ${redColour}masterPassword${noColour} to your config\nfor this to work.\nYou can get a sample config by running upshift setup config\n"
          next=false
          exit 1
        else
          # This means we are on CI and password exists
          echo -e "${masterPassword}" | sudo -S gem install cocoapods
        fi
      else
        # This means we are NOT on CI and we don't care about the password
        sudo gem install cocoapods
        
      fi
    fi

    # https://guides.cocoapods.org/using/pod-install-vs-update.html
    # We want to keep pods on their own version, hence not updating
    pod install 2>&1 | tee "pod-install-${TIMESHTAMP}.log"
    POD_INSTALL=$(<"pod-install-${TIMESHTAMP}.log")
    POD_INSTALL_FAILED=$(grep "Invalid" -c <<< "${POD_INSTALL}")

    if [ "${POD_INSTALL_FAILED}" -gt 0 ]; then
      ShowError
      printf "Damn, pod install ${redColour}not successful${noColour}. You should check this up.\n\n"
      next=false
    else
      printf "\nPods are now ${greenColour}up to date${noColour}, one less thing to think about! 🍺\n\n"
    fi      
  else
    printf "\nIt looks like this project doesn't use ${greenColour}pods${noColour}. You're awesome!\n\n"
  fi
}

##
## Get the XCode version in use
##
function XCodeVersion {

  StartAction "XCodeVersion"

  XCODE_VERSION=$(xcodebuild -version | grep "Xcode" | tr -d "Xcode ")

  if [ "${XCODE_VERSION}" != "" ]; then

    printf "We are currently using XCode ${blueColour}${XCODE_VERSION}${noColour}\n\n"

    if [ "${xcodeVersion}" != "" ]; then

      # Alright now we check if the versions match
      if [ "${XCODE_VERSION}" != "${xcodeVersion}" ]; then
        # Looks like the xcode version required and available do not match
        # Check if this system has the XCode version required
        # TODO : This will vary based on how you setup XCode, find out if there is a better way
        printf "We expect XCode versions to be placed like this\n/Applications/Xcode-7.2.app\n/Applications/Xcode-7.3.app\n\n"
        if [ -d "/Applications/Xcode-${xcodeVersion}.app/" ]; then
          # Looks like this version is available on this machine

          printf "${blueColour}Switching${noColour} Xcode versions\n\n"
          if [ "${CI}" == true ]; then
            # If we are on CI, expect masterPassword
            if [ "${masterPassword}" != "" ]; then
              echo -e "${masterPassword}" | sudo xcode-select -switch "/Applications/Xcode-${xcodeVersion}.app/"
            else
              ShowError
              printf "Alright, so we need the sudo password to change the Xcode version.\nWould you be a sweetheart and add it\nto the ${blueColour}masterPassword${noColour} key in the config\n\n"
              next=false
              exit 1          
            fi
          else
            # If we are NOT on CI, don't expect password
            sudo xcode-select -switch "/Applications/Xcode-${xcodeVersion}.app/"
          fi

          # Maye it's done, check and confirm
          XCODE_VERSION=$(xcodebuild -version | grep "Xcode" | tr -d "Xcode ")

          if [ "${XCODE_VERSION}" == "${xcodeVersion}" ]; then
            printf "We are now using XCode ${blueColour}${XCODE_VERSION}${noColour}\n"
          else
            ShowError
            printf "Something went wrong. Maybe we messed up big time or the password that you entered was wrong.\n\n"
            next=false
          fi
        else
          ShowError
          printf "We do not have XCode ${xcodeVersion} on this machine. About time to get it.\n\n"
          next=false
        fi
      fi

    else
      printf "Your config shows that you are not picky about the\n${blueColour}xcodeVersion${noColour}. We will use whatever is available!\n\n"
    fi

  else 
    ShowError
    printf "We do not have ${redColour}XCode${noColour} installed on this machine. About time to get it.\n\n"
    next=false
  fi

}

##
## Install xcpretty
##
function XCPretty {

  StartAction "XCPretty"

  # If 'next' is false, exit
  if [ ${next} == false ]; then
    ShowPreviousFailed
    return
  fi

  XCPRETTY_VERSION=$(xcpretty --version 2>&1)
  XCPRETTY_INSTALLED=$(grep 'command not found' -c <<< "${XCPRETTY_VERSION}")

  if [ "${XCPRETTY_INSTALLED}" -gt 0 ]; then
    # XCPretty is not installed, let's install it first

    if [ "${CI}" == true ]; then
      # CI is true, we now need password
      if [ "${masterPassword}" == "" ]; then
        # masterPassword exists, do your thing
        echo -e "${masterPassword}" | sudo -S gem install xcpretty
      else
        # Show error and stop
        ShowError
        printf "Alright, so it seems we need to install xcpretty and that requires\nadmin permissions. You need to add  ${redColour}masterPassword${noColour} to your config\nfor this to work.\nYou can get a sample config by running upshift setup config"
        next=false
      fi
    else
      # User should type in the password
      sudo -S gem install xcpretty
      printf "\nXCPretty is now ${greenColour}installed${noColour}, one less thing to think about! 🍺\n\n"
    fi

  else
    printf "Woot! XCPretty is already installed\n"
  fi
}

##
## iOS Build Project
##
function BuildiOS {

  StartAction "BuildiOS"

  # If 'next' is false, exit
  if [ ${next} == false ]; then
    ShowPreviousFailed
    return
  fi

  # Make a TIMESHTAMP for log file
  TIMESHTAMP=$(date +%Y%m%d%H%M%S)

  # Either use projectName defined by the user, or pick it automatically
  if [ "${projectName}" == "" ]; then
    printf "Dude, you need to define ${redColour}projectName${noColour} in your config.\nSince you haven't, we are going to pick it automatically,\nwhich will take time on every build.\n\n"

    PROJECT_NAME_STRING=$(xcodebuild -showBuildSettings | grep PROJECT_NAME)
    projectName=${PROJECT_NAME_STRING#"    PROJECT_NAME = "}

    if [ "${projectName}" == "" ]; then
      ShowError
      printf "Damn, couldn't even find it automatically. Are you sure this is an iOS repo?\n"
      next=false
      return
    fi

    printf "Found ${blueColour}${projectName}${noColour} automatically. Using this now.\n\n"
  fi

  # Build using workspace if user asks, if it uses cocoapods user workspace automatically, othewise use xcodeproj
  PROJECT_TYPE="project"
  EXTENSION=".xcodeproj"

  if [ "${useWorkspace}" == true ]; then
    # Since the user is requesting for it, decision is done, we love our users.
    PROJECT_TYPE="workspace"
    EXTENSION=".xcworkspace"
  else
    if [ -f "Podfile" ]; then
      # If a Podfile exits, then guys use Cocoapods, load up the workspace by default
      PROJECT_TYPE="workspace"
      EXTENSION=".xcworkspace"
    fi
  fi

  PROJECT_PATH="${projectName}${EXTENSION}";

  # Check if we have `iPhone` variable set
  if [ "${iPhone}" == "" ]; then
    ShowError
    printf "Looks like you forgot to set iPhone in config. Do that first and come back.\n\nExample iPhone=\'${blueColour}iPhone 6\'${noColour}\n"
    next=false
    return
  fi

  if [ "${scheme}" != "" ];then

    # Load up the simulator first, so that it gets ready while the build happens
    # Check if the simulator is already open
    OUTPUT=$(ps -aef | grep "Simulator.app" -c)
    if [ "${OUTPUT}" -gt 1 ]; then
      # There's a simulator already running
      printf "The simulator is already ${greenColour}running${noColour}!\n\n"
    else
      # Load up the simulator

      # Find the correct simulator
      # From here - https://en.wikipedia.org/wiki/Xcode - Xcode 7.0 - 7.x (since Swift 2.0 support)
      iPhoneOS=""
      case "${xcodeVersion}" in
        "7.3.1") iPhoneOS="(9.3)";;
        "7.3") iPhoneOS="(9.3)";;
        "7.2.1") iPhoneOS="(9.2)";;
        "7.2") iPhoneOS="(9.2)";;
        "7.1.1") iPhoneOS="(9.1)";;
        "7.1") iPhoneOS="(9.1)";;
        "7.0.1") iPhoneOS="(9.0)";;
        "7.0") iPhoneOS="(9.0)";;
        *)
          ShowError
          printf "The device that you wanted to load [${iPhone} ${iPhoneOS}] does not exit\n\n"
          next=false
          return
        esac

      # Find out if the target exists on this device on note
      printf "Getting the list of available simulators\n"
      DEVICES_AVAILABLE=$(instruments -s devices)
      DEVICE_FOUND=$(grep "${iPhone} ${iPhoneOS}" -c <<< "${DEVICES_AVAILABLE}" )

      if [ "${DEVICE_FOUND}" -gt 0 ]; then
        printf "Starting up the ${greenColour}simulator${noColour}!\n\n"
        xcrun instruments -w "${iPhone} ${iPhoneOS}" 1>/dev/null 2>&1 
      else
        ShowError
        printf "The device that you wanted to load [${iPhone} ${iPhoneOS}] does not exit\n\n"
        next=false
        return
      fi

    fi

    # Build the effing thing
    # TODO : Clean the effing thing before you start
    printf "Compiling the beautiful codebase\n\n"

    set -o pipefail && xcodebuild -"${PROJECT_TYPE}" "${PROJECT_PATH}" -scheme "${scheme}" -hideShellScriptEnvironment -sdk iphonesimulator -destination "platform=iphonesimulator,name=${iPhone}" -derivedDataPath build | tee "xcode-build-${TIMESHTAMP}.log" | xcpretty

    BUILD_RESULTS=$(<"xcode-build-${TIMESHTAMP}.log");
    BUILD_SUCCEDED=$(grep "BUILD SUCCEEDED" -c <<< "${BUILD_RESULTS}")

    if [ "${BUILD_SUCCEDED}" -gt 0 ]; then
      # The build succeded
      printf "The build was ${greenColour}successful${noColour} 🍺\n\n"
    else
      ShowError
      printf "It seems the build ${redColour}failed${noColour}. You need to look this up\n\n"
      next=false
    fi

  else
    ShowError
    printf "Dude, you need to define the ${blueColour}scheme${noColour} that you would like to build.\nYou can get a sample config by running upshift setup config\nYou can pick one here\n\n"
    xcodebuild -list
    next=false
  fi
}

##
## Archive iOS
##

function ArchiveIOS {

  StartAction "ArchiveIOS"

  # If 'next' is false, exit
  if [ ${next} == false ]; then
    ShowPreviousFailed
    return
  fi

  # Make a TIMESHTAMP for log file
  TIMESHTAMP=$(date +%Y%m%d%H%M%S)

  # Either use projectName defined by the user, or pick it automatically
  if [ "${projectName}" == "" ]; then
    printf "Dude, you need to define ${redColour}projectName${noColour} in your config.\nSince you haven't, we are going to pick it automatically,\nwhich will take time on every build.\n\n"

    PROJECT_NAME_STRING=$(xcodebuild -showBuildSettings | grep PROJECT_NAME)
    projectName=${PROJECT_NAME_STRING#"    PROJECT_NAME = "}

    if [ "${projectName}" == "" ]; then
      ShowError
      printf "Damn, couldn't even find it automatically. Are you sure this is an iOS repo?\n"
      next=false
      return
    fi

    printf "Found ${blueColour}${projectName}${noColour} automatically. Using this now.\n\n"
  fi

  # Build using workspace if user asks, if it uses cocoapods user workspace automatically, othewise use xcodeproj
  PROJECT_TYPE="project"
  EXTENSION=".xcodeproj"

  if [ "${useWorkspace}" == true ]; then
    # Since the user is requesting for it, decision is done, we love our users.
    PROJECT_TYPE="workspace"
    EXTENSION=".xcworkspace"
  else
    if [ -f "Podfile" ]; then
      # If a Podfile exits, then guys use Cocoapods, load up the workspace by default
      PROJECT_TYPE="workspace"
      EXTENSION=".xcworkspace"
    fi
  fi

  PROJECT_PATH="${projectName}${EXTENSION}";

  if [ "${scheme}" != "" ]; then

    set -o pipefail && xcodebuild -"${PROJECT_TYPE}" "${PROJECT_PATH}" -scheme "${scheme}" -derivedDataPath build -archivePath "build/${projectName}.xcarchive" archive | tee "xcode-archive-${TIMESHTAMP}.log" | xcpretty

    ARCHIVE_RESULTS=$(<"xcode-archive-${TIMESHTAMP}.log");
    ARCHIVE_SUCCEDED=$(grep "ARCHIVE SUCCEEDED" -c <<< "${ARCHIVE_RESULTS}")

    if [ "${ARCHIVE_SUCCEDED}" -gt 0 ]; then
      # The archive succeded

      # Get the UUID from profiles
      # https://gist.github.com/mxpr/8208289a63ca4e3a35a4
      # Loop through all files, if you get a UDID, add them to the list of profiles
      if [ -d "./profiles/" ]; then
        # The profiles folder exists
        for profileName in profiles/*.mobileprovision; do
          uuid=$(/usr/libexec/PlistBuddy -c 'Print UUID' /dev/stdin <<< $(security cms -D -i "${profileName}" 2>/dev/null))

          # If a UUID exist, then copy it, if it hasn't already been copied
          if [ "${uuid}" != "" ]; then
            # Copy file to UUID folder
            `cp -f ${profileName} ~/Library/MobileDevice/Provisioning\ Profiles/${uuid}.mobileprovision`
            printf "Copy provisioning file ${uuid}.mobileprovision to the system\n"
          fi

        done
      else
        ShowError
        printf "You need to add your provisioning profiles to a folder called profiles\n\n"
        next=false
        exit
      fi


      # Create an IPA (Export the effing thing)
      printf "Create an IPA\n"
      set -o pipefail && xcodebuild -exportArchive -exportOptionsPlist "profiles/export.plist" -archivePath "build/${projectName}.xcarchive" -exportPath "build/${projectName}.ipa" 2>&1 | tee "xcode-archive-${TIMESHTAMP}.log" | xcpretty

      EXPORT_RESULTS=$(<"xcode-archive-${TIMESHTAMP}.log")
      EXPORT_SUCCEDED=$(grep "EXPORT SUCCEEDED" -c <<< "${EXPORT_RESULTS}")

      if [ "${EXPORT_SUCCEDED}" -gt 0 ]; then
        printf "The build was ${greenColour}successful${noColour} 🍺\n\n"
      else
        ShowError
        printf "It seems the export ${redColour}failed${noColour}. You need to look this up\n\n"
        next=false
      fi
    else
      ShowError
      printf "It seems the build ${redColour}failed${noColour}. You need to look this up\n\n"
      next=false
    fi
  else
    ShowError
    printf "Dude, you need to define the ${blueColour}scheme${noColour} that you would like to build.\nYou can get a sample config by running upshift setup config\nYou can pick one here\n\n"
    xcodebuild -list
    next=false
  fi

}

##
## Deploy iOS to a simulator
##
function DeployiOSSimulator {

  StartAction "DeployiOSSimulator"

  # If 'next' is false, exit
  if [ ${next} == false ]; then
    ShowPreviousFailed
    return
  fi

  # Get the product full name
  FULL_PRODUCT_NAME=$(xcodebuild -showBuildSettings | grep FULL_PRODUCT_NAME)
  fullProductName=${FULL_PRODUCT_NAME#"    FULL_PRODUCT_NAME = "}

  # Get the bundle identifier
  PRODUCT_BUNDLE_IDENTIFIER=$(xcodebuild -showBuildSettings | grep PRODUCT_BUNDLE_IDENTIFIER)
  productBundleIdentifier=${PRODUCT_BUNDLE_IDENTIFIER#"    PRODUCT_BUNDLE_IDENTIFIER = "}

  if [ "${fullProductName}" != "" ] && [ "${productBundleIdentifier}" != "" ]; then

    printf "About to ${blueColour}deploy${noColour} ${fullProductName} (${productBundleIdentifier}) to the simulator\n\n"

    # TODO : Find a good way to delete the app from the simulator
    #xcrun simctl uninstall booted ${productBundleIdentifier}

    #Details here http://dduan.net/post/2015/02/build-and-run-ios-apps-in-commmand-line/

    # Find the path of the build
    # Check if Debug-iphonesimulator path exits
    BUILD_PATH=""
    if [ -f "./build/Build/Products/Debug-iphonesimulator/${fullProductName}" ]; then
      # All okay, just use this
      BUILD_PATH="./build/Build/Products/Debug-iphonesimulator/${fullProductName}\n"
    else 
      # So that doesn't exits, damn, let's find the folder
      PROBABLE_PATH=$(find . | grep "/${fullProductName}$")

      if [[ -d "${PROBABLE_PATH}" && ! -L "${PROBABLE_PATH}" ]]; then
        # Alright so the path that we found exits
        BUILD_PATH=${PROBABLE_PATH}
      else
        # Couldn't find the path, screw this
        ShowError
        printf "Damn. Could't find where the app was built. Time to shut down\n"
        next=false
        return
      fi

    fi

    # Deploy the app to the simulator
    xcrun simctl install booted "${BUILD_PATH}"

    # Open up the app on the simulator
    printf "Starting the app\n\n"
    xcrun simctl launch booted "${productBundleIdentifier}"
  else
    ShowError
    printf "It looks like this is not an iOS repository. Do you know what's going on?\n\n"
    next=false
  fi

}

##
## Setup script
##
function SetupScript {

  StartAction "SetupScript"

  # If 'next' is false, exit
  if [ ${next} == false ]; then
    ShowPreviousFailed
    return
  fi

  # Only these are available in $PATH on a fresh system
  # /usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin
  # We will install in /usr/local/bin

  # Check if /usr/local/upshift directory exists, if not create it
  if [ ! -d "/usr/local/upshift/${version}" ]; then
    # Directory doesn't exist, create it

    # If we are on docker, don't need sudo because there is no password required
    if [ "${DOCKER}" == true ]; then
      mkdir -p /usr/local/upshift/${version}
    else
      sudo mkdir -p /usr/local/upshift/${version}
    fi
    
    #printf "Setup the folder in /usr/local/upshift/${version}\n"
  fi

  # Throw yourself into the folder above, it not already there
  if [ ! -f "/usr/local/upshift/${version}/upshift" ]; then
    # Copy away

    # If we are on docker, don't need sudo because there is no password required
    if [ "${DOCKER}" == true ]; then
      cp -rf ./upshift.temp /usr/local/upshift/${version}/upshift
    else
      sudo cp -rf ./upshift.temp /usr/local/upshift/${version}/upshift
    fi

    
    #printf "Copied the script into /usr/local/upshift/${version}/upshift\n"

    # Now add a link to the above file in /usr/local/bin
    if [ -L "/usr/local/bin/upshift" ]; then
      # Remove only if it exits

      # If we are on docker, don't need sudo because there is no password required
      if [ "${DOCKER}" == true ]; then
        rm -f /usr/local/bin/upshift
      else
        sudo rm -f /usr/local/bin/upshift
      fi
      
    fi

    # If we are on docker, don't need sudo because there is no password required
    if [ "${DOCKER}" == true ]; then
      ln -s /usr/local/upshift/${version}/upshift /usr/local/bin
    else
      sudo ln -s /usr/local/upshift/${version}/upshift /usr/local/bin
    fi

    printf "Installation has been ${greenColour}successfully${noColour} completed\n"
  else
    printf "This version is already ${greenColour}installed${noColour}\n"
  fi


}

##
## Add Configuration File
##
function AddConfig {

  StartAction "AddConfig"

  # If 'next' is false, exit
  if [ ${next} == false ]; then
    ShowPreviousFailed
    return
  fi

  # Check if config.ci exists in
  if [ -f "config.ci" ]; then
    printf "A config file ${redColour}already exists${noColour}\n"
  else
    # Write the config file
    echo debug=false >> config.ci
    echo alwaysCleanBeforeBuild=true >> config.ci
    echo alwaysUninstallOlderBuilds=true >> config.ci
    echo package=\'\' >> config.ci
    echo mainActivity=\'\' >> config.ci
    echo gitRepositoryURL=\'\' >> config.ci
    echo gitRepositoryBranch=\'\' >> config.ci
    echo masterPassword=\'\' >> config.ci
    echo projectName=\'\' >> config.ci
    echo useWorkspace=false >> config.ci
    echo scheme=\'\' >> config.ci
    echo iPhone=\'\' >> config.ci
    echo xcodeVersion=\'\' >> config.ci

    printf "We just add a very basic confi.ci file in this folder.\n"
  fi

}

function SetupGradle {

  StartAction "SetupGradle"

  # If 'next' is false, exit
  if [ ${next} == false ]; then
    ShowPreviousFailed
    return
  fi

  # Check if Gradle is installed
  GRADLE_VERSION=$(gradle -v 2>&1)
  GRADLE_INSTALLED=$(grep 'command not found' -c <<< "${GRADLE_VERSION}")

  if [ "${GRADLE_INSTALLED}" -gt 0 ]; then
    # Gradle is not installed, damn it
    ShowError
    printf "Alright, so you ${redColour}don't have gradle installed${noColour}. You need to do this when you setup the machine\n"
    next=false
    return
  else
    # Gradle is already installed

    # Check if gradle wrapper exits
    if [ -f "./gradlew" ]; then
      # GradleW exits, all good
      printf "You are ${greenColour}all good${noColour} boss\n"
    else
      # Let's setup the gradle wrapper
      gradle wrapper

      printf "Alright, we have setup ${blueColour}./gradlew${noColour}, you are all done now\n"
    fi

  fi

}

function UpgradeVersion {

  StartAction "UpgradeVersion"

  # If 'next' is false, exit
  if [ ${next} == false ]; then
    ShowPreviousFailed
    return
  fi

  LATEST_VERSION=$(curl -sS https://raw.githubusercontent.com/leftshifters/upshift/master/release 2>&1)
  LATEST_VERSION_RESULTS=$(grep 'curl:' -c <<< "${LATEST_VERSION}")

  if [ "${LATEST_VERSION_RESULTS}" -gt 0 ]; then
    printf "There was a problem in confirming the version number.\nIgnoring the message and moving on\n"
    printf "${LATEST_VERSION}\n\n"
  else
    if [ "${LATEST_VERSION}" == "${version}" ]; then
      printf "You are using the latest version of upshift v${version}\n\n"
    else
      printf "Moving you to the latest version of upshift v${version}\n"
      VERSION_UPGRADE=$(curl -fsSL https://raw.githubusercontent.com/leftshifters/upshift/master/upshift > upshift.temp && chmod +x upshift.temp && ./upshift.temp install)
      printf "${VERSION_UPGRADE}\n"
      rm upshift.temp
    fi    
  fi

}

function SendEmail {

  StartAction "SendEmail"

  # If 'next' is false, exit
  if [ ${next} == false ]; then
    ShowPreviousFailed
    return
  fi

  MAIL_TO=$1
  MAIL_SUBJECT=$2
  MAIL_BODY=$3

  SEND_EMAIL=$(echo "${MAIL_BODY}" | mail -s "${MAIL_SUBJECT}" "${MAIL_TO}")
  printf "${SEND_EMAIL}\n"

}

# Setup Jobs

jobQueue=()

function RunJobs {
  for action in "${jobQueue[@]}"
    do
      :
      "${action}"
  done
}

function FindJobQueue {

  case "${platform}" in
    "android")
      StartScript
      case "${job}" in
        "build")      jobQueue=("UpgradeVersion" "GitPull" "GitSubmodules" "SetupGradle" "AssembleAndroid");;
        "emulator")   jobQueue=("UpgradeVersion" "StartEmulator" "GitPull" "GitSubmodules" "SetupGradle" "InstallOnAndroid");;
        *)
                      ShowError
                      printf "Yo! We only support two commands for Android right now, build and emulator\n"
                      ;;
      esac
      ;;
    "ios")
      StartScript
      case "${job}" in
        "build")      jobQueue=("UpgradeVersion" "XCodeVersion" "GitPull" "GitSubmodules" "SetupPods" "XCPretty" "BuildiOS" "DeployiOSSimulator");;
        "distribute") jobQueue=("UpgradeVersion" "XCodeVersion" "GitPull" "GitSubmodules" "SetupPods" "XCPretty" "ArchiveIOS");;
        *)
                      ShowError
                      printf "Yo! We only support one commands for iOS right now: build\n"
                      ;;
      esac
      ;;
    "setup")
      case "${job}" in
        "clone")      jobQueue=("UpgradeVersion" "GitClone" "GitSubmodules" "SetupPods");;
        "config")     jobQueue=("UpgradeVersion" "AddConfig");;
        *)
                      ShowError
                      printf "Yo! We only support one commands for Setup right now: clone\n"
                      ;;
      esac
      ;;
    "install")
      jobQueue=("SetupScript");;
    "-v")
      printf "Upshift v${version}\n"
      exit
      ;;
    "action")
      case "${job}" in
        "SetupSSH")           jobQueue=("SetupSSH");;
        "InstallOnAndroid")   jobQueue=("InstallOnAndroid");;
        "GitPull")            jobQueue=("GitPull");;
        "GitClone")           jobQueue=("GitClone");;
        "StartEmulator")      jobQueue=("StartEmulator");;
        "AndroidDevices")     jobQueue=("AndroidDevices");;
        "AssembleAndroid")    jobQueue=("AssembleAndroid");;
        "GitSubmodules")      jobQueue=("GitSubmodules");;
        "SetupPods")          jobQueue=("SetupPods");;
        "XCodeVersion")       jobQueue=("XCodeVersion");;
        "XCPretty")           jobQueue=("XCPretty");;
        "BuildiOS")           jobQueue=("BuildiOS");;
        "ArchiveIOS")         jobQueue=("ArchiveIOS");;
        "DeployiOSSimulator") jobQueue=("DeployiOSSimulator");;
        "SetupScript")        jobQueue=("SetupScript");;
        "SetupGradle")        jobQueue=("SetupGradle");;
        "UpgradeVersion")     jobQueue=("UpgradeVersion");;
        *)
                      ShowError
                      printf "Yo! This command is not supported right now\n"
                      ;;
        esac
        ;;
    *)
      printf "
UPSHIFT(1)               Upshift Commands Manual               UPSHIFT(1)

${boldStyle}NAME${normalStyle}
    ${boldStyle}upshift${normalStyle} -- help you build and test mobile apps

${boldStyle}SYNOPSIS${normalStyle}
    ${boldStyle}upshift${normalStyle} ${underlineStyle}platform${noUnderlineStyle} ${underlineStyle}job${noUnderlineStyle}

${boldStyle}DESCRIPTION${normalStyle}
     The ${boldStyle}upshift${normalStyle} utility helps you clone, build and test your iOS and
     Android projects.

     The following options are currently available

     ${boldStyle}upshift${normalStyle} ${underlineStyle}android${noUnderlineStyle} ${underlineStyle}build${noUnderlineStyle}

     This command pulls the latest code, checks and install submodules
     and finally builds the Android project.

     ${boldStyle}upshift${normalStyle} ${underlineStyle}android${noUnderlineStyle} ${underlineStyle}emulator${noUnderlineStyle}

     This command starts up the emulator on the system, pulls the latest
     code, installs the git submodules and installs the app on the emulator.

     ${boldStyle}upshift${normalStyle} ${underlineStyle}ios${noUnderlineStyle} ${underlineStyle}build${noUnderlineStyle}

     This command checks if you are building on the correct Xcode version,
     pulls the code, installs submodules, installs pods, build the latest
     iOS project and deploys it on the iOS simulator

     ${boldStyle}upshift${normalStyle} ${underlineStyle}setup${noUnderlineStyle} ${underlineStyle}clone${noUnderlineStyle}

     This command helps you clone a new repository, install submodules and
     installs pods.

     ${boldStyle}upshift${normalStyle} ${underlineStyle}setup${noUnderlineStyle} ${underlineStyle}config${noUnderlineStyle}

     This command will create an empty config file in the current folder
     that you are in. Config files are setup in the main folder and are
     called config.ci . The variable defined in this file gets more
     priority than the ones defined inside upshift

     ${boldStyle}upshift${normalStyle} ${underlineStyle}install${noUnderlineStyle}

     This command installs this script on your machine

     ${boldStyle}upshift${normalStyle} ${underlineStyle}action${noUnderlineStyle} ${underlineStyle}SetupSSH${noUnderlineStyle}
     ${boldStyle}upshift${normalStyle} ${underlineStyle}action${noUnderlineStyle} ${underlineStyle}InstallOnAndroid${noUnderlineStyle}
     ${boldStyle}upshift${normalStyle} ${underlineStyle}action${noUnderlineStyle} ${underlineStyle}GitPull${noUnderlineStyle}
     ${boldStyle}upshift${normalStyle} ${underlineStyle}action${noUnderlineStyle} ${underlineStyle}GitClone${noUnderlineStyle}
     ${boldStyle}upshift${normalStyle} ${underlineStyle}action${noUnderlineStyle} ${underlineStyle}StartEmulator${noUnderlineStyle}
     ${boldStyle}upshift${normalStyle} ${underlineStyle}action${noUnderlineStyle} ${underlineStyle}AndroidDevices${noUnderlineStyle}
     ${boldStyle}upshift${normalStyle} ${underlineStyle}action${noUnderlineStyle} ${underlineStyle}AssembleAndroid${noUnderlineStyle}
     ${boldStyle}upshift${normalStyle} ${underlineStyle}action${noUnderlineStyle} ${underlineStyle}GitSubmodules${noUnderlineStyle}
     ${boldStyle}upshift${normalStyle} ${underlineStyle}action${noUnderlineStyle} ${underlineStyle}SetupPods${noUnderlineStyle}
     ${boldStyle}upshift${normalStyle} ${underlineStyle}action${noUnderlineStyle} ${underlineStyle}XCodeVersion${noUnderlineStyle}
     ${boldStyle}upshift${normalStyle} ${underlineStyle}action${noUnderlineStyle} ${underlineStyle}XCPretty${noUnderlineStyle}
     ${boldStyle}upshift${normalStyle} ${underlineStyle}action${noUnderlineStyle} ${underlineStyle}BuildiOS${noUnderlineStyle}
     ${boldStyle}upshift${normalStyle} ${underlineStyle}action${noUnderlineStyle} ${underlineStyle}DeployiOSSimulator${noUnderlineStyle}
     ${boldStyle}upshift${normalStyle} ${underlineStyle}action${noUnderlineStyle} ${underlineStyle}SetupScript${noUnderlineStyle}
     ${boldStyle}upshift${normalStyle} ${underlineStyle}action${noUnderlineStyle} ${underlineStyle}SetupGradlepSSH${noUnderlineStyle}
     ${boldStyle}upshift${normalStyle} ${underlineStyle}action${noUnderlineStyle} ${underlineStyle}UpgradeVersion${noUnderlineStyle}

     These commands will allow you to run each specific action separately

     ${boldStyle}upshift${normalStyle} ${underlineStyle}-v${noUnderlineStyle}

     This command gets you the latest version number

${boldStyle}COMPATIBILITY${normalStyle}
      This software has only been tried and tested on Mac OSX 10.11.4. If
      you're planning to use this on windows, your're on your own!

Leftshift Tech.               May 10, 2016                Leftshift Tech.
      "
      ;;
  esac
}

# Start running the scripts
FindJobQueue
RunJobs




# Ending

EndScript
