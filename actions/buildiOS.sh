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

