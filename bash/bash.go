package bash

import (
	"bufio"
	// "bytes"
	"errors"
	log "github.com/Sirupsen/logrus"
	"io"
	"os/exec"
	"strings"
)

// Get started with init
func init() {

}

// Execute arbitary bash script like "ls -la"
func Bash(command string) (string, error) {

	// Inspiration from http://nathanleclaire.com/blog/2014/12/29/shelled-out-commands-in-golang/
	// var output bytes.Buffer
	ch := make(chan string, 1)

	// Covert command in []string
	com := strings.Fields(command)

	if len(com) == 0 {
		return "", errors.New("You need to enter something for this to work")
	}

	name := com[0]  // Pick the name from the first value
	args := com[1:] // Rest everything is an arg)

	// Time to run the process
	cmd := exec.Command(name, args...)

	// Get an output pipe so that you can read Stdout
	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		return "", errors.New("We could not connect to the stdout for " + name)
	}
	setupScanner(stdOut, ch)

	// Get an output pipe so that you can reach Stderr
	stdErr, err := cmd.StderrPipe()
	if err != nil {
		return "", errors.New("We could not connect to stderr for " + name)
	}
	setupScanner(stdErr, ch)

	// Start the command
	err = cmd.Start()
	if err != nil {
		return "", errors.New("There was a problem starting the command " + name)
	}

	// Fire up wait to finish up and close things
	err = cmd.Wait()
	if err != nil {
		return "", errors.New("It wasn't the best of endings, there were some problems wrapping up, but isn't too much to worry about")
	}

	return <-ch, nil
}

func setupScanner(pipe io.ReadCloser, ch chan string) {
	scanner := bufio.NewScanner(pipe)
	go func() {
		for scanner.Scan() {
			readText := scanner.Text()
			log.Info("Scanner <", readText, ">")
			ch <- readText
		}
		// if scanner.Err() != nil {
		// 	log.Info("ScannerErr <", scanner.Err().Error(), ">")
		// 	ch <- scanner.Err().Error()
		// }
	}()
}
