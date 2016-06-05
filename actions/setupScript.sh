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

