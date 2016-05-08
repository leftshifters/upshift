            #!/bin/bash

            # Details about the script came from here
            # https://help.github.com/articles/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent/

            # This is for font color and style adjustment
            RED='\033[0;31m'
            GREEN='\033[0;32m'
            BLUE='\033[0;34m'
            NOCOLOR='\033[0m'
            BOLD=$(tput bold)
            NORMAL=$(tput sgr0)

            set -e

            # Dump everything only if debug is on
            if [ "${debug}" == "1" ];then
              set -v
            fi

            # An empty space to separate from the preceeding block
            printf "\n"

            # Check if email has been defined by the user
            if [ "${emailForSSHKey}" != "" ]; then
              # TODO : Allow the user to add the keys to a non default place
              # Check if an id_rsa already exists at the defualt location
              if [ ! -f ~/.ssh/id_rsa ]; then
                printf "File does not exist at ~/.ssh/id_rsa"
                echo -ne '\n' | ssh-keygen -t rsa -b 4096 -C "${emailForSSHKey}"

                # Show the created keys on the screen
                ID_RSA=$(<~/.ssh/id_rsa)
                ID_RSA_PUB=$(<~/.ssh/id_rsa.pub)

                printf "${BOLD}id_rsa${NORMAL}\n"
                printf "${ID_RSA}"
                printf "\n\n${BOLD}id_rsa.pub${NORMAL}\n"
                printf "${ID_RSA_PUB}"

                printf "All done 🍺\n\n"

              else
                printf "${BOLD}Error!${NORMAL} Looks like an id_rsa already exists at ~/.ssh/id_rsa \n\n"
                exit 1
              fi
            else
              printf "${BOLD}Error!${NORMAL} Parameter ${RED}emailForSSHKey${NOCOLOR} not found \n\n"
              exit 1
            fi

