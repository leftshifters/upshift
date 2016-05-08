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
            if [ "${alwaysCleanBeforeBuild}" == "1" ]; then
              printf "Time to ${GREEN}clean up${NOCOLOR}\n\n"
              ./gradlew clean 2>&1 | tee gradle-clean-$TIMESHTAMP.log
            else
              printf "Skipping :clean:, you have ${BLUE}alwaysCleanBeforeBuild${NOCOLOR} turned off\n"
            fi

            # Uninstall older builds if the setting so desires
            if [ "${alwaysUninstallOlderBuilds}" == "1" ]; then
              printf "Time to ${GREEN}uninstall${NOCOLOR} older builds\n\n"
              ./gradlew uninstallAll 2>&1 | tee gradle-uninstall-$TIMESHTAMP.log
            else
              printf "Skipping :uninstallAll:, you have ${BLUE}alwaysUninstallOlderBuilds${NOCOLOR} turned off\n"
            fi

            # Now time to build again
            printf "\nTime to run ${GREEN}installDebug${NOCOLOR} on this thing\n\n"
            ./gradlew installDebug --stacktrace 2>&1 | tee gradle-install-$TIMESHTAMP.log

            # Get the logged results and try to make some sense out of it
            BUILD_RESULTS=$(<gradle-install-$TIMESHTAMP.log)

            # When you build via Gradle, it seems it always sends BUILD SUCCESSFUL in the results
            # This could mess up if in some build configuration, there are two messages, one is BUILD SUCCESSFUL and one otherwise
            BUILD_SUCCESSFUL=$(echo $BUILD_RESULTS | grep "BUILD SUCCESSFUL" -c)

            printf "\n\n${GREEN}Super${NOCOLOR}! The build was fine.\n"

            # If the build was successful, let 'em know
            if [ "$BUILD_SUCCESSFUL" != "1" ]; then
              printf "Damn, it looks like something went ${RED}wrong${NOCOLOR}. You should check this up.\n\n"
              exit 1
            else
              # TODO : Someday figure out how to get package and mainActivity automatically
              # Check if package is empty
              if [ "${package}" != "" ];then
                if [ "${mainActivity}" != "" ]; then
                  printf "Starting activity ${BLUE}${mainActivity}${NOCOLOR} in package ${BLUE}${package}${NOCOLOR}\n"

                  # Start the activity and package
                  adb shell am start -n ${package}/${package}.${mainActivity}

                  # Tell the user everything is nice and easy
                  printf "\nAlright, the build was ${GREEN}successful${NOCOLOR} 🍺\n\n"
                else
                  # The mainActivity is empty it seems
                  printf "Alright, the build was ${GREEN}successful${NOCOLOR}, but there was no ${BLUE}mainActivity${NOCOLOR} defined, so couldn't start it automatically 🍺\n\n"
                fi
              else
                # The package is empty it seems
                printf "Alright, the build was ${GREEN}successful${NOCOLOR}, but there was no ${BLUE}package${NOCOLOR} defined, so couldn't start it automatically 🍺\n\n"
              fi
            fi
