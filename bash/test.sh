#!/bin/bash
 set -v
# set -e

# AF=$(date +%Y%m%d)
# echo $AF

# function book {
#	AF=$(date +%Y%m%d)
#	echo $AF
# }
# book

# echo $1 $2 $3
# echo $# ### Number of parameters

# echo `uname`

# echo -e "Can you please enter two words? "; read word1 word2; echo "Here is your input: \"$word1\" \"$word2\""
# echo -e "Here's how you reac a line"; read; echo $REPLY
# echo -e "What are your favorite colours ? "; read -a colours; echo ${colours[0]}

# clear

# trap bashtrap INT
# bashtrap() {
#	echo "Ctrl+C detected, executing trap"
#}

#sleep 1

#for a in `seq 1 10`; do
#	echo "$a/10 to Exit"
#	sleep 1;
#done
#echo "Exit bash trap"

#seq 1 10

#linux=( 'Debian Linux' 'Redhat Linux' 'Ubuntu Linux' )
#items=${#linux[@]}
#for (( i=0;i<$items;i++)); do
#    echo ${linux[${i}]}
#done 

# if [ -d $directory ]; then
# if [ $choice -eq 2 ]; then
# -lt <, -gt >, -le <=, -ge >=, -eq ==, -ne !=
# Strings : =, !=, <, >, -n s1 (not empty), -z s1 (is empty)

#for f in $( ls /var/ ); do

# while [ $COUNT -gt 0 ]; do

# until [ $COUNT -gt 5 ]; do

#find $DIR -type f

# function function_A { echo $1 }

# select word in "linux" "bash" "scripting" "tutorial"; do
#	echo "The word you have selected is: $word"
#	break
#done

# echo "What is your preferred programming / scripting language"
# echo "1) bash"
# read case;
# case $case in
#     1) echo "You selected bash";;
# esac


# round floating point number with bash
#for bash_rounded_number in $(printf %.0f 3.3446); do
#	echo "Rounded number with bash:" $bash_rounded_number
#done 

# calculations using BC http://www.basicallytech.com/blog/archive/23/command-line-calculations-using-bc/
# https://linuxconfig.org/bash-scripting-tutorial

# more on bash - http://tldp.org/HOWTO/Bash-Prog-Intro-HOWTO.html#toc5
# even more - http://ryanstutorials.net/bash-scripting-tutorial/

# echo "1234561" | tr "1" "3" // Returns 3234563

# http://linux.die.net/abs-guide/x15683.html


#OUTPUT=$(adb devices | grep -v "List of devices attached")
#echo $OUTPUT

#   line=$(grep "revision" <<< "$packages")

			
#			emulator_name=$1
#
#            OUTPUT=$(ps -aef | grep emulator | grep "sdk/tools" -c)
#            if [ "$OUTPUT" == "0" ]; then
#              EMULATOR_RESULTS=$(nohup $ANDROID_HOME/tools/emulator -avd ${emulator_name} 2>emulator.log 1>emulator.log &)
#              # This is a big #HACK, only errors are returned in the first two seconds, I suck
#              sleep 2
#              EMULATOR_RESULTS=$(<emulator.log)
#              
#              if [ $(echo $EMULATOR_RESULTS | grep "PANIC" | wc -l) == "1" ]; then
#                echo $EMULATOR_RESULTS | grep "PANIC"
#                exit 1
#              else 
#                if [ $(echo $EMULATOR_RESULTS | grep "ERROR" | wc -l) == "1" ]; then
#                  echo $EMULATOR_RESULTS | grep "ERROR"
#                  exit 1
#              	else
#                	echo Loading the emulator
#                fi
#              fi
#            else
#              echo The emulator is already running
#            fi

#            DEVICES_FOUND=$(adb devices | grep -v "List of devices attached" | grep -v emulator -c)
#            let devices=DEVICES_FOUND-1;
#            echo $devices


#function getLastedSupportLibraryRevision {
#  packages=$(android list sdk)
#  echo $packages
#  #line=$(grep "Local Maven repository for Support Libraries, revision" <<< "$packages")
#  line=$(grep "revision" <<< "$packages")
#  echo $line;
#  line=$(echo $packages | grep "revision")
#  echo $line;
#  echo ${line: -2}
#}
#
#getLastedSupportLibraryRevision

# Black        0;30     Dark Gray     1;30
# Red          0;31     Light Red     1;31
# Green        0;32     Light Green   1;32
# Brown/Orange 0;33     Yellow        1;33
# Blue         0;34     Light Blue    1;34
# Purple       0;35     Light Purple  1;35
# Cyan         0;36     Light Cyan    1;36
# Light Gray   0;37     White         1;37

#            RED='\033[0;31m'
#            GREEN='\033[0;32m'
#            BLUE='\033[0;34m'
#            NOCOLOR='\033[0m'
# printf "Parameter ${BLUE}emailForSSHKey${NOCOLOR} not found\n"


# find . | grep -v "build" | grep AndroidManifest.xml | grep "app/src"






# /Users/sudhanshuraheja/.stepman/step_collections/1462621641/collection 


# stepman share start -c git@github.com:leftshifters/bitrise-steplib.git

# Go to repo and git tag -a 0.5.0 -m "messsage"
#stepman share create --tag 0.5.0 --git https://github.com/sudhanshuraheja/bitrise-install-android-build.git --stepid android-install

#stepman audit

#stepman share finish

# create a Pull Request



# Find where is XCode Installed
#XCODE_PATH=$(xcode-select --print-path)
#printf "$XCODE_PATH\n"

# Install the latest command line tools
#xcode-select --install
# If already installed # xcode-select: error: command line tools are already installed, use "Software Update" to install updates

# Switching to a different XCode
#sudo xcode-select -switch <path/to/>Xcode.app


# xcodebuild -list -project Deezeno.xcodeproj/
# xcodebuild -list -workspace Deezeno.xcworkspace/
# xcodebuild -scheme Deezeno build
# xcodebuild -scheme Deezeno clean
# xcodebuild -scheme Deezeno analyze


# echo "${machinePassword}" | sudo xcode-select -switch /Applications/Xcode.app



# Check .gitmodules if it exists and number of files greater than 0
#git submodule init
#git submodule update
#xcodebuild -scheme Deezeno build
# Catch errors and warnings

# Change bundle identifier
#/usr/libexec/PlistBuddy -c "Print :CFBundleIdentifier" "${info_plist_file}")"
#/usr/libexec/PlistBuddy -c "Set :CFBundleIdentifier ${bundle_identifier}" "${info_plist_file}"
# Info Plist is at /Deezeno/Info.plist

# Change bundle version
#/usr/libexec/PlistBuddy -c "Print CFBundleVersion" "${CONFIG_project_info_plist_path}"
#/usr/libexec/PlistBuddy -c "Set :CFBundleVersion ${CONFIG_new_bundle_version}" "${CONFIG_project_info_plist_path}"

# Change bundle version code
# ---- Current Bundle Short Version String:
#/usr/libexec/PlistBuddy -c "Print CFBundleShortVersionString" "${CONFIG_project_info_plist_path}"
#/usr/libexec/PlistBuddy -c "Set :CFBundleShortVersionString ${CONFIG_new_build_short_version_string}" "${CONFIG_project_info_plist_path}"

# Get XCode Version
#xcodebuild -version

#xcodebuild -showsdks
#OS X SDKs:
#	OS X 10.11                    	-sdk macosx10.11
#
#iOS SDKs:
#	iOS 9.3                       	-sdk iphoneos9.3
#
#iOS Simulator SDKs:
#	Simulator - iOS 9.3           	-sdk iphonesimulator9.3
#
#tvOS SDKs:
#	tvOS 9.2                      	-sdk appletvos9.2
#
#tvOS Simulator SDKs:
#	Simulator - tvOS 9.2          	-sdk appletvsimulator9.2
#
#watchOS SDKs:
#	watchOS 2.2                   	-sdk watchos2.2
#
#watchOS Simulator SDKs:
#	Simulator - watchOS 2.2       	-sdk watchsimulator2.2

# xcodebuild -showBuildSettings
#    ## AVAILABLE_PLATFORMS = appletvos appletvsimulator iphoneos iphonesimulator macosx watchos watchsimulator
#    # BUILD_DIR = /Users/sudhanshuraheja/Library/Developer/Xcode/DerivedData/Deezeno-fpdbioebihpiyrbfohmfktdwgmgf/Build/Products
#    CODE_SIGNING_ALLOWED = YES
#    CODE_SIGNING_REQUIRED = YES
#    CODE_SIGN_ENTITLEMENTS = GoJek/GO-JEK.entitlements
#  CODE_SIGN_IDENTITY = iPhone Developer
#    CORRESPONDING_SIMULATOR_SDK_NAME = iphonesimulator9.3
#    FULL_PRODUCT_NAME = Deezeno.app / GO-JEK.app
#  INFOPLIST_FILE = Deezeno/Info.plist
#  PRODUCT_BUNDLE_IDENTIFIER = com.leftshift.deezeno
#  PRODUCT_NAME = Deezeno
#    PRODUCT_SETTINGS_PATH = /Users/sudhanshuraheja/Code/bitrise/deezeno-ios/Deezeno/Info.plist
#  PROJECT_FILE_PATH = /Users/sudhanshuraheja/Code/bitrise/deezeno-ios/Deezeno.xcodeproj
#    SOURCE_ROOT = /Users/sudhanshuraheja/Code/bitrise/deezeno-ios
#    TARGETNAME = Deezeno
#    TARGET_NAME = Deezeno

#    FACEBOOK_APP_ID = 830829963645099
#    FLURRY_API_KEY = WC229WJF2BF6FFK767B3
#    GOOGLE_MAPS_API_KEY = AIzaSyBOaKFEQG3Tumx-WD4NAEZBSZd8unhPHBc

#xcodebuild -list
#Information about project "Deezeno":
#    Targets:
#        Deezeno
#        DeezenoTests
#        DeezenoUITests
#
#    Build Configurations:
#        Debug
#        Release
#
#    If no build configuration is specified and -scheme is not passed then "Release" is used.
#
#    Schemes:
#        Deezeno


#xcodebuild -version
#Xcode 7.3
#Build version 7D175

#//usr/libexec/PlistBuddy 


#set -o pipefail


#et -o pipefail && xcodebuild -workspace './GoJek.xcworkspace' -scheme 'GoJek Staging' -configuration 'Debug' -destination 'generic/platform=iOS' -archivePath '/var/lib/jenkins/Library/Developer/Xcode/Archives/2016-05-09/GO-JEK-Dev-b0cd652 2016-05-09 19.14.16.xcarchive' archive CODE_SIGN_IDENTITY='iPhone Developer: Gojek Ci (AKQYP36WDV)' | tee /var/lib/jenkins/Library/Logs/gym/GO-JEK-GoJek\ Staging.log | xcpretty[0m[0m




# xcodebuild -scheme "schemeName" -showBuildSettings >> mynew.xcconfig



# Sudhanshus-MacBook-Air:gojek-ios sudhanshuraheja$ git blame GoJek/NewBookingCompleteController.swift -e | grep " 158)"
# 163564e6 (<vinod@leftshift.io>                           2016-02-02 20:03:55 +0530  158)                     if checkErrorRequest(Int(response!.statusCode),json: JSON(JSONResponse!)) {


# https://gist.github.com/johanneswuerbach/5559514


# https://www.objc.io/issues/6-build-tools/travis-ci/


# agvtool what-marketing-version
# https://github.com/nomad/shenzhen/blob/master/lib/shenzhen/agvtool.rb


# https://www.objc.io/issues/6-build-tools/travis-ci/
# https://docs.travis-ci.com/user/customizing-the-build/
# https://gist.github.com/johanneswuerbach/5559514
# https://discuss.circleci.com/t/ios-code-signing/1231





# https://github.com/heroku/stack-images/blob/master/bin/cedar.sh


# https://github.com/ainoya/docker-android-project/blob/master/Dockerfile
# https://github.com/gfx/docker-android-project/blob/master/Dockerfile