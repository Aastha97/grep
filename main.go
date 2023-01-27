package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var searchStr, fileName, content string
	var err error
	var isCaseInSensitivity = flag.Bool("i", false, "Ignore case when searching")
	var isWordMatch = flag.Bool("w", false, "Word match when searching")
	var outputFile = flag.String("o", "", "File to write the matches")
	flag.Parse()
	if len(os.Args) < 2 {
		fmt.Println("Usage: grep [option(s)...] pattern [file_name]")
		return
	}
	nonFlagValues := flag.Args()
	if strings.Contains(os.Args[len(os.Args)-1], ".txt") {
		searchStr = nonFlagValues[0]
		fileName = nonFlagValues[1]
	} else {
		searchStr = os.Args[len(os.Args)-1]
	}

	if fileName == "" {
		var inputNo int
		fmt.Println("Enter the number of inputs")
		fmt.Scanln(&inputNo)
		fmt.Println("Enter string in different line: ")
		inputs := make([]string, inputNo)
		in := bufio.NewReader(os.Stdin)

		for i := 0; i < inputNo; i++ {
			inputs[i], _ = in.ReadString('\n')
		}
		content = strings.Join(inputs, "\r\n")
	} else {
		content, err = readFile(fileName)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
	}

	output := searchString(searchStr, content, *isCaseInSensitivity, *isWordMatch)
	if *outputFile != "" {
		writeFile(output, *outputFile)
	} else {
		finalResult(strings.Join(output, "\n"), err)
	}

}
