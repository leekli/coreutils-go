package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Retrieve any args given by the user not incl. file name (which is arg 0)
	userArgsGiven := os.Args[1:]

	// If no args given, display the usage help screen & exit
	if len(userArgsGiven) == 0 {
		cutMessage := CutHelpMessage()

		fmt.Println(cutMessage)

		return
	}

	// Get the file path/name given as the final argument and check it exists, exit if not
	var filePath string

	if len(userArgsGiven) > 1 {
		filePath = userArgsGiven[len(userArgsGiven) - 1]
	}

	fileExists := CheckFileExists(filePath)

	if !fileExists {
		fmt.Printf("cut: %v: No such file or directory\n", filePath)

		return
	}

	var fieldToPrint int
	var delimiterToUse string

	// The -f flag
	// -f list 		The list specifies fields, separated in the input by the field delimiter character (see the -d option).  Output fields are separated by a single occurrence of the field delimiter character.
	if strings.Contains(userArgsGiven[0], "-f") {
		reNumsOnly := regexp.MustCompile("[0-9]+")
		fNumsList := reNumsOnly.FindAllString(userArgsGiven[0], -1)
		fieldNum, err := strconv.ParseInt(fNumsList[0], 10, 32)

		if err != nil {	
			fmt.Println("There was an error with the field number provided with the -f flag")

			return
		}

		fieldToPrint = int(fieldNum)
	}

	// The -d flag
	// -d delim 	Use delim as the field delimiter character instead of the tab character.
	if strings.Contains(userArgsGiven[1], "-d") {
		reDelimChar := regexp.MustCompile("^-d(.)$")
		dDelimCharMatch := reDelimChar.FindStringSubmatch(userArgsGiven[1])

		if len(dDelimCharMatch) > 1 {
			delimiterToUse = dDelimCharMatch[1]
		} else {
			delimiterToUse = "\t"
		}
	}

	returnCode, returnMessage := PrintBySpecifiedField(filePath, fieldToPrint, delimiterToUse)

	if returnCode < 0 {
		fmt.Println(returnMessage)
	}

	fmt.Println(returnMessage)
}

func CutHelpMessage() string {
	var cutMessage string;

	cutMessage += "usage: "

	cutMessage += "	cut -f list [-d delim]"

	return cutMessage
}

func CheckFileExists (filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
        return true
    } else if os.IsNotExist(err) {
        return false
    }

	return false
}

func PrintBySpecifiedField(filePath string, fieldNum int, delimiterChar string) (int, string) {
	// Set default delimiter of tab-spaced if still empty
	if delimiterChar == "" {
		delimiterChar = "\t"
	}

	if fieldNum == 0 {
		return -1, "cut: values may not include zero\n"
	}

	file, err := os.Open(filePath)

    if err != nil {
		errMsg := fmt.Sprintf("cut: %v: Error opening file", err)
		
        return -1, errMsg
    }

	var textToReturn string

    defer file.Close()
	scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()

        fields := strings.Split(line, delimiterChar)

        if fieldNum > len(fields) {
			textToReturn += fmt.Sprintln(" ")

            continue
        }

		line = fmt.Sprintln(fields[fieldNum - 1])
        textToReturn += line
    }

    if err := scanner.Err(); err != nil {
		errMsg := fmt.Sprintf("cut: %v: Error opening file", err)
		
        return -1, errMsg
    }

	finalString := strings.Trim(textToReturn, "\n")

	return 0, finalString
}