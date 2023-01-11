package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Choose one:/n 1. search string /n 2. sub string ")
	var option int
	fmt.Scanln(&option)

	switch option {
	case 1:
		var searchStr, fileName string
		fmt.Scan(&searchStr, &fileName)
		result, err := searchString(searchStr, fileName)

		if err == nil {
			fmt.Println(result)
		} else {
			fmt.Println(err)
		}
	case 2:
		var searchStr, fileName string
		fmt.Scan(&searchStr, &fileName)
		result, err := searchSubString(searchStr, fileName)

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
