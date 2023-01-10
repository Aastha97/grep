package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var searchStr, fileName string
	fmt.Scan(&searchStr, &fileName)
	fmt.Println(searchString(searchStr, fileName))
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
