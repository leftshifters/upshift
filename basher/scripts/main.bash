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

ArchiveIOS() {
	PROJECT_TYPE=$1
	PROJECT_PATH=$2
	SCHEME=$3
	PROJECT_NAME=$4
	LOG_PATH=$5
	mkdir -p .upshift/logs/
	mkdir -p .upshift/build/
	set -o pipefail && xcodebuild -"$1" "$2" -scheme "$3" -derivedDataPath .upshift/build -archivePath .upshift/$4.xcarchive archive | tee "$5" | xcpretty
}

ExportIOS() {
	PROJECT_NAME=$1
	LOG_PATH=$2
	mkdir -p .upshift/logs/
	mkdir -p .upshift/build/
	set -o pipefail && xcodebuild -exportArchive -exportOptionsPlist .private/export.plist -archivePath .upshift/$1.xcarchive -exportPath .upshift/$1.ipa 2>&1 | tee "$2" | xcpretty
}


# // set -o pipefail && xcodebuild -exportArchive -exportOptionsPlist "profiles/export.plist" -archivePath "build/${projectName}.xcarchive" -exportPath "build/${projectName}.ipa" 2>&1 | tee "xcode-archive-${TIMESHTAMP}.log" | xcpretty

PopulateProvisioningProfiles() {
	# Get the UUID from .private
	# https://gist.github.com/mxpr/8208289a63ca4e3a35a4
	# Loop through all files, if you get a UDID, add them to the list of profiles
	if [ -d "./.private/" ]; then
		# The .private folder exists
		foundProfiles=false
		for profileName in .private/*.mobileprovision; do
			uuid=$(/usr/libexec/PlistBuddy -c 'Print UUID' /dev/stdin <<< $(security cms -D -i "${profileName}" 2>/dev/null))

			# If a UUID exist, then copy it, if it hasn't already been copied
			if [ "${uuid}" != "" ]; then
				# Copy file to UUID folder
				`cp -f ${profileName} ~/Library/MobileDevice/Provisioning\ Profiles/${uuid}.mobileprovision`
				printf "Transporting file ${uuid}.mobileprovision from .private to Library\n"
				foundProfiles=true
			fi
		done

		if [ "$foundProfiles" == true ]; then
			exit 0
		else
			printf "Dude, You need to add some provisioning profiles in .private\n"
			exit 1
		fi
	else
		printf "Hey, you need to add your provisioning profiles in .private\n"
		exit 1
	fi
}