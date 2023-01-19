package main

import (
	"testing"
)

// test function
var searchStr, fileName string
var option int
var output []string

func TestSearchStringFromFileZeroMatch(t *testing.T) {

	searchStr = "golang"
	option = 1
	fileName = "grepS2.txt"
	content, err := readFile(fileName)
	if err != nil {
		t.Errorf("Cannot read the file %v, %v", fileName, err)
	}
	output, _ = searchString(searchStr, content, option)
	if len(output) != 0 {
		t.Errorf("expected no error, but got the string %v in the file %v", searchStr, fileName)
	}
}

func TestSearchStringFromFileOneMatch(t *testing.T) {
	searchStr = "golang"
	option = 1
	fileName = "grepS1.txt"
	content, err := readFile(fileName)
	if err != nil {
		t.Errorf("Cannot read the file %v, %v", fileName, err)
	}
	output, _ = searchString(searchStr, content, option)
	if len(output) > 1 {
		t.Errorf("expected no error, but got the string %v in the file %v", searchStr, fileName)
	}
}

func TestSearchStringFromFileMoreThenOnceMatch(t *testing.T) {
	searchStr = "golang"
	option = 1
	fileName = "grepS4.txt"
	content, err := readFile(fileName)
	if err != nil {
		t.Errorf("Cannot read the file %v, %v", fileName, err)
	}
	output, _ = searchString(searchStr, content, option)
	if len(output) < 2 {
		t.Errorf("expected no error, but got the string %v in the file %v", searchStr, fileName)
	}
}
