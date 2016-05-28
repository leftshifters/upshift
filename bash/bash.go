package bash

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func init() {
	fmt.Println("Initializing Bash Package")
}

// Execute arbitary bash script like "ls -la"
func Bash(command string) string {

	// Literally flicked from http://nathanleclaire.com/blog/2014/12/29/shelled-out-commands-in-golang/

	commandParams := strings.Fields(command)

	if len(commandParams) == 0 {
		// TODO : Find better ways to show and pass errors
		fmt.Println("There was nothing in the command")
		return ""
	}

	// fmt.Println("CMD: ", commandParams[:1])
	fmt.Println("PRM: ", commandParams[1:])

	cmd := exec.Command(commandParams[0], commandParams[1:]...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		os.Exit(1)
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error waiting for command", err)
		os.Exit(1)
	}

	return command
}

// func main() {
// 	Bash("")
// 	Bash("ls")
// 	Bash("ls -la")
// 	Bash("ls -la -la")
// }
