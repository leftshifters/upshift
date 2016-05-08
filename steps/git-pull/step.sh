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

            # Check if the branch name is defined
            # TODO : Automatically pick the current branch
            if [ "${gitRepositoryBranch}" != "" ]; then

              printf "Alright, let's ${GREEN}pull${NOCOLOR} the ${gitRepositoryBranch} branch for this repo\n\n"
            
              # Alright, let's pull
              git pull origin ${gitRepositoryBranch} 2>&1 | tee git-pull-$TIMESHTAMP.log

              # Load up the results to see if there were any errors
              PULL_RESULTS=$(<git-pull-$TIMESHTAMP.log)
              # If there was a fatal error, tell the user there's something wrong
              if [ "$(printf "${PULL_RESULTS}" | grep "fatal" -c)" -gt "0" ]; then
                printf "\n${BOLD}Error!${NORMAL} Something went wrong with the pull, you should look this up\n\n"
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

