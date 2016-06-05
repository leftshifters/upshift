main() {
    echo "Hello World" | reverse
}

gradleTasks() {
	gradle tasks 2>&1 | tee ./this-file-$1
}

upgradeScript() {
	curl -fsSL https://raw.githubusercontent.com/leftshifters/upshift/master/upshift > upshift.temp && chmod +x upshift.temp && ./upshift.temp install && rm upshift.temp
}