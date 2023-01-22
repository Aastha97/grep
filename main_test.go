package main

import (
	"strings"
	"testing"
)

// test function
type CommonVar struct {
	searchStr   string
	fileName    string
	newFileName string
	output      []string
	content     string
	err         error
	inputs      []string
}

var commonVar CommonVar

// story 1 test cases
func TestSearchStringFromFileZeroMatch(t *testing.T) {
	commonVar.searchStr = "golang"
	commonVar.fileName = "grepS2.txt"
	commonVar.content, commonVar.err = readFile(commonVar.fileName)
	expected := ""
	if commonVar.err != nil {
		t.Errorf("Cannot read the file %v, %v", commonVar.fileName, commonVar.err)
	}
	commonVar.output, _ = searchString(commonVar.searchStr, commonVar.content)
	if expected != strings.Join(commonVar.output, " ") {
		t.Errorf("expected no error, but got the string %v in the file %v", commonVar.searchStr, commonVar.fileName)
	}
}

func TestSearchStringFromFileOneMatch(t *testing.T) {
	commonVar.searchStr = "golang"
	commonVar.fileName = "grepS1.txt"
	commonVar.content, commonVar.err = readFile(commonVar.fileName)
	expected := "golang"
	if commonVar.err != nil {
		t.Errorf("Cannot read the file %v, %v", commonVar.fileName, commonVar.err)
	}
	commonVar.output, _ = searchString(commonVar.searchStr, commonVar.content)
	if expected != strings.Join(commonVar.output, " ") {
		t.Errorf("expected no error, but got the string %v in the file %v not exactly once", commonVar.searchStr, commonVar.fileName)
	}
}

func TestSearchStringFromFileMoreThenOnceMatch(t *testing.T) {
	commonVar.searchStr = "golang"
	commonVar.fileName = "grepS4.txt"
	commonVar.content, commonVar.err = readFile(commonVar.fileName)
	expected := "golang golang allows Multithreading"
	if commonVar.err != nil {
		t.Errorf("Cannot read the file %v, %v", commonVar.fileName, commonVar.err)
	}
	commonVar.output, _ = searchString(commonVar.searchStr, commonVar.content)
	if expected != strings.Join(commonVar.output, " ") {
		t.Errorf("expected no error, but got the string %v in the file %v", commonVar.searchStr, commonVar.fileName)
	}
}

// story 2 test cases
func TestSearchSubstring(t *testing.T) {
	commonVar.searchStr = "go"

	commonVar.inputs = []string{"golang", "java", "going", "golang supports in multithearding"}
	expected := "golang going golang supports in multithearding"
	commonVar.output, _ = searchString(commonVar.searchStr, strings.Join(commonVar.inputs, "\r\n"))
	if expected != strings.Join(commonVar.output, " ") {
		t.Errorf("expected no error, but also printed String: java")
	}
}

func TestSearchSubstringWithCaseSensitivity(t *testing.T) {
	commonVar.searchStr = "go"
	commonVar.inputs = []string{"golang", "Golang is a wonderful language", "going", "golang supports in multithearding"}
	expected := "golang going golang supports in multithearding"
	commonVar.output, _ = searchString(commonVar.searchStr, strings.Join(commonVar.inputs, "\r\n"))
	if expected != strings.Join(commonVar.output, " ") {
		t.Errorf("expected no error, but also printed String: Golang is a wonderful language")
	}
}

// story 3 test cases
func TestSearchStringFromFileAndStoreOutputOneMatch(t *testing.T) {
	commonVar.searchStr = "golang"
	commonVar.fileName = "grepS3.txt"
	commonVar.newFileName = "output.txt"
	expected := "golang\n"
	commonVar.content, commonVar.err = readFile(commonVar.fileName)
	if commonVar.err != nil {
		t.Errorf("Cannot read the file %v, %v", commonVar.fileName, commonVar.err)
	}
	commonVar.output, _ = searchString(commonVar.searchStr, commonVar.content)
	if len(commonVar.output) == 0 {
		t.Errorf("expected no error, but cannot find the string %v in file %v", commonVar.searchStr, commonVar.fileName)
	}
	writeFile(commonVar.output, commonVar.newFileName)
	commonVar.content, commonVar.err = readFile(commonVar.newFileName)

	if expected != commonVar.content {
		t.Errorf("wrong data printed in file: %v", commonVar.fileName)
	}

}

func TestSearchStringFromFileAndStoreOutputMultipleMatch(t *testing.T) {
	commonVar.searchStr = "golang"
	commonVar.fileName = "grepS4.txt"
	commonVar.newFileName = "output1.txt"
	expected := "golang\ngolang allows Multithreading\n"
	commonVar.content, commonVar.err = readFile(commonVar.fileName)
	if commonVar.err != nil {
		t.Errorf("Cannot read the file %v, %v", commonVar.fileName, commonVar.err)
	}
	commonVar.output, _ = searchString(commonVar.searchStr, commonVar.content)
	if len(commonVar.output) == 0 {
		t.Errorf("expected no error, but cannot find the string %v in file %v", commonVar.searchStr, commonVar.fileName)
	}
	writeFile(commonVar.output, commonVar.newFileName)
	commonVar.content, commonVar.err = readFile(commonVar.newFileName)

	if expected != commonVar.content {
		t.Errorf("wrong data printed in file: %v", commonVar.fileName)
	}

}

// story 4 test cases
func TestSearchSubstringAndStoreOutputWithCaseInsensitivity(t *testing.T) {
	commonVar.searchStr = "go"
	commonVar.fileName = "grepS3.txt"
	commonVar.newFileName = "output2.txt"
	expected := "golang\nGolang a wonderful language\ngo\n"
	commonVar.content, commonVar.err = readFile(commonVar.fileName)
	if commonVar.err != nil {
		t.Errorf("Cannot read the file %v, %v", commonVar.fileName, commonVar.err)
	}
	commonVar.output, _ = searchString(commonVar.searchStr, commonVar.content)
	if len(commonVar.output) == 0 {
		t.Errorf("expected no error, but cannot find the string %v in file %v", commonVar.searchStr, commonVar.fileName)
	}
	writeFile(commonVar.output, commonVar.newFileName)
	commonVar.content, commonVar.err = readFile(commonVar.newFileName)

	if expected != commonVar.content {
		t.Errorf("wrong data printed in file: %v", commonVar.fileName)
	}
}
