package main

import(
	"fmt"
	"strings"
	// "os/exec"
	// "bufio"
)

// Execute arbitary bash script like "ls -la"
func Bash(command string) {
	var commandParams []string
	commandParams = strings.Fields(command)

	if len(commandParams) > 0 {
		fmt.Println("CMD: ", commandParams[:1])
		fmt.Println("PRM: ", commandParams[1:])		
	} else {
		fmt.Println("There was nothing in the command")
	}

}

func main() {
	Bash("")
	Bash("ls")
	Bash("ls -la")
	Bash("ls -la -la")
}