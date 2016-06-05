##
## setup-ssh
##

function SetupSSH {

  StartAction "SetupSSH"

  # Details about the script came from here
  # https://help.github.com/articles/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent/

  # If 'next' is false, exit
  if [ ${next} == false ]; then
    ShowPreviousFailed
    return
  fi

  # Check if email has been defined by the user
  if [ "${emailForSSHKey}" != "" ]; then
    # Check if an id_rsa already exists at the defualt location
    if [ ! -f ~/.ssh/id_rsa ]; then
      printf "File does not exist at ~/.ssh/id_rsa"
      echo -e '\n' | ssh-keygen -t rsa -b 4096 -C "${emailForSSHKey}"

      # Show the created keys on the screen
      ID_RSA=$(<~/.ssh/id_rsa)
      ID_RSA_PUB=$(<~/.ssh/id_rsa.pub)

      printf "${boldStyle}id_rsa${normalStyle}\n"
      printf "${ID_RSA}"
      printf "\n\n${boldStyle}id_rsa.pub${normalStyle}\n"
      printf "${ID_RSA_PUB}"

      printf "All done 🍺\n"

    else
      ShowError
      printf "Can't do this, looks like an id_rsa already exists at ~/.ssh/id_rsa, get rid of that first\n"
      next=false
    fi
  else
    ShowError
    printf "Dude, you need to add the <${redColour}emailForSSHKey${noColour}> parameter to get this to work\nYou can get a sample config by running upshift setup config\n"
    next=false
  fi
}
