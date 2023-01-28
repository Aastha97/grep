package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func searchString(searchStr, line string, isCaseInsensitive, isWordMatch bool) string {
	var caseInsensitivityFlag, wordMatchFlag, output string = "", "", ""

	if isCaseInsensitive {
		caseInsensitivityFlag = "(?i)"
	}
	if isWordMatch {
		wordMatchFlag = "\\b"
	}

	matched, _ := regexp.MatchString(caseInsensitivityFlag+wordMatchFlag+searchStr+wordMatchFlag, line)
	if matched {
		output = line
	}

	return output
}

// func readFile(fileName string) (string, error) {
// 	data, err := os.ReadFile(fileName)
// 	if err != nil {
// 		return "Error reading file", err
// 	}
// 	content := string(data)
// 	return content, err
// }

func writeFile(newString []string, fileName string) (string, error) {
	content, err := os.Open(fileName)
	var warning string = fileName + " file already exists"
	if err != nil {
		content, _ := os.Create(fileName)
		for _, searchStr := range newString {
			content.Write([]byte(searchStr + "\n"))
		}
		fileScanner := bufio.NewScanner(content)
		return fileScanner.Text(), nil
	}
	defer content.Close()
	return warning, err
}

// func readFileLineByLine(fileName string) (string, error) {
// 	var line string
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		fmt.Println(err)
// 		return "", err
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	if scanner.Scan() {
// 		line = scanner.Text()
// 	}
// 	if err := scanner.Err(); err != nil {
// 		return "", err
// 	}
// 	return line, nil
// }

func finalResult(result string) {
	fmt.Println("\n" + result)
}
