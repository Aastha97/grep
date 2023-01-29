package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var searchStr, fileName, dirName, line, output string
	var result, files []string
	var err error
	var file *os.File
	var isCaseInSensitivity = flag.Bool("i", false, "Ignore case when searching")
	var isWordMatch = flag.Bool("w", false, "Word match when searching")
	var outputFile = flag.String("o", "", "File to write the matches")
	var recursive = flag.Bool("r", false, "recursive search from directory")
	flag.Parse()
	fmt.Printf("recursive: %v\n", *recursive)
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
			files = traverseDir(dirName)

			for _, fileName = range files {
				file, err = os.Open(fileName)
				if err != nil {
					fmt.Println(err)
				}
				defer file.Close()
				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					line = scanner.Text()
					output = searchString(searchStr, line, *isCaseInSensitivity, *isWordMatch)
					if output != "" {
						result = append(result, fileName+": "+output)
					}
				}
				if err := scanner.Err(); err != nil {
					log.Fatal(err)
				}
			}
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

	} else {
		if *recursive {
			dirName = nonFlagValues[1]
			if err != nil {
				fmt.Println(err)
			}
			files = traverseDir(dirName)

			for _, fileName = range files {
				file, err = os.Open(fileName)
				if err != nil {
					fmt.Println(err)
				}
				defer file.Close()
				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					line = scanner.Text()
					output = searchString(searchStr, line, *isCaseInSensitivity, *isWordMatch)
					if output != "" {
						result = append(result, fileName+": "+output)
					}
				}
				if err := scanner.Err(); err != nil {
					log.Fatal(err)
				}
			}

		} else {
			fileName = nonFlagValues[1]
			file, err = os.Open(fileName)
			if err != nil {
				fmt.Println(err)
			}
			defer file.Close()
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line = scanner.Text()

				output = searchString(searchStr, line, *isCaseInSensitivity, *isWordMatch)

				if output != "" {
					result = append(result, output)
				}

			}

			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}

		}
		if *outputFile != "" {
			writeFile(result, *outputFile)
		} else {
			finalResult(strings.Join(result, "\n"))
		}
	}

}
