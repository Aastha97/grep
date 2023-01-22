package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 || len(os.Args) > 4 {
		fmt.Println("Usage: grep [option...] [pattern] [file_name]")
		return
	}

	if len(os.Args) == 3 {
		pattern := os.Args[1]
		fileName := os.Args[2]
		content, err := readFile(fileName)
		result, _ := searchString(pattern, content)
		var output string
		if len(result) != 0 {
			output = "I found the " + pattern + " in the " + fileName + " file"
		} else {
			output = "Cannot find the " + pattern + " in the " + fileName + " file"
		}
		finalResult(output, err)
	}
}
