package bash

import (
	"bufio"
	"bytes"
	"errors"
	log "github.com/Sirupsen/logrus"
	// "io"
	"os/exec"
	"strings"
)

// Create a struct for errors
type BashErrors struct {
	errorMessage string
}

func (e *BashErrors) Error() string {
	return e.errorMessage
}

// Get started with init
func init() {

}

// Expect multiple commands
type Command struct {
	name string
	args []string
	next string // This defines how the pipe would work, it can be | or && for now
	cmd  *exec.Cmd
}

// Execute arbitary bash script like "ls -la"
func Bash(command string) (string, error) {

	// Inspiration from http://nathanleclaire.com/blog/2014/12/29/shelled-out-commands-in-golang/
	var outputBuffer bytes.Buffer
	var errorBuffer bytes.Buffer

	commands, err := createCommandQueue(command)
	if err != nil {
		return "", errors.New(err.Error())
	}
	// log.Info("CMDS", commands)

	if len(commands) == 0 {
		return "", &BashErrors{"There are no commands to process"}
	}

	errScanner := bufio.NewScanner(&errorBuffer)
	go func() {
		log.Info("I ran too")
		for errScanner.Scan() {
			log.Info("y", errScanner.Text())
		}
	}

	setupPipes(commands, errorBuffer)
	runCommands(commands, errorBuffer)

	return outputBuffer.String(), nil
}

func runCommands(commands []Command, errorBuffer bytes.Buffer) {
	for i, command := range commands {
		log.Info(i, command.cmd.Stdin)
		log.Info(i, command.cmd.Stdout)

		if err := command.cmd.Start(); err != nil {
			log.Error(i, "Problem starting command")
		}

		if err := command.cmd.Wait(); err != nil {
			log.Error(i, "Problem waiting for command ", err.Error())
		}
	}
}

func setupPipes(commands []Command, errorBuffer bytes.Buffer) {
	log.Info("CMD", commands)
	for i, command := range commands[:len(commands)-1] {
		out, _ := command.cmd.StdoutPipe()
		command.cmd.Stderr = &errorBuffer
		commands[i+1].cmd.Stdin = out

		scanner := bufio.NewScanner(out)
		go func() {
			log.Info("I did run")
			for scanner.Scan() {
				log.Info("x", scanner.Text())
			}
		}()
	}
}

func createCommandQueue(command string) ([]Command, error) {

	// Create an empty list of commands
	commands := []Command{}

	// Parse the command string to create one big array with all commands and args
	commandParams := strings.Fields(command)

	// Check if there are no commands, and just return if it is
	if len(commandParams) == 0 {
		// return commands, errors.New("The command was empty")
		return commands, &BashErrors{"The command was empty"}
	}

	//commands = append(commands, Command{name: commandParams[0], args: commandParams[1:]})

	inCommandSeq := 0
	tempCommand := Command{}

	for _, param := range commandParams {
		// Figure out if param is a name, arg or separator
		if inCommandSeq == 0 {
			tempCommand.name = param
		} else if param == "&&" {
			// Ignore
			log.WithFields(log.Fields{"package": "Bash", "command": command}).Error("We don't handle && in bash commands, keem them out")
			return commands, &BashErrors{"Found an unsupported action in the shell script"}
		} else if param == "|" {
			tempCommand.next = param
			tempCommand.cmd = exec.Command(tempCommand.name, tempCommand.args...)
			// Move tempCommand to commands and get a empty one again
			commands = append(commands, tempCommand)
			tempCommand.name = ""
			tempCommand.args = nil
			tempCommand.next = ""

			inCommandSeq = -1
		} else {
			tempCommand.args = append(tempCommand.args, param)
		}
		inCommandSeq = inCommandSeq + 1
	}

	tempCommand.cmd = exec.Command(tempCommand.name, tempCommand.args...)
	commands = append(commands, tempCommand)

	return commands, nil
}
