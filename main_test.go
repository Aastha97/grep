package main

import (
	"strings"
	"testing"
)

// test function
type CommonVar struct {
	searchStr           string
	fileName            string
	newFileName         string
	output              []string
	content             string
	err                 error
	inputs              []string
	caseSensitivityflag bool
	wordMatchFlag       bool
}

var commonVar CommonVar

// story 1 test cases
func TestSearchStringFromFileZeroMatch(t *testing.T) {
	commonVar.searchStr = "react"
	commonVar.fileName = "grepS1.txt"
	commonVar.caseSensitivityflag = true
	commonVar.wordMatchFlag = false
	commonVar.content, commonVar.err = readFile(commonVar.fileName)
	expected := ""
	if commonVar.err != nil {
		t.Errorf("Cannot read the file %v, %v", commonVar.fileName, commonVar.err)
	}
	commonVar.output = searchString(commonVar.searchStr, commonVar.content, commonVar.caseSensitivityflag, commonVar.wordMatchFlag)
	if expected != strings.Join(commonVar.output, " ") {
		t.Errorf("expected no error, but got the string %v in the file %v", commonVar.searchStr, commonVar.fileName)
	}
}

func TestSearchStringFromFileOneMatch(t *testing.T) {
	commonVar.searchStr = "Java"
	commonVar.fileName = "grepS1.txt"
	commonVar.caseSensitivityflag = true
	commonVar.wordMatchFlag = false
	commonVar.content, commonVar.err = readFile(commonVar.fileName)
	expected := "Java"
	if commonVar.err != nil {
		t.Errorf("Cannot read the file %v, %v", commonVar.fileName, commonVar.err)
	}
	commonVar.output = searchString(commonVar.searchStr, commonVar.content, commonVar.caseSensitivityflag, commonVar.wordMatchFlag)
	if expected != strings.Join(commonVar.output, " ") {
		t.Errorf("expected no error, but did not got the string %v in the file %v not exactly once", commonVar.searchStr, commonVar.fileName)
	}
}

func TestSearchStringFromFileMoreThenOnceMatch(t *testing.T) {
	commonVar.searchStr = "golang"
	commonVar.fileName = "grepS1.txt"
	commonVar.caseSensitivityflag = true
	commonVar.wordMatchFlag = false
	commonVar.content, commonVar.err = readFile(commonVar.fileName)
	expected := "golang golang going golang allows Multithreading "
	if commonVar.err != nil {
		t.Errorf("Cannot read the file %v, %v", commonVar.fileName, commonVar.err)
	}
	commonVar.output = searchString(commonVar.searchStr, commonVar.content, commonVar.caseSensitivityflag, commonVar.wordMatchFlag)
	if expected != strings.Join(commonVar.output, " ") {
		t.Errorf("expected no error, but did not got the string %v in the file %v", commonVar.searchStr, commonVar.fileName)
	}
}

func TestSearchStringFromFileWithCaseInsensitivity(t *testing.T) {
	commonVar.searchStr = "Go"
	commonVar.fileName = "grepS1.txt"
	commonVar.caseSensitivityflag = false
	commonVar.wordMatchFlag = false
	commonVar.content, commonVar.err = readFile(commonVar.fileName)
	expected := "golang Golang is a wonderful language golang going Love Golang language golang allows Multithreading go"
	if commonVar.err != nil {
		t.Errorf("Cannot read the file %v, %v", commonVar.fileName, commonVar.err)
	}
	commonVar.output = searchString(commonVar.searchStr, commonVar.content, commonVar.caseSensitivityflag, commonVar.wordMatchFlag)
	if expected != strings.Join(commonVar.output, " ") {
		t.Errorf("expected no error, but did not got the string %v in the file %v", commonVar.searchStr, commonVar.fileName)
	}
}
func TestSearchStringFromFileWithCaseInsensitivityAndWordMatch(t *testing.T) {
	commonVar.searchStr = "Golang"
	commonVar.fileName = "grepS1.txt"
	commonVar.caseSensitivityflag = false
	commonVar.wordMatchFlag = true
	commonVar.content, commonVar.err = readFile(commonVar.fileName)
	expected := "golang Golang is a wonderful language golang going Love Golang language golang allows Multithreading"
	if commonVar.err != nil {
		t.Errorf("Cannot read the file %v, %v", commonVar.fileName, commonVar.err)
	}
	commonVar.output = searchString(commonVar.searchStr, commonVar.content, commonVar.caseSensitivityflag, commonVar.wordMatchFlag)
	if expected != strings.Join(commonVar.output, " ") {
		t.Errorf("expected no error, but got the string %v in the file %v", commonVar.searchStr, commonVar.fileName)
	}
}

func TestSearchStringFromFileWithWordMatch(t *testing.T) {
	commonVar.searchStr = "go"
	commonVar.fileName = "grepS1.txt"
	commonVar.caseSensitivityflag = false
	commonVar.wordMatchFlag = true
	commonVar.content, commonVar.err = readFile(commonVar.fileName)
	expected := "go"
	if commonVar.err != nil {
		t.Errorf("Cannot read the file %v, %v", commonVar.fileName, commonVar.err)
	}
	commonVar.output = searchString(commonVar.searchStr, commonVar.content, commonVar.caseSensitivityflag, commonVar.wordMatchFlag)
	if expected != strings.Join(commonVar.output, " ") {
		t.Errorf("expected no error, but got the string %v in the file %v", commonVar.searchStr, commonVar.fileName)
	}
}

// story 2 test cases
func TestSearchSubstring(t *testing.T) {
	commonVar.searchStr = "go"
	commonVar.caseSensitivityflag = false
	commonVar.wordMatchFlag = false
	commonVar.inputs = []string{"golang", "java", "Going", "golang supports in multithearding"}
	expected := "golang golang supports in multithearding"
	commonVar.output = searchString(commonVar.searchStr, commonVar.content, commonVar.caseSensitivityflag, commonVar.wordMatchFlag)
	if expected != strings.Join(commonVar.output, " ") {
		t.Errorf("expected no error, but also printed String: Going")
	}
}

func TestSearchSubstringWithCaseSensitivity(t *testing.T) {
	commonVar.searchStr = "go"
	commonVar.caseSensitivityflag = true
	commonVar.wordMatchFlag = false
	commonVar.inputs = []string{"golang", "Golang is a wonderful language", "going", "golang supports in multithearding"}
	expected := "golang going golang supports in multithearding"
	commonVar.output = searchString(commonVar.searchStr, commonVar.content, commonVar.caseSensitivityflag, commonVar.wordMatchFlag)
	if expected != strings.Join(commonVar.output, " ") {
		t.Errorf("expected no error, but also printed String: Golang is a wonderful language")
	}
}

// story 3 test cases
func TestSearchStringFromFileAndStoreOutputOneMatch(t *testing.T) {
	commonVar.searchStr = "Java"
	commonVar.fileName = "grepS1.txt"
	commonVar.caseSensitivityflag = true
	commonVar.wordMatchFlag = false
	commonVar.content, commonVar.err = readFile(commonVar.fileName)
	expected := "Java"
	if commonVar.err != nil {
		t.Errorf("Cannot read the file %v, %v", commonVar.fileName, commonVar.err)
	}
	commonVar.output = searchString(commonVar.searchStr, commonVar.content, commonVar.caseSensitivityflag, commonVar.wordMatchFlag)
	if expected != strings.Join(commonVar.output, " ") {
		t.Errorf("expected no error, but did not got the string %v in the file %v not exactly once", commonVar.searchStr, commonVar.fileName)
	}
	writeFile(commonVar.output, commonVar.newFileName)
	commonVar.content, commonVar.err = readFile(commonVar.newFileName)

	if expected != commonVar.content {
		t.Errorf("wrong data printed in file: %v", commonVar.fileName)
	}

}

func TestSearchStringFromFileAndStoreOutputMultipleMatch(t *testing.T) {
	commonVar.searchStr = "golang"
	commonVar.fileName = "grepS1.txt"
	commonVar.caseSensitivityflag = true
	commonVar.wordMatchFlag = false
	commonVar.content, commonVar.err = readFile(commonVar.fileName)
	expected := "golang golang going golang allows Multithreading "
	if commonVar.err != nil {
		t.Errorf("Cannot read the file %v, %v", commonVar.fileName, commonVar.err)
	}
	commonVar.output = searchString(commonVar.searchStr, commonVar.content, commonVar.caseSensitivityflag, commonVar.wordMatchFlag)
	if expected != strings.Join(commonVar.output, " ") {
		t.Errorf("expected no error, but did not got the string %v in the file %v", commonVar.searchStr, commonVar.fileName)
	}
	writeFile(commonVar.output, commonVar.newFileName)
	commonVar.content, commonVar.err = readFile(commonVar.newFileName)

	if expected != commonVar.content {
		t.Errorf("wrong data printed in file: %v", commonVar.fileName)
	}

}

// story 4 test cases
func TestSearchSubstringAndStoreOutputWithCaseInsensitivity(t *testing.T) {
	commonVar.searchStr = "Go"
	commonVar.fileName = "grepS1.txt"
	commonVar.caseSensitivityflag = false
	commonVar.wordMatchFlag = false
	commonVar.content, commonVar.err = readFile(commonVar.fileName)
	expected := "golang Golang is a wonderful language golang going Love Golang language golang allows Multithreading go"
	if commonVar.err != nil {
		t.Errorf("Cannot read the file %v, %v", commonVar.fileName, commonVar.err)
	}
	commonVar.output = searchString(commonVar.searchStr, commonVar.content, commonVar.caseSensitivityflag, commonVar.wordMatchFlag)
	if expected != strings.Join(commonVar.output, " ") {
		t.Errorf("expected no error, but did not got the string %v in the file %v", commonVar.searchStr, commonVar.fileName)
	}
	writeFile(commonVar.output, commonVar.newFileName)
	commonVar.content, commonVar.err = readFile(commonVar.newFileName)

	if expected != commonVar.content {
		t.Errorf("wrong data printed in file: %v", commonVar.fileName)
	}
}
