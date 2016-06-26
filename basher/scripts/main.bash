MakeFolders() {
	mkdir -p .upshift/logs/
	mkdir -p .upshift/build/
}

UpgradeScript() {
	curl -fsSL https://raw.githubusercontent.com/leftshifters/upshift/master/install.sh > upshift.temp && chmod +x upshift.temp && ./upshift.temp && rm upshift.temp
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
	MakeFolders
	git pull $1 $2 2>&1 | tee $3
}

GitSubmoduleInit() {
	LOGFILE=$1
	MakeFolders
	git submodule init 2>&1 | tee $1
}

GitSubmoduleUpdate() {
	LOGFILE=$1
	MakeFolders
	git submodule update 2>&1 | tee $1
}

PodInstall() {
	LOGFILE=$1
	MakeFolders
	pod install 2>&1 | tee $1	
}

PodRepoUpdate() {
	LOGFILE=$1
	UpshiftConfig
	pod repo update --verbose 2>&1 | tee $1	
}

SetupGradleW() {
	MakeFolders
	gradle wrapper
}

StartSimulator() {
	DEVICE=$1
	MakeFolders
	xcrun instruments -w "$1" 1>/dev/null 2>&1
}

CompileIOS() {
	PROJECT_TYPE=$1
	PROJECT_PATH=$2
	SCHEME=$3
	DEVICE=$4
	LOG_PATH=$5
	MakeFolders
	set -o pipefail && xcodebuild -"$1" "$2" -scheme "$3" -hideShellScriptEnvironment -sdk iphonesimulator -destination "platform=iphonesimulator,name=$4" -derivedDataPath .upshift/build | tee "$5" | xcpretty
}

ArchiveIOS() {
	PROJECT_TYPE=$1
	PROJECT_PATH=$2
	SCHEME=$3
	PROJECT_NAME=$4
	LOG_PATH=$5
	MakeFolders
	set -o pipefail && xcodebuild -"$1" "$2" -scheme "$3" -derivedDataPath .upshift/build -archivePath .upshift/$4.xcarchive archive | tee "$5" | xcpretty
}

ExportIOS() {
	PROJECT_NAME=$1
	LOG_PATH=$2
	MakeFolders
	set -o pipefail && xcodebuild -exportArchive -exportOptionsPlist .private/export.plist -archivePath .upshift/$1.xcarchive -exportPath .upshift/ 2>&1 | tee "$2" | xcpretty
}

FetchAndRepairProvisioningProfiles() {
	# this uses sigh
	# download all into .private folder
	# and then run the function PopulateProvisioningProfiles
	ACCOUNT_EMAIL=$1
	sigh repair -u $1
	sigh download_all -u $1
	mkdir -p ./.private
	mv *.mobileprovision .private
	PopulateProvisioningProfiles
}

FindProvisioningProfile() {
	DEVELOPER_ACCOUNT=$1
	BUNDLE_IDENTIFIER=$2
	sigh -u $1 -a $2 --adhoc
}

PopulateProvisioningProfiles() {
	# Get the UUID from .private
	# https://gist.github.com/mxpr/8208289a63ca4e3a35a4
	# Loop through all files, if you get a UDID, add them to the list of profiles
	MakeFolders
	if [ -d "./.private/" ]; then
		# The .private folder exists
		foundProfiles=false
		for profileName in .private/*.mobileprovision; do
			uuid=$(/usr/libexec/PlistBuddy -c 'Print UUID' /dev/stdin <<< $(security cms -D -i "${profileName}" 2>/dev/null))

			# If a UUID exist, then copy it, if it hasn't already been copied
			if [ "${uuid}" != "" ]; then
				if [ -f ~/Library/MobileDevice/Provisioning\ Profiles/${uuid}.mobileprovision ]; then
					# Silently ignore
					printf "${profileName} exists in the Library\n"
				else
					# Copy file to UUID folder
					cp -f "${profileName}" ~/Library/MobileDevice/Provisioning\ Profiles/${uuid}.mobileprovision
					printf "Moving ${profileName} [${uuid}] to Library\n"
				fi
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

AndroidClean() {
	LOG_PATH=$1
	MakeFolders
	./gradlew clean 2>&1 | tee $1
}

AndroidUninstall() {
	LOG_PATH=$1
	MakeFolders
	./gradlew uninstallAll 2>&1 | tee $1
}

AndroidInstallDebug() {
	LOG_PATH=$1
	MakeFolders
	./gradlew installDebug --stacktrace 2>&1 | tee $1
}

AndroidAssemble() {
	LOG_PATH=$1
	MakeFolders
	./gradlew assemble --stacktrace 2>&1 | tee $1
}

AndroidLint() {
	LOG_PATH=$1
	MakeFolders
	./gradlew lint 2>&1 | tee $1
}

AndroidStartActivity() {
	PACKAGE=$1
	MAIN_ACTIVITY=$2
	MakeFolders
	adb shell am start -n $1/$1.$2
}

AndroidLaunchEmulator() {

	MakeFolders

	redColour='\033[0;31m'
	greenColour='\033[0;32m'
	blueColour='\033[0;34m'
	noColour='\033[0m'

	EMULATOR_NAME=$1
	LOG_PATH=$2
	EMULATOR_RESULTS=$(nohup "$ANDROID_HOME/tools/emulator" -avd "$1" 2>$2 1>$2 &)

	# TODO : This is a big #HACK, only errors are returned in the first two seconds, I suck and I don't know a way out
	sleep 2
	EMULATOR_RESULTS=$(<$2)

	# Check if there was a PANIC [to test this, put in the wrong emulator name]
	PANIC_COUNT=$(echo "${EMULATOR_RESULTS}" | grep "PANIC" -c)

	# If there was a panic, throw it at the user, they messed up, not us
	if [ "${PANIC_COUNT}" -gt 0 ]; then
		printf "The emulator won't load up. Maybe the ${redColour}emulatorName${noColour} key isn't correct\n"
		printf "Here's what the log says:\n\n"
		printf "${EMULATOR_RESULTS}\n\n"
		exit 10 # 10 for Panic
	else
		# Seems like there is no panic, let's check for errors
		ERROR_COUNT=$(echo "${EMULATOR_RESULTS}" | grep "ERROR" -c)
		if [ "${ERROR_COUNT}" -gt 0 ]; then
			printf "${EMULATOR_RESULTS}\n\n"
			exit 11 # 11 for Error
		else
			printf "All set baby, ${greenColour}starting${noColour} to load up the emulator\n"
			# ADB gives this option to wait for the device till it comes up, let's just depend on it,
			# this is really mess with us when there is a problem with the emulator fails to load because of it's own bugs
			adb wait-for-device

			# Now that the device is available, the worst is over
			# Check if the emulator has finished botting, if not, sleedp for 2 seconds and try this again, this is our final trigger
			printf "Seems like the device is now ${greenColour}available${noColour}, we are getting close\n"
			SCREEN_LOADING=$(adb shell getprop sys.boot_completed | tr -d '\r')
			while [ "$SCREEN_LOADING" != "1" ]; do
				sleep 4
				printf "Check if the emulator has finished booting, why is this thing so ${blueColour}damn${noColour} slow?\n"
				SCREEN_LOADING=$(adb shell getprop sys.boot_completed | tr -d '\r')
			done

			# Alright, everything is now done. This is just used to unlock the device.
			printf "Almost ${greenColour}done${noColour}, touch the device once\n"
			adb shell input keyevent 82
			printf "${greenColour}Super!${noColour} The emulator is now running. You're one lucky person 🍺\n"
			exit 0
		fi
	fi

}

AndroidUpgradeSDK() {
	# Copied from here - http://stackoverflow.com/a/31900427/57914
	# I still don't know how does this work
	( sleep 5 && while [ 1 ]; do sleep 1; echo y; done ) | android update sdk --no-ui
}

AndroidInstallSDK() {
	echo y | android update sdk --all --no-ui --filter "android-23"
	echo y | android update sdk --all --no-ui --filter "android-22"
	echo y | android update sdk --all --no-ui --filter "android-21"
	echo y | android update sdk --all --no-ui --filter "android-20"
	echo y | android update sdk --all --no-ui --filter "android-19"
	echo y | android update sdk --all --no-ui --filter "android-18"
	echo y | android update sdk --all --no-ui --filter "android-17"
	echo y | android update sdk --all --no-ui --filter "android-16"
	echo y | android update sdk --all --no-ui --filter "tools"
	echo y | android update sdk --all --no-ui --filter "platform-tools"
	echo y | android update sdk --all --no-ui --filter "extra-android-m2repository"
	echo y | android update sdk --all --no-ui --filter "extra-google-m2repository"
	echo y | android update sdk --all --no-ui --filter "sys-img-armeabi-v7a-android-22"
}

UpshiftConfig() {
	mkdir -p ~/.upshift
	touch ~/.upshift/config
}

InstallCertificates() {
	BASE_PATH=$1
	security import $1/apple.cer -k ~/Library/Keychains/login.keychain -T /usr/bin/codesign -T /usr/bin/security
	security import $1/distribution.p12 -k ~/Library/Keychains/login.keychain -T /usr/bin/codesign -T /usr/bin/security -P ""
	security import $1/distribution.cer -k ~/Library/Keychains/login.keychain -T /usr/bin/codesign -T /usr/bin/security
}

CreateAppOnItunes() {
	DEVELOPER_ACCOUNT=$1
	BUNDLE_IDENTIFIER=$2
	PROJECT_NAME=$3
	produce -u $1 -a $2 --app_name "$3"
}

UploadIPAoniTunes() {
	DEVELOPER_ACCOUNT=$1
	IPA_PATH=$2
	pilot upload -u $1 -i "$2" pi--verbose
}