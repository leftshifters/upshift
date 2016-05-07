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



