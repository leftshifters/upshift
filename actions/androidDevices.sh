##
## android-devices
##
function AndroidDevices {

  StartAction "AndroidDevices"

  # If 'next' is false, exit
  if [ ${next} == false ]; then
    ShowPreviousFailed
    return
  fi

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
    ShowError
    printf "No physical devices were found\n\n"
    next=false
  fi
}

