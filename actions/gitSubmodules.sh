##
## Setup Submodule if they exist
##

function GitSubmodules {

  StartAction "GitSubmodules"

  # If 'next' is false, exit
  if [ ${next} == false ]; then
    ShowPreviousFailed
    return
  fi

  # Make a TIMESHTAMP for log file
  TIMESHTAMP=$(date +%Y%m%d%H%M%S)

  # Check if .gitmodules exists
  if [ -f ".gitmodules" ]; then
    # If the file exists, we need to run init and update and catch errors
    git submodule init 2>&1 | tee "git-submodule-init-${TIMESHTAMP}.log"
    git submodule update 2>&1 | tee "git-submodule-update-${TIMESHTAMP}.log"

    SUBMODULE_RESULTS=$(<"git-submodule-update-${TIMESHTAMP}.log")

    if [ $(echo "${SUBMODULE_RESULTS}" | grep "fatal:" -c) -gt 0 ] || [ $(echo "${SUBMODULE_RESULTS}" | grep "error:" -c) -gt 0 ]; then
      ShowError
      printf "Damn, initializing submodules was ${redColour}not successful${noColour}. You should check this up.\n\n"
      next=false
    else
      printf "\nSubmodules are now ${greenColour}setup${noColour}, one less thing to think about! 🍺\n\n"
    fi
    # Else Quietly ignore
  else
    printf "\nIt looks like this project doesn't use ${greenColour}submodules${noColour}.\n\n"
  fi
}

