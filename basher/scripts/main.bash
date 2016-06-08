UpgradeScript() {
	curl -fsSL https://raw.githubusercontent.com/leftshifters/upshift/master/upshift > upshift.temp && chmod +x upshift.temp && ./upshift.temp install && rm upshift.temp
}

SetupXcpretty() {
	CI = $1
	RootPassword = $2

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