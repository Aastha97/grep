package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func searchString(searchStr, content string) ([]string, error) {
	list := strings.Split(content, "\r\n")
	var finalList []string
	for _, str := range list {
		for _, word := range strings.Fields(str) {
			if word == searchStr {
				finalList = append(finalList, str)
			}
		}
	}
	return finalList, nil
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
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}

// func traverseAllfilesInExecDirectory(searchStr string, option int) {
// 	exec, err := os.Executable()
// 	if err != nil {
// 		panic(err)
// 	}
// 	dir := filepath.Dir(exec)

// 	files, err := os.ReadDir(dir)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	for _, file := range files {
// 		content, err := readFile(file.Name())
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		output, _ := searchString(searchStr, content)
// 		fmt.Println(file.Name() + ":" + strings.Join(output, "\n"))
// 	}

// }
