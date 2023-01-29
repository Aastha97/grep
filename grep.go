package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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

func traverseDir(dirName string) []string {
	var files []string

	err := filepath.Walk(dirName, func(path string, info os.FileInfo, err error) error {

		if err != nil {

			fmt.Println(err)
			return nil
		}

		if !info.IsDir() && filepath.Ext(path) == ".txt" {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	return files
}
func writeFile(newString []string, fileName string) (string, error) {
	var warning string = "\n" + fileName + " file already exists, cannot write in file"
	content, err := os.Open(fileName)
	if err == nil {
		fmt.Println(warning)
		return warning, err
	}
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

func readFileLineByLine(fileName, searchStr string, isCaseInSensitivity, isWordMatch, recursive bool) []string {
	var result []string
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		output := searchString(searchStr, line, isCaseInSensitivity, isWordMatch)
		if output != "" {
			if recursive {
				result = append(result, fileName+": "+output)
			} else {
				result = append(result, output)
			}

		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func finalResult(result string) {
	fmt.Println("\n" + result)
}

func recursiveCallFromDir(dirName, searchStr string, isCaseInSensitivity, isWordMatch, recursive bool) []string {
	var result, finalList []string
	files := traverseDir(dirName)
	for _, fileName := range files {
		result = readFileLineByLine(fileName, searchStr, isCaseInSensitivity, isWordMatch, recursive)
		finalResult(strings.Join(result, "\n"))
		finalList = append(finalList, strings.Join(result, "\n"))
	}
	return finalList
}
