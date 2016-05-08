#!/bin/bash

# Turning this off, because this ends the script if 0 devices are found
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

# Touch device to that they can get unlocked, otherwise ADB just ignores them
# Just make sure you hide the STDERR because well we don't care too much about it
TOUCH_DEVICE=$(adb -d shell input keyevent 82 2>&1)

# Find out how many devices are available, and are not emulators
DEVICES_FOUND=$(adb devices | grep -v "List of devices attached" | grep -v emulator -c)
# For some reason we get one added to the value, so let's just substract it
let DEVICES_FOUND=DEVICES_FOUND-1;

# Show appropriate error message
if [ "${DEVICES_FOUND}" -gt 0 ]; then
  printf "${DEVICES_FOUND} physical device(s) were found 🍺\n\n"
else
  printf "${BOLD}Error!${NORMAL} No physical devices were found\n\n"
  exit 1
fi

