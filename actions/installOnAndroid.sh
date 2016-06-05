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
