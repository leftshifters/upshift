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
      PROBABLE_PATH=$(find ./build/Build/Products/Debug-iphonesimulator | grep "/${fullProductName}$")

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

