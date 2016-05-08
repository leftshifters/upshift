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

# TODO : Add the project name to the logfile and move it to a common location on the server
# TODO : Use the logs to show details on a screen somewhere

# If alwaysCleanBeforeBuild then run clean
if [ "${alwaysCleanBeforeBuild}" == 1 ]; then
  printf "Time to ${GREEN}clean up${NOCOLOR}\n\n"
  ./gradlew clean 2>&1 | tee gradle-clean-$TIMESHTAMP.log
else
  printf "Skipping :clean:, you have ${BLUE}alwaysCleanBeforeBuild${NOCOLOR} turned off\n"
fi

# Now time to build again
printf "\nTime to run ${GREEN}assembleDebug${NOCOLOR} on this thing\n\n"
./gradlew assembleDebug --stacktrace 2>&1 | tee gradle-assemble-$TIMESHTAMP.log

# Get the logged results and try to make some sense out of it
BUILD_RESULTS=$(<gradle-assemble-$TIMESHTAMP.log)

# When you build via Gradle, it seems it always sends BUILD SUCCESSFUL in the results
# This could mess up if in some build configuration, there are two messages, one is BUILD SUCCESSFUL and one otherwise
BUILD_SUCCESSFUL=$(echo $BUILD_RESULTS | grep "BUILD SUCCESSFUL" -c)

# If the build was successful, let 'em know
if [ "$BUILD_SUCCESSFUL" != "1" ]; then
  printf "\nDamn, it looks like something went ${RED}wrong${NOCOLOR}. You should check this up.\n\n"
  exit 1
else
  printf "\nAlright, the build was ${GREEN}successful${NOCOLOR} 🍺\n\n"
fi

