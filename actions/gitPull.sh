##
## git-pull
##

function GitPull {

  StartAction "GitPull"

  # If this is coming via CI, don't pull
  if [ "${CI}" == true ]; then
    ShowCIFailed
    return
  fi

  # If 'next' is false, exit
  if [ ${next} == false ]; then
    ShowPreviousFailed
    return
  fi

  # Make a TIMESHTAMP for log file
  TIMESHTAMP=$(date +%Y%m%d%H%M%S)

  # Check if the branch name is defined
  if [ "${gitRepositoryBranch}" != "" ]; then

    printf "Alright, let's ${greenColour}pull${noColour} the ${gitRepositoryBranch} branch for this repo\n\n"

    # Alright, let's pull
    git pull origin ${gitRepositoryBranch} 2>&1 | tee "git-pull-${TIMESHTAMP}.log"

    # Load up the results to see if there were any errors
    PULL_RESULTS=$(<"git-pull-${TIMESHTAMP}.log")
    PULL_RESULTS_FATAL=$(grep "fatal:" -c <<< "${PULL_RESULTS}")
    PULL_RESULTS_ERROR=$(grep "error:" -c <<< "${PULL_RESULTS}")
    # If there was a fatal error, tell the user there's something wrong
    if [ "${PULL_RESULTS_FATAL}" -gt "0" ] || [ "${PULL_RESULTS_ERROR}" -gt "0" ]; then
      ShowError
      printf "Something went wrong with the pull, you should look this up\n\n"
      next=false
    else
      # All done
      printf "\nAll done ${greenColour}baby${noColour}! 🍺.\n\n"
    fi

  else
    # The user hasn't added the required keys
    ShowError
    printf "Dude, you need to add the ${blueColour}gitRepositoryBranch${noColour} for this to work\nYou can get a sample config by running upshift setup config\n\n"
    next=false
  fi
}

