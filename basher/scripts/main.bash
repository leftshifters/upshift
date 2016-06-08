UpgradeScript() {
	curl -fsSL https://raw.githubusercontent.com/leftshifters/upshift/master/upshift > upshift.temp && chmod +x upshift.temp && ./upshift.temp install && rm upshift.temp
}

SetupXcpretty() {
	CI=$1
	RootPassword=$2

	if [ "${CI}" == "true" ]; then
		# CI is true, we now need password
		if [ "${RootPassword}" != "" ]; then
			echo -e "${RootPassword}" | sudo -S gem install xcpretty
			exit 0
		else
			exit 1
		fi
	else
		# User should type in the password
		sudo -S gem install xcpretty
		exit 0
	fi
}

GitPull() {
	REMOTE=$1
	BRANCH=$2
	TIMESTAMP=$3
	mkdir -p .upshift/logs/
	git pull $1 $2 2>&1 | tee $3

	# PULL_RESULTS=$(<".upshift/logs/git-pull-$3.log")
	# PULL_RESULTS_FATAL=$(grep "fatal:" -c <<< "${PULL_RESULTS}")
	# PULL_RESULTS_ERROR=$(grep "error:" -c <<< "${PULL_RESULTS}")
	# # If there was a fatal error, tell the user there's something wrong
	# if [ "${PULL_RESULTS_FATAL}" -gt "0" ] || [ "${PULL_RESULTS_ERROR}" -gt "0" ]; then
	# 	exit 1
	# else
	# 	# All done
	# 	exit 0
	# fi
}