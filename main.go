package main

import (
	bash "upshift/bash"
)

func main() {
	bash.Bash("")
	bash.Bash("ls")
	bash.Bash("ls -la")
	bash.Bash("ls -la -la")
}
