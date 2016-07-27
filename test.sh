#!/usr/bin/env bash

set -e
echo "" > coverage.txt

for d in $(find ./* -maxdepth 10 -type d | grep -v actions); do
    printf "Testing $d\n"
    if ls $d/*.go &> /dev/null; then
        printf "Checking $d\n"
        go test -v -timeout 99999s -coverprofile=profile.out -covermode=atomic $d
        if [ -f profile.out ]; then
            printf "$d/profile.out moving to coverage.txt\n"
            cat profile.out >> coverage.txt
            rm profile.out
        else
            printf "$d/profile.out not found\n"
        fi
    fi
done
