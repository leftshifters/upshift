##
## Add Configuration File
##
function AddConfig {

  StartAction "AddConfig"

  # If 'next' is false, exit
  if [ ${next} == false ]; then
    ShowPreviousFailed
    return
  fi

  # Check if config.ci exists in
  if [ -f "config.ci" ]; then
    printf "A config file ${redColour}already exists${noColour}\n"
  else
    # Write the config file
    echo debug=false >> config.ci
    echo alwaysCleanBeforeBuild=true >> config.ci
    echo alwaysUninstallOlderBuilds=true >> config.ci
    echo package=\'\' >> config.ci
    echo mainActivity=\'\' >> config.ci
    echo gitRepositoryURL=\'\' >> config.ci
    echo gitRepositoryBranch=\'\' >> config.ci
    echo masterPassword=\'\' >> config.ci
    echo projectName=\'\' >> config.ci
    echo useWorkspace=false >> config.ci
    echo scheme=\'\' >> config.ci
    echo iPhone=\'\' >> config.ci
    echo xcodeVersion=\'\' >> config.ci

    printf "We just add a very basic confi.ci file in this folder.\n"
  fi

}

