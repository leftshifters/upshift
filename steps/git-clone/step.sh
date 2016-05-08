#!/bin/bash

# Turning this off, because this ends the script if 0 get returned anywhere
# set -e

# This is for font color and style adjustment
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NOCOLOR='\033[0m'
BOLD=$(tput bold)
NORMAL=$(tput sgr0)            

# An empty space to separate from the preceeding block
if [ "${debug}" == "1" ];then
  set -v
fi

# An empty space to separate from the preceeding block
printf "\n"

# Make a TIMESHTAMP for log file
TIMESHTAMP=$(date +%Y%m%d%H%M%S)

# Check if the repo URL is defined
if [ "${gitRepositoryURL}" != "" ]; then
  # Check if the branch name is defined
  # TODO : Automatically pick the current branch
  if [ "${gitRepositoryBranch}" != "" ]; then

    printf "Alright, let's ${GREEN}clone${NOCOLOR} the ${BLUE}${gitRepositoryBranch}${NOCOLOR} branch for the ${BLUE}${gitRepositoryURL}${NOCOLOR} repo\n\n"
  
    # Alright, let's pull
    # But you can't pull into an empty directly, now since you already have bitrise.yml and .bitrise.secrets.yml in your directory,
    #   you will need to clone into another folder, and move stuff back here
    #   we can't do what the rest of the world tries to do - which is git init, add remote,
    #   because we want to ensure we do depth=1 and not download the whole repo, which can be painful at times
    git clone -b ${gitRepositoryBranch} ${gitRepositoryURL} tmp --depth 1  2>&1 | tee git-clone-$TIMESHTAMP.log
    mv tmp/* . 2>&1 | tee git-clone-$TIMESHTAMP.log
    mv tmp/.* . 2>/dev/null | tee git-clone-$TIMESHTAMP.log
    rm -rf tmp/ 2>&1 | tee git-clone-$TIMESHTAMP.log

    # Load up the results to see if there were any errors
    CLONE_RESULTS=$(<git-clone-$TIMESHTAMP.log)
    # If there was a fatal error, tell the user there's something wrong
    if [ "$(printf "${CLONE_RESULTS}" | grep "fatal" -c)" -gt "0" ] || [ "$(printf "${CLONE_RESULTS}" | grep "error" -c)" -gt "0" ]; then
      printf "\n${BOLD}Error!${NORMAL} Something went wrong with the clone, you should look this up\n\n"
      exit 1
    else
      # All done
      printf "\nAll done ${GREEN}baby${NOCOLOR}! 🍺.\n\n"
      exit 0
    fi

  else
    # The user hasn't added the required keys
    printf "${BOLD}Error!${NORMAL} Dude, you need to add the ${BLUE}gitRepositoryBranch${NOCOLOR} for this to work\n\n"
    exit 1
  fi
else
  # The user hasn't added the required keys
  printf "${BOLD}Error!${NORMAL} Dude, you need to add the ${BLUE}gitRepositoryURL${NOCOLOR} for this to work\n\n"
  exit 1
fi

