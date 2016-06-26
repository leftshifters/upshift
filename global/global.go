package global

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
	// "time"
	"upshift/basher"
	"upshift/utils"
)

var conf Config

// Define structs for Config
type Config struct {
	// Save when was android sdk last updated
	AndroidSDKUpdatedTime int32
	IOSDeveloperAccounts  string
	IOSCertificatePath    string
}

// conf.AndroidSDKUpdated = int32(time.Now().Unix())

func Get() (Config, error) {

	// Run the basher script to make sure the folder and file exits
	basher.Run("UpshiftConfig", []string{})

	// See if a TOML file is available in this folder
	filePath := os.Getenv("HOME") + "/.upshift/config"
	fmt.Println("Reading the global config file at " + filePath)
	tomlBytes, err := utils.ReadIfFileExists(filePath)
	if err != nil {
		fmt.Println("The global config seems empty, this machine isn't used much it seems")
	}

	// If you get data, try to decode it and save into our conf object
	tomlData := string(tomlBytes)
	_, err = toml.Decode(tomlData, &conf)
	if err != nil {
		return conf, errors.New("Could not decode the toml file at " + filePath)
	}

	return conf, nil
}

func Set(c Config) error {

	// Run the basher script to make sure the folder and file exits
	basher.Run("UpshiftConfig", []string{})

	// Create the file name so that you don't have to write this over and over again
	filePath := os.Getenv("HOME") + "/.upshift/config"

	// Create a byte buffer
	var tomlBuffer bytes.Buffer
	e := toml.NewEncoder(&tomlBuffer)
	err := e.Encode(c)
	if err != nil {
		return err
	}

	// Write to disk
	err = ioutil.WriteFile(filePath, []byte(tomlBuffer.String()), 0644)
	if err != nil {
		return errors.New("We could not write the config file, the OS told us to get lost\n" + err.Error())
	}

	return nil
}
