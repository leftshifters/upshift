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

            # TODO : Create an emulator if one doesn't exist
            # TODO : Maybe run ./gradlew connectedCheck to see if everything is working fine

            # Check if Boot Animation is still on
            # https://devmaze.wordpress.com/2011/12/12/starting-and-stopping-android-emulators/
            # adb shell getprop init.svc.bootanim
            # We don't really care about this right now

            # Check if a process which calls itself the emulator is running
            # TODO : may check this using ADB Devices
            # TODO : Gets fucked up when adb fucks up, keeps ranting multiple devices found (not the exact message)
            OUTPUT=$(ps -aef | grep emulator | grep "sdk/tools" -c)
            # If 0 processes are called emulator, it means we need to load up one
            if [ "$OUTPUT" == "0" ]; then
              EMULATOR_RESULTS=$(nohup $ANDROID_HOME/tools/emulator -avd ${emulatorName} 2>emulator.log 1>emulator.log &)
              # TODO : This is a big #HACK, only errors are returned in the first two seconds, I suck and I don't know a way out
              # TODO : Another potential problem, we redirect both 1,2 in reset mode (>), the file could get overwritten
              sleep 2
              EMULATOR_RESULTS=$(<emulator.log)
              
              # Check if there was a PANIC [to test this, put in the wrong emulator name]
              PANIC_COUNT=$(echo ${EMULATOR_RESULTS} | grep "PANIC" -c)

              # If there was a panic, throw it at the user, they messed up, not us
              if [ "${PANIC_COUNT}" -gt 0 ]; then
                printf "${BOLD}Error!${NORMAL} The emulator won't load up. Maybe the ${RED}emulatorName${NOCOLOR} key isn't correct\n"
                printf "Here's what Mr. Log says:\n\n"
                printf "${EMULATOR_RESULTS}\n\n"
                exit 1
              else 
                # Seems like there is no panic, let's check for errors
                # TODO : Can't remember how do you get an error
                ERROR_COUNT=$(echo ${EMULATOR_RESULTS} | grep "ERROR" -c)

                if [ "${ERROR_COUNT}" -gt 0 ]; then
                  printf "${EMULATOR_RESULTS}\n\n"
                  exit 1
                else
                  printf "All set baby, ${GREEN}starting${NOCOLOR} to load up the emulator\n"
                  # ADB gives this option to wait for the device till it comes up, let's just depend on it, 
                  # this is really mess with us when there is a problem with the emulator fails to load because of it's own bugs
                  adb wait-for-device

                  # Now that the device is available, the worst is over
                  # Check if the emulator has finished botting, if not, sleedp for 2 seconds and try this again, this is our final trigger
                  printf "Seems like the device is now ${GREEN}available${NOCOLOR}, we are getting close\n"
                  SCREEN_LOADING=$(adb shell getprop sys.boot_completed | tr -d '\r')
                  while [ "$SCREEN_LOADING" != "1" ]; do
                    sleep 4
                    printf "Check if the emulator has finished booting, why is this thing so ${BLUE}damn${NOCOLOR} slow?\n"
                    SCREEN_LOADING=$(adb shell getprop sys.boot_completed | tr -d '\r')
                  done

                  # Alright, everything is now done. This is just used to unlock the device.
                  printf "Almost ${GREEN}done${NOCOLOR}, touch the device once\n"
                  adb shell input keyevent 82
                  printf "${GREEN}Super!${NOCOLOR} The emulator is now running. You're one lucky person 🍺\n"
                fi
              fi
            else
              printf "${GREEN}Dude${NOCOLOR}, Looks like the emulator is already running!\n\n"
              exit 0
            fi

