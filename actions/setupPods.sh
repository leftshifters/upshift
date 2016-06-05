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

