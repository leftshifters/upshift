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

