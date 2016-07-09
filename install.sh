#!/usr/bin/env bash

# 7. Application Version
version="0.8.4"

# Check which OS are we on and download the appropriate binary
OS=$(uname)
INSTALL_URL=""

printf "Detected OS $OS\n"

if [ "$OS" == "Darwin" ]; then
  INSTALL_URL="https://github.com/leftshifters/upshift/releases/download/${version}/upshift-darwin-${version}"
else
  INSTALL_URL="https://github.com/leftshifters/upshift/releases/download/${version}/upshift-linux-${version}"
fi

printf "Downloading from $INSTALL_URL\n"
INSTALL=$(curl --show-error --progress-bar --fail --location "${INSTALL_URL}" > upshift)
printf "Download completed\n"

# Only these are available in $PATH on a fresh system
# /usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin
# We will install in /usr/local/bin

# Check if /usr/local/upshift directory exists, if not create it
if [ ! -d "/usr/local/upshift/${version}" ]; then
  # Directory doesn't exist, create it
  sudo mkdir -p /usr/local/upshift/${version}
  printf "Created /usr/local/upshift/${version}\n"
fi

# Throw yourself into the folder above, it not already there
if [ ! -f "/usr/local/upshift/${version}/upshift" ]; then
  # Copy away
  sudo cp -rf ./upshift /usr/local/upshift/${version}/upshift

  # Now add a link to the above file in /usr/local/bin
  if [ -L "/usr/local/bin/upshift" ]; then
    # Remove only if it exists
    sudo rm -f /usr/local/bin/upshift
  fi

  sudo ln -s /usr/local/upshift/${version}/upshift /usr/local/bin
  sudo chmod +x /usr/local/bin/upshift
  printf "Created /usr/local/bin/upshift\n"
  rm upshift

  printf "Installation has been ${greenColour}successfully${noColour} completed\n"
else
  printf "This version is already ${greenColour}installed${noColour}\n"
fi
