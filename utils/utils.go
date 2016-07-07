package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	c "upshift/colours"
)

// Log Message, this shows up in green and underlined
func LogMessage(message string) {
	fmt.Println("✅  " + c.Green + c.Bold + message + c.Default)
}

// Log Information, this shows up in blue
func LogInfo(message string) {
	fmt.Println("🔰  " + c.Green + c.Bold + "Maybe you should " + c.Underline + "know this" + c.Default)
	fmt.Println(message)
}

// Log an error, show them this shit in color, red most probably
func LogError(message string) {
	fmt.Println("☎️  " + c.Red + c.Bold + "Shit! Something broke" + c.Default)
	fmt.Println(message)
}

// Create a string array from an bash output
func CreateList(data string, ignore []string) []string {
	rows := strings.Split(data, "\n")
	var items []string
	for _, row := range rows {
		ignoredRow := false
		for _, ign := range ignore {
			if strings.Contains(row, ign) == true {
				ignoredRow = true
			}
			row = strings.TrimSpace(strings.Replace(row, ign, "", 1))
		}
		if row != "" && ignoredRow == false {
			items = append(items, row)
		}
	}
	return items
}

// Read last few bytes of a file
func FileTail(filePath string, size int64) (string, error) {
	// Check if file exits, if it doesn't just return an error
	if FileExists(filePath) == false {
		return "", errors.New("File does not exist " + filePath)
	}

	// If file exists, go ahead and read the shit out of it
	file, err := os.Open(filePath)
	if err != nil {
		return "", errors.New("Could not open file " + filePath + "\n" + err.Error())
	}
	// Close the file when you are about to close the function
	defer file.Close()

	// Get the file size from the OS
	fileStat, err := os.Stat(filePath)
	// Calculate where you want to start reading
	startByte := fileStat.Size() - size
	// Calculate the amount of data that you need read
	readSize := size
	// If the file size is less than the amount of data you're trying to read, adjust the two above
	if startByte < 0 {
		startByte = 0
		readSize = fileStat.Size()
	}

	// Make a byte buffer to read the amount of data
	byteBuffer := make([]byte, readSize)
	// Ignore the number of bytes read, we just need the effing output
	_, err = file.ReadAt(byteBuffer, startByte)
	if err != nil {
		// Reading failed, dump the error
		return "", errors.New("Could not read file" + filePath + "\n" + err.Error())
	}

	return string(byteBuffer), nil
}

// Read a file if it exists
func FileRead(filePath string) (string, error) {
	// Check if file exits, if it doesn't just return an error
	if FileExists(filePath) == false {
		return "", errors.New("File does not exist " + filePath)
	}

	// If file exists, go ahead and read the shit out of it
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", errors.New("Could not read file" + filePath + err.Error())
	}

	return string(fileBytes), nil
}

// Simply to check if a file exists or not
func FileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	if err == nil {
		return true
	} else {
		return false
	}
}
