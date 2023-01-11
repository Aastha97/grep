package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Choose one:\n 1. search string \n 2. sub string \n 3. Write text in a new file ")
	var option int
	fmt.Scanln(&option)
	var str, fileName string
	fmt.Scan(&str, &fileName)
	switch option {
	case 1:
		result, err := searchString(str, fileName)

		if err == nil {
			fmt.Println(result)
		} else {
			fmt.Println(err)
		}
	case 2:
		result, err := searchSubString(str, fileName)

		if err == nil {
			fmt.Println(result)
		} else {
			fmt.Println(err)
		}
	case 3:
		result, err := addStrToNewFile(str, fileName)

		if err == nil {
			fmt.Println(result)
		} else {
			fmt.Println(err)
		}
	}

}

func searchString(searchStr, fileName string) (string, error) {
	var result string
	resultTrue := "I found the " + searchStr + " in the " + fileName + " file"
	resultFalse := "Cannot find the " + searchStr + " in the " + fileName + " file"

	content, err := os.Open(fileName)

	if err != nil {
		return "", err
	}
	defer content.Close()

	fileScanner := bufio.NewScanner(content)

	for fileScanner.Scan() {
		if fileScanner.Text() == searchStr {
			result = resultTrue
			break
		} else {
			result = resultFalse
		}
	}

	return result, err
}

func searchSubString(searchStr, fileName string) ([]string, error) {
	content, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}
	defer content.Close()

	fileScanner := bufio.NewScanner(content)
	list := make([]string, 5)
	var i int = 0

	for fileScanner.Scan() {
		if strings.Contains(fileScanner.Text(), searchStr) {
			list[i] = fileScanner.Text()
			i++
		}
	}

	return list, err
}

func addStrToNewFile(newString, fileName string) (string, error) {
	content, err := os.Open(fileName)
	var warning string = fileName + " file already exists"
	if err != nil {
		content, _ := os.Create(fileName)
		content.Write([]byte("You have added the text " + newString))
		fileScanner := bufio.NewScanner(content)
		return fileScanner.Text(), nil
	}
	defer content.Close()
	return warning, err
}
