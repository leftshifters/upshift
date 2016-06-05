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

