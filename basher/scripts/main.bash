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
	LOGFILE=$3
	mkdir -p .upshift/logs/
	git pull $1 $2 2>&1 | tee $3
}

GitSubmoduleInit() {
	LOGFILE=$1
	mkdir -p .upshift/logs/
	git submodule init 2>&1 | tee $1
}

GitSubmoduleUpdate() {
	LOGFILE=$1
	mkdir -p .upshift/logs/
	git submodule update 2>&1 | tee $1
}