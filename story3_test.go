package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestSearchStringFromFileAndStoreOutput(t *testing.T) {
	commonVar.searchStr = "golang"
	commonVar.fileName = "grepS3.txt"
	commonVar.newFileName = "output.txt"
	commonVar.option = 3
	commonVar.content, commonVar.err = readFile(commonVar.fileName)
	if commonVar.err != nil {
		t.Errorf("Cannot read the file %v, %v", commonVar.fileName, commonVar.err)
	}
	commonVar.output, _ = searchString(commonVar.searchStr, commonVar.content, commonVar.option)
	if len(commonVar.output) == 0 {
		t.Errorf("expected no error, but cannot find the string %v in file %v", commonVar.searchStr, commonVar.fileName)
	}
	writeFile(commonVar.output, commonVar.newFileName)
	commonVar.content, commonVar.err = readFile(commonVar.newFileName)
	//doubt
	data := strings.Join(commonVar.output, "\n")
	if reflect.DeepEqual(data, commonVar.content) == false {
		t.Errorf("Wrong content printed in new file")
	}
}

// new file already exists
