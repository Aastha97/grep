package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	var output string
	str, option, fileName, inputs, newFileName := takeInputs()
	switch option {
	case 1:
		content, err := readFile(fileName)
		result, _ := searchString(str, content, option)
		if len(result) != 0 {
			output = "I found the " + str + " in the " + fileName + " file"
		} else {
			output = "Cannot find the " + str + " in the " + fileName + " file"
		}
		finalResult(output, err)
	case 2:
		fmt.Println("All inputs done")
		result, _ := searchString(str, strings.Join(inputs, "\r\n"), option)
		finalResult(strings.Join(result, "\n"), nil)
	case 3:
		content, err := readFile(fileName)
		if err != nil {
			finalResult(content, err)
		}
		result, _ := searchString(str, content, option)
		writeFile(result, newFileName)
		content, err = readFile(newFileName)
		finalResult(content, err)

	case 4:
		content, err := readFile(fileName)
		if err != nil {
			finalResult(content, err)
		}
		result, _ := searchString(str, content, option)
		writeFile(result, newFileName)
		content, err = readFile(newFileName)
		finalResult(content, err)
	}
}

func searchString(searchStr, content string, option int) ([]string, error) {
	list := strings.Split(content, "\r\n")
	var finalList []string
	var i int = 0
	for _, str := range list {
		for _, word := range strings.Fields(str) {
			if (option == 1 || option == 3) && word == searchStr {
				finalList = append(finalList, str)
				i++
			} else if option == 2 && strings.Contains(word, searchStr) {
				finalList = append(finalList, str)
				i++
			} else if option == 4 {
				matched, _ := regexp.MatchString("(?i)"+searchStr, word)
				if matched {
					finalList = append(finalList, str)
					i++
				}
			}
		}
	}
	return finalList, nil
}

func readFile(fileName string) (string, error) {
	data, err := ioutil.ReadFile(fileName)
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
		for _, str := range newString {
			content.Write([]byte(str + "\n"))
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

func takeInputs() (string, int, string, []string, string) {
	var str, fileName, newFileName string
	var inputs []string
	fmt.Println("Enter the string : ")
	fmt.Scanln(&str)
	fmt.Println("Choose one:\n 1. Find in a file \n 2. from the terminal \n 3. Write the sting in a file \n 4. Find the string(case-insensitive) in a file then save the output in another file \n 5. Find the string in all the files ")
	var option, inputNo int
	fmt.Scanln(&option)
	if option == 1 || option == 3 || option == 4 {
		fmt.Println("Enter the file Name to search in: ")
		fmt.Scanln(&fileName)
	}
	if option == 2 {
		fmt.Println("Enter the number of inputs")
		fmt.Scanln(&inputNo)
		fmt.Println("Enter string in different line: ")
		inputs = make([]string, inputNo)
		for i := 0; i < inputNo; i++ {
			fmt.Scanln(&inputs[i])
		}
	}
	if option == 3 || option == 4 {
		fmt.Println("Enter the file Name to store the ouput: ")
		fmt.Scanln(&newFileName)
	}
	return str, option, fileName, inputs, newFileName
}
