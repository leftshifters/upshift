main() {
    echo "Hello World" | reverse
}

gradleTasks() {
	gradle tasks 2>&1 | tee ./this-file-$1
}