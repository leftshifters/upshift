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

