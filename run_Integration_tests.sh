#!/usr/bin/env bash

ColorOff='\033[0m'
Red='\033[0;31m'
Green='\033[0;32m'
Cyan='\033[0;36m'

function success {
  echo -e "${Green}success${ColorOff}"
}

function failed {
  echo -e "${Red}failed${ColorOff}"
}

function assertEquals {
  if [[ "$1" == "$2" ]]
  then
    success
  else
    failed
  fi
}

function assertFileExist {
  if test -f "$1"
  then
    success
  else
    failed
  fi
}

# build
go build -o intergrationTests/commandor .
containerHash=$(docker build -q intergrationTests)

# run tests
echo -en "test: ${Cyan}commandor touch${ColorOff} - "
output=$(docker run --rm -it "${containerHash}" commandor touch)
output=$(echo "${output}" | sed 's/.$//')
assertEquals "${output}" "hello world"

echo -en "test: ${Cyan}commandor go-run${ColorOff} - "
output=$(docker run --rm -it "${containerHash}" commandor go-run)
output=$(echo "${output}" | sed 's/.$//')
assertEquals "${output}" "hello world"

# clean
rm intergrationTests/commandor