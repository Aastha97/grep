package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func searchString(searchStr, content string, isCaseInsensitive, isWordMatch bool) []string {
	var caseInsensitivityFlag string = ""
	var wordMatchFlag string = ""
	var finalList []string

	if isCaseInsensitive {
		caseInsensitivityFlag = "(?i)"
	}
	if isWordMatch {
		wordMatchFlag = "\\b"
	}

	list := strings.Split(content, "\r\n")
	for _, str := range list {
		if strings.Contains(str, searchStr) {
			finalList = append(finalList, str)
			break
		}
		for _, word := range strings.Fields(str) {
			matched, _ := regexp.MatchString(caseInsensitivityFlag+wordMatchFlag+searchStr+wordMatchFlag, word)
			if matched {
				finalList = append(finalList, str)
				break
			}

		}
	}
	return finalList
}

func readFile(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return "Error reading file", err
	}
	content := string(data)
	return content, err
}

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

func finalResult(result string, err error) {
	if err == nil {
		fmt.Println("\n" + result)
	} else {
		fmt.Println(err)
	}
}
