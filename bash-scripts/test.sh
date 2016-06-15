#!/bin/bash

# Find where is XCode Installed
#XCODE_PATH=$(xcode-select --print-path)

# xcodebuild -list -project Deezeno.xcodeproj/
# xcodebuild -list -workspace Deezeno.xcworkspace/
# xcodebuild -scheme Deezeno build
# xcodebuild -scheme Deezeno clean
# xcodebuild -scheme Deezeno analyze

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

# xcodebuild -scheme "schemeName" -showBuildSettings >> mynew.xcconfig

# Sudhanshus-MacBook-Air:gojek-ios sudhanshuraheja$ git blame GoJek/NewBookingCompleteController.swift -e | grep " 158)"
# 163564e6 (<vinod@leftshift.io>                           2016-02-02 20:03:55 +0530  158)                     if checkErrorRequest(Int(

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

docker-machine create --driver virtualbox default

docker run -it ubuntu bash

docker run -d -P -v $HOME/gocode/bin:/usr/local/upshift/bin --name upshift ubuntu:16.04 /usr/local/upshift/bin/upshift -v
docker run -P -v $HOME/gocode/bin:/usr/local/upshift/bin -it --name upshift ubuntu:16.04 /bin/bash
docker rm upshift
#docker run --name upshift ubuntu:16.04 upshift -v
docker run --name upshift ubuntu:16.04 upshift -v && docker rm upshift
#docker rm upshift
#-d to run in background
#-P to expose ports

docker build -t upshift:latest .

env GOOS=linux GOARCH=amd64 go build -v upshift

docker run --name upshift ubuntu:16.04

