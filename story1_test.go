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
	option      int
	output      []string
	content     string
	err         error
	inputs      []string
}

var commonVar CommonVar

func TestSearchStringFromFileZeroMatch(t *testing.T) {
	commonVar.searchStr = "golang"
	commonVar.fileName = "grepS2.txt"
	commonVar.option = 1
	commonVar.content, commonVar.err = readFile(commonVar.fileName)
	if commonVar.err != nil {
		t.Errorf("Cannot read the file %v, %v", commonVar.fileName, commonVar.err)
	}
	commonVar.output, _ = searchString(commonVar.searchStr, commonVar.content, commonVar.option)
	if strings.Contains(strings.Join(commonVar.output, " "), "golang") {
		t.Errorf("expected no error, but got the string %v in the file %v", commonVar.searchStr, commonVar.fileName)
	}
}

func TestSearchStringFromFileOneMatch(t *testing.T) {
	commonVar.searchStr = "golang"
	commonVar.fileName = "grepS1.txt"
	commonVar.option = 1
	commonVar.content, commonVar.err = readFile(commonVar.fileName)
	if commonVar.err != nil {
		t.Errorf("Cannot read the file %v, %v", commonVar.fileName, commonVar.err)
	}
	commonVar.output, _ = searchString(commonVar.searchStr, commonVar.content, commonVar.option)
	if len(commonVar.output) > 1 {
		t.Errorf("expected no error, but got the string %v in the file %v", commonVar.searchStr, commonVar.fileName)
	}
}

func TestSearchStringFromFileMoreThenOnceMatch(t *testing.T) {
	commonVar.searchStr = "golang"
	commonVar.fileName = "grepS4.txt"
	commonVar.option = 1
	commonVar.content, commonVar.err = readFile(commonVar.fileName)
	if commonVar.err != nil {
		t.Errorf("Cannot read the file %v, %v", commonVar.fileName, commonVar.err)
	}
	commonVar.output, _ = searchString(commonVar.searchStr, commonVar.content, commonVar.option)
	if len(commonVar.output) < 2 {
		t.Errorf("expected no error, but got the string %v in the file %v", commonVar.searchStr, commonVar.fileName)
	}
}
