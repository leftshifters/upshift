##
## git-clone
##

function GitClone {

  StartAction "GitClone"

  # If 'next' is false, exit
  if [ ${next} == false ]; then
    ShowPreviousFailed
    return
  fi

  # Make a TIMESHTAMP for log file
  TIMESHTAMP=$(date +%Y%m%d%H%M%S)

  # Check if the repo URL is defined
  if [ "${gitRepositoryURL}" != "" ]; then
    # Check if the branch name is defined
    if [ "${gitRepositoryBranch}" != "" ]; then

      # Check if this is already a git repository

      if [ -f ".git/config" ]; then

        printf "You have a .git/config, I believe you already have a cloned repo here!\n"

      else

        printf "Alright, let's ${greenColour}clone${noColour} the ${blueColour}${gitRepositoryBranch}${noColour} branch for the ${blueColour}${gitRepositoryURL}${noColour} repo\n\n"

        # Alright, let's pull
        # But you can't pull into an empty directly, now since you already have bitrise.yml and .bitrise.secrets.yml in your directory,
        #   you will need to clone into another folder, and move stuff back here
        #   we can't do what the rest of the world tries to do - which is git init, add remote,
        #   because we want to ensure we do depth=1 and not download the whole repo, which can be painful at times
        git clone -b "${gitRepositoryBranch}" "${gitRepositoryURL}" tmp --progress --depth 1  2>&1 | tee "git-clone-${TIMESHTAMP}.log"
        mv tmp/* . 2>&1 | tee "git-clone-${TIMESHTAMP}.log"
        mv tmp/.* . 2>/dev/null | tee "git-clone-${TIMESHTAMP}.log"
        rm -rf tmp/ 2>&1 | tee "git-clone-${TIMESHTAMP}.log"

        # Load up the results to see if there were any errors
        CLONE_RESULTS=$(<"git-clone-${TIMESHTAMP}.log")
        # If there was a fatal error, tell the user there's something wrong
        if [ "$(printf "${CLONE_RESULTS}" | grep "fatal" -c)" -gt "0" ] || [ "$(printf "${CLONE_RESULTS}" | grep "error" -c)" -gt "0" ]; then
          ShowError
          printf "Something failed while we were cloning, you should look this up\n\n"
          next=false
        else
          # All done
          printf "\nAll done ${greenColour}baby${noColour}! 🍺.\n\n"
        fi

      fi

    else
      # The user hasn't added the required keys
      ShowError
      printf "Dude, you need to add the ${blueColour}gitRepositoryBranch${noColour} for this to work\nYou can get a sample config by running upshift setup config\n\n"
      next=false
    fi
  else
    # The user hasn't added the required keys
    ShowError
    printf "Dude, you need to add the ${blueColour}gitRepositoryURL${noColour} for this to work\nYou can get a sample config by running upshift setup config\n\n"
    next=false
  fi
}

