package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var searchStr, fileName, dirName, line, output string
	var result []string
	var err error
	var isCaseInSensitivity = flag.Bool("i", false, "Ignore case when searching")
	var isWordMatch = flag.Bool("w", false, "Word match when searching")
	var outputFile = flag.String("o", "", "File to write the matches")
	var recursive = flag.Bool("r", false, "recursive search from directory")
	flag.Parse()
	if len(os.Args) < 2 {
		fmt.Println("Usage: grep [option(s)...] pattern [file_name]")
		return
	}
	nonFlagValues := flag.Args()
	searchStr = nonFlagValues[0]
	if len(nonFlagValues) == 1 {
		if *recursive {
			dirName, err = os.Getwd()
			if err != nil {
				fmt.Println(err)
			}
			result, _ = recursiveCallFromDir(dirName, searchStr, *isCaseInSensitivity, *isWordMatch, *recursive)
		} else {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				line = scanner.Text()
				output = searchString(searchStr, line, *isCaseInSensitivity, *isWordMatch)
				if output != "" {
					fmt.Println(output + "\n")
				}
			}
		}
		if *outputFile != "" {
			writeFile(result, *outputFile)
		}
	} else {
		if *recursive {
			dirName = nonFlagValues[1]
			if err != nil {
				fmt.Println(err)
			}
			result, _ = recursiveCallFromDir(dirName, searchStr, *isCaseInSensitivity, *isWordMatch, *recursive)
		} else {
			fileName = nonFlagValues[1]
			result, err = readFileLineByLine(fileName, searchStr, *isCaseInSensitivity, *isWordMatch, *recursive)
			finalResult(strings.Join(result, "\n"), err)
		}
		if *outputFile != "" {
			writeFile(result, *outputFile)
		}
	}

}
