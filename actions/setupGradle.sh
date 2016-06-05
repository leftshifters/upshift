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

