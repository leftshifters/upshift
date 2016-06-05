class UpgradeScript {
	
}

// function UpgradeVersion {

//   StartAction "UpgradeVersion"

//   # If 'next' is false, exit
//   if [ ${next} == false ]; then
//     ShowPreviousFailed
//     return
//   fi

//   LATEST_VERSION=$(curl -sS https://raw.githubusercontent.com/leftshifters/upshift/master/release 2>&1)
//   LATEST_VERSION_RESULTS=$(grep 'curl:' -c <<< "${LATEST_VERSION}")

//   if [ "${LATEST_VERSION_RESULTS}" -gt 0 ]; then
//     printf "There was a problem in confirming the version number.\nIgnoring the message and moving on\n"
//     printf "${LATEST_VERSION}\n\n"
//   else
//     if [ "${LATEST_VERSION}" == "${version}" ]; then
//       printf "You are using the latest version of upshift v${version}\n\n"
//     else
//       printf "Moving you to the latest version of upshift v${version}\n"
//       VERSION_UPGRADE=$(curl -fsSL https://raw.githubusercontent.com/leftshifters/upshift/master/upshift > upshift.temp && chmod +x upshift.temp && ./upshift.temp install)
//       printf "${VERSION_UPGRADE}\n"
//       rm upshift.temp
//     fi    
//   fi

// }
