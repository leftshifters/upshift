##
## start-emulator
##

function StartEmulator {

  StartAction "StartEmulator"

  # If 'next' is false, exit
  if [ ${next} == false ]; then
    ShowPreviousFailed
    return
  fi

  # Check if Boot Animation is still on
  # https://devmaze.wordpress.com/2011/12/12/starting-and-stopping-android-emulators/
  # adb shell getprop init.svc.bootanim
  # We don't really care about this right now

  # Check if a process which calls itself the emulator is running
  OUTPUT=$(ps -aef | grep emulator | grep "sdk/tools" -c)
  # If 0 processes are called emulator, it means we need to load up one
  if [ "$OUTPUT" == "0" ]; then

    if [ "${emulatorName}" != "" ]; then

      EMULATOR_RESULTS=$(nohup "$ANDROID_HOME/tools/emulator" -avd "${emulatorName}" 2>emulator.log 1>emulator.log &)
      # TODO : This is a big #HACK, only errors are returned in the first two seconds, I suck and I don't know a way out
      # TODO : Another potential problem, we redirect both 1,2 in reset mode (>), the file could get overwritten
      sleep 2
      EMULATOR_RESULTS=$(<emulator.log)

      # Check if there was a PANIC [to test this, put in the wrong emulator name]
      PANIC_COUNT=$(echo "${EMULATOR_RESULTS}" | grep "PANIC" -c)

      # If there was a panic, throw it at the user, they messed up, not us
      if [ "${PANIC_COUNT}" -gt 0 ]; then
        ShowError
        printf "The emulator won't load up. Maybe the ${redColour}emulatorName${noColour} key isn't correct\n"
        printf "Here's what Mr. Log says:\n\n"
        printf "${EMULATOR_RESULTS}\n\n"
        next=false
      else
        # Seems like there is no panic, let's check for errors
        ERROR_COUNT=$(echo "${EMULATOR_RESULTS}" | grep "ERROR" -c)

        if [ "${ERROR_COUNT}" -gt 0 ]; then
          printf "${EMULATOR_RESULTS}\n\n"
          next=false
        else
          printf "All set baby, ${greenColour}starting${noColour} to load up the emulator\n"
          # ADB gives this option to wait for the device till it comes up, let's just depend on it,
          # this is really mess with us when there is a problem with the emulator fails to load because of it's own bugs
          adb wait-for-device

          # Now that the device is available, the worst is over
          # Check if the emulator has finished botting, if not, sleedp for 2 seconds and try this again, this is our final trigger
          printf "Seems like the device is now ${greenColour}available${noColour}, we are getting close\n"
          SCREEN_LOADING=$(adb shell getprop sys.boot_completed | tr -d '\r')
          while [ "$SCREEN_LOADING" != "1" ]; do
            sleep 4
            printf "Check if the emulator has finished booting, why is this thing so ${blueColour}damn${noColour} slow?\n"
            SCREEN_LOADING=$(adb shell getprop sys.boot_completed | tr -d '\r')
          done

          # Alright, everything is now done. This is just used to unlock the device.
          printf "Almost ${greenColour}done${noColour}, touch the device once\n"
          adb shell input keyevent 82
          printf "${greenColour}Super!${noColour} The emulator is now running. You're one lucky person 🍺\n"
        fi
      fi

    else
      ShowError
      printf "Dude, you need to define ${blueColour}emulatorName${noColour} for this to work\nYou can get a sample config by running upshift setup config\n\n"
      next=false
    fi
  else
    printf "${greenColour}Dude${noColour}, Looks like the emulator is already running!\n\n"
  fi
}

