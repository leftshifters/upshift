#!/bin/bash

# 1. Read the configuration files
debug=false
alwaysCleanBeforeBuild=true
alwaysUninstallOlderBuilds=true
package=""
mainActivity=""
gitRepositoryBranch=""

# 2. Load up config from config file
if [ -f "./config.cfg" ]; then
  source ./config.cfg
fi

# 3. Dump commands to the screen, only if one wants to debug
if [ "${debug}" == true ];then
  set -v
fi

# 4. Make sure things look good. Here are some font and color adjustments
redColour='\033[0;31m'
greenColour='\033[0;32m'
blueColour='\033[0;34m'
noColour='\033[0m'

boldStyle=$(tput bold)
normalStyle=$(tput sgr0)

# 5. Exit script on error
# set -e
# (Maye not)

# 6. Setup Global Variables
next=true





# Setup Internal Functions

function StartScript {
  printf "\n${blueColour}Loading up...${noColour}\n"
}

function EndScript {
  printf "\n"
}

function ShowError {
  printf "\n${redColour}Something went wrong!${noColour}\n"
}

StartScript

# Setup Actions

##
## setup-ssh
## 

function SetupSSH {

  # Details about the script came from here
  # https://help.github.com/articles/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent/

  # Check if email has been defined by the user
  if [ "${emailForSSHKey}" != "" ]; then
    # TODO : Allow the user to add the keys to a non default place
    # Check if an id_rsa already exists at the defualt location
    if [ ! -f ~/.ssh/id_rsa ]; then
      printf "File does not exist at ~/.ssh/id_rsa"
      echo -ne '\n' | ssh-keygen -t rsa -b 4096 -C "${emailForSSHKey}"

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
    printf "Dude, you need to add the <${redColour}emailForSSHKey${noColour}> parameter to get this to work\n"
  fi
}

##
## install-on-android
##

function InstallOnAndroid {

  # Make a TIMESHTAMP for the log file
  TIMESHTAMP=$(date +%Y%m%d%H%M%S)

  # TODO : Add the project name to the logfile and move it to a common location on the server
  # TODO : Use the logs to show details on a screen somewhere

  # If alwaysCleanBeforeBuild then run clean
  if [ "${alwaysCleanBeforeBuild}" == true ]; then
    printf "Time to ${greenColour}clean up${noColour}\n\n"
    ./gradlew clean 2>&1 | tee gradle-clean-$TIMESHTAMP.log
  else
    printf "Skipping :clean:, you have ${blueColour}alwaysCleanBeforeBuild${noColour} turned off\n"
  fi

  # Uninstall older builds if the setting so desires
  if [ "${alwaysUninstallOlderBuilds}" == true ]; then
    printf "Time to ${greenColour}uninstall${noColour} older builds\n\n"
    ./gradlew uninstallAll 2>&1 | tee gradle-uninstall-$TIMESHTAMP.log
  else
    printf "Skipping :uninstallAll:, you have ${blueColour}alwaysUninstallOlderBuilds${noColour} turned off\n"
  fi

  # Now time to build again
  printf "\nTime to run ${greenColour}installDebug${noColour} on this thing\n\n"
  ./gradlew installDebug --stacktrace 2>&1 | tee gradle-install-$TIMESHTAMP.log

  # Get the logged results and try to make some sense out of it
  BUILD_RESULTS=$(<gradle-install-$TIMESHTAMP.log)

  # When you build via Gradle, it seems it always sends BUILD SUCCESSFUL in the results
  # This could mess up if in some build configuration, there are two messages, one is BUILD SUCCESSFUL and one otherwise
  BUILD_SUCCESSFUL=$(echo $BUILD_RESULTS | grep "BUILD SUCCESSFUL" -c)

  # If the build was successful, let 'em know
  if [ "$BUILD_SUCCESSFUL" != "1" ]; then
    ShowError
    printf "Damn, it looks like something went ${redColour}wrong${noColour}. You should check this up.\n\n"
    next=false
  else
    printf "\n\n${greenColour}Super${noColour}! The build was fine.\n"
    # TODO : Someday figure out how to get package and mainActivity automatically
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

  # Make a TIMESHTAMP for log file
  TIMESHTAMP=$(date +%Y%m%d%H%M%S)

  # Check if the branch name is defined
  # TODO : Automatically pick the current branch
  if [ "${gitRepositoryBranch}" != "" ]; then

    printf "Alright, let's ${greenColour}pull${noColour} the ${gitRepositoryBranch} branch for this repo\n\n"

    # Alright, let's pull
    git pull origin ${gitRepositoryBranch} 2>&1 | tee git-pull-$TIMESHTAMP.log

    # Load up the results to see if there were any errors
    PULL_RESULTS=$(<git-pull-$TIMESHTAMP.log)
    # If there was a fatal error, tell the user there's something wrong
    if [ "$(printf "${PULL_RESULTS}" | grep "fatal" -c)" -gt "0" ]; then
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
    printf "Dude, you need to add the ${blueColour}gitRepositoryBranch${noColour} for this to work\n\n"
    next=false
  fi
}

##
## git-clone
##

function GitClone {

  # Make a TIMESHTAMP for log file
  TIMESHTAMP=$(date +%Y%m%d%H%M%S)

  # Check if the repo URL is defined
  if [ "${gitRepositoryURL}" != "" ]; then
    # Check if the branch name is defined
    # TODO : Automatically pick the current branch
    if [ "${gitRepositoryBranch}" != "" ]; then

      printf "Alright, let's ${greenColour}clone${noColour} the ${blueColour}${gitRepositoryBranch}${noColour} branch for the ${blueColour}${gitRepositoryURL}${noColour} repo\n\n"
    
      # Alright, let's pull
      # But you can't pull into an empty directly, now since you already have bitrise.yml and .bitrise.secrets.yml in your directory,
      #   you will need to clone into another folder, and move stuff back here
      #   we can't do what the rest of the world tries to do - which is git init, add remote,
      #   because we want to ensure we do depth=1 and not download the whole repo, which can be painful at times
      git clone -b ${gitRepositoryBranch} ${gitRepositoryURL} tmp --depth 1  2>&1 | tee git-clone-$TIMESHTAMP.log
      mv tmp/* . 2>&1 | tee git-clone-$TIMESHTAMP.log
      mv tmp/.* . 2>/dev/null | tee git-clone-$TIMESHTAMP.log
      rm -rf tmp/ 2>&1 | tee git-clone-$TIMESHTAMP.log

      # Load up the results to see if there were any errors
      CLONE_RESULTS=$(<git-clone-$TIMESHTAMP.log)
      # If there was a fatal error, tell the user there's something wrong
      if [ "$(printf "${CLONE_RESULTS}" | grep "fatal" -c)" -gt "0" ] || [ "$(printf "${CLONE_RESULTS}" | grep "error" -c)" -gt "0" ]; then
        ShowError
        printf "Something failed while we were cloning, you should look this up\n\n"
        next=false
      else
        # All done
        printf "\nAll done ${greenColour}baby${noColour}! 🍺.\n\n"
      fi

    else
      # The user hasn't added the required keys
      ShowError
      printf "Dude, you need to add the ${blueColour}gitRepositoryBranch${noColour} for this to work\n\n"
      exit 1
    fi
  else
    # The user hasn't added the required keys
    ShowError
    printf "Dude, you need to add the ${blueColour}gitRepositoryURL${noColour} for this to work\n\n"
    next=false
  fi
}

##
## start-emulator
##

function StartEmulator {

  # TODO : Create an emulator if one doesn't exist
  # TODO : Maybe run ./gradlew connectedCheck to see if everything is working fine

  # Check if Boot Animation is still on
  # https://devmaze.wordpress.com/2011/12/12/starting-and-stopping-android-emulators/
  # adb shell getprop init.svc.bootanim
  # We don't really care about this right now

  # Check if a process which calls itself the emulator is running
  # TODO : may check this using ADB Devices
  # TODO : Gets fucked up when adb fucks up, keeps ranting multiple devices found (not the exact message)
  OUTPUT=$(ps -aef | grep emulator | grep "sdk/tools" -c)
  # If 0 processes are called emulator, it means we need to load up one
  if [ "$OUTPUT" == "0" ]; then

    if [ "${emulatorName}" != "" ]; then

      EMULATOR_RESULTS=$(nohup $ANDROID_HOME/tools/emulator -avd ${emulatorName} 2>emulator.log 1>emulator.log &)
      # TODO : This is a big #HACK, only errors are returned in the first two seconds, I suck and I don't know a way out
      # TODO : Another potential problem, we redirect both 1,2 in reset mode (>), the file could get overwritten
      sleep 2
      EMULATOR_RESULTS=$(<emulator.log)
      
      # Check if there was a PANIC [to test this, put in the wrong emulator name]
      PANIC_COUNT=$(echo ${EMULATOR_RESULTS} | grep "PANIC" -c)

      # If there was a panic, throw it at the user, they messed up, not us
      if [ "${PANIC_COUNT}" -gt 0 ]; then
        ShowError
        printf "The emulator won't load up. Maybe the ${redColour}emulatorName${noColour} key isn't correct\n"
        printf "Here's what Mr. Log says:\n\n"
        printf "${EMULATOR_RESULTS}\n\n"
        next=false
      else 
        # Seems like there is no panic, let's check for errors
        # TODO : Can't remember how do you get an error
        ERROR_COUNT=$(echo ${EMULATOR_RESULTS} | grep "ERROR" -c)

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
      printf "Dude, you need to define ${blueColour}emulatorName${noColour} for this to work\n\n"
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
  # Make a TIMESHTAMP for log file
  TIMESHTAMP=$(date +%Y%m%d%H%M%S)

  # TODO : Add the project name to the logfile and move it to a common location on the server
  # TODO : Use the logs to show details on a screen somewhere

  # If alwaysCleanBeforeBuild then run clean
  if [ "${alwaysCleanBeforeBuild}" == 1 ]; then
    printf "Time to ${greenColour}clean up${noColour}\n\n"
    ./gradlew clean 2>&1 | tee gradle-clean-$TIMESHTAMP.log
  else
    printf "Skipping :clean:, you have ${blueColour}alwaysCleanBeforeBuild${noColour} turned off\n"
  fi

  # Now time to build again
  printf "\nTime to run ${greenColour}assembleDebug${noColour} on this thing\n\n"
  ./gradlew assembleDebug --stacktrace 2>&1 | tee gradle-assemble-$TIMESHTAMP.log

  # Get the logged results and try to make some sense out of it
  BUILD_RESULTS=$(<gradle-assemble-$TIMESHTAMP.log)

  # When you build via Gradle, it seems it always sends BUILD SUCCESSFUL in the results
  # This could mess up if in some build configuration, there are two messages, one is BUILD SUCCESSFUL and one otherwise
  BUILD_SUCCESSFUL=$(echo $BUILD_RESULTS | grep "BUILD SUCCESSFUL" -c)

  # If the build was successful, let 'em know
  if [ "$BUILD_SUCCESSFUL" != "1" ]; then
    ShowError
    printf "Damn, the build was not ${redColour}successful${noColour}. You should check this up.\n\n"
    next=false
  else
    printf "\nAlright, the build was ${greenColour}successful${noColour} 🍺\n\n"
  fi

}

#SetupSSH
#InstallOnAndroid
#GitPull
#GitClone
#StartEmulator
#AndroidDevices
#AssembleAndroid


EndScript