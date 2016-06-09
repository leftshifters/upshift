UpgradeScript() {
	curl -fsSL https://raw.githubusercontent.com/leftshifters/upshift/master/upshift > upshift.temp && chmod +x upshift.temp && ./upshift.temp install && rm upshift.temp
}

SetupGem() {
	GEM=$1
	CI=$2
	RootPassword=$3

	if [ "$2" == "true" ]; then
		# CI is true, we now need password
		if [ "$3" != "" ]; then
			echo -e "$3" | sudo -S gem install $1
			exit 0
		else
			exit 1
		fi
	else
		# User should type in the password
		sudo -S gem install $1
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

PodInstall() {
	LOGFILE=$1
	mkdir -p .upshift/logs/
	pod install 2>&1 | tee $1	
}

SetupGradleW() {
	gradle wrapper
}

StartSimulator() {
	DEVICE=$1
	xcrun instruments -w "$1" 1>/dev/null 2>&1
}

CompileIOS() {
	PROJECT_TYPE=$1
	PROJECT_PATH=$2
	SCHEME=$3
	DEVICE=$4
	LOG_PATH=$5
	mkdir -p .upshift/logs/
	mkdir -p .upshift/build/
	set -o pipefail && xcodebuild -"$1" "$2" -scheme "$3" -hideShellScriptEnvironment -sdk iphonesimulator -destination "platform=iphonesimulator,name=$4" -derivedDataPath .upshift/build | tee "$5" | xcpretty
}