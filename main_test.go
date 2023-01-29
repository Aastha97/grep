package main

import (
	"strings"
	"testing"
)

type CommonVar struct {
	searchStr             string
	fileName              string
	output                []string
	err                   error
	caseInSensitivityflag bool
	wordMatchFlag         bool
	recursive             bool
}

var commonVar CommonVar

func TestSearchStringFromFileZeroMatch(t *testing.T) {
	commonVar.searchStr = "react"
	commonVar.fileName = "grepS1.txt"
	commonVar.caseInSensitivityflag = false
	commonVar.wordMatchFlag = false
	commonVar.recursive = false
	expected := ""
	if commonVar.err != nil {
		t.Errorf("Cannot read the file %v, %v", commonVar.fileName, commonVar.err)
	}
	commonVar.output = readFileLineByLine(commonVar.fileName, commonVar.searchStr, commonVar.caseInSensitivityflag, commonVar.wordMatchFlag, commonVar.recursive)
	if expected != strings.Join(commonVar.output, " ") {
		t.Errorf("expected no error, but got the string %v in the file %v", commonVar.searchStr, commonVar.fileName)
	}
}

// func TestSearchStringFromFileOneMatch(t *testing.T) {
// 	commonVar.searchStr = "Java"
// 	commonVar.fileName = "grepS1.txt"
// 	commonVar.caseInSensitivityflag = false
// 	commonVar.wordMatchFlag = false
// 	commonVar.content, commonVar.err = readFile(commonVar.fileName)
// 	expected := "Java"
// 	if commonVar.err != nil {
// 		t.Errorf("Cannot read the file %v, %v", commonVar.fileName, commonVar.err)
// 	}
// 	commonVar.output = searchString(commonVar.searchStr, commonVar.content, commonVar.caseInSensitivityflag, commonVar.wordMatchFlag)
// 	if expected != strings.Join(commonVar.output, " ") {
// 		t.Errorf("expected no error, but did not got the string %v in the file %v not exactly once", commonVar.searchStr, commonVar.fileName)
// 	}
// }

// func TestSearchStringFromFileMoreThenOnceMatch(t *testing.T) {
// 	commonVar.searchStr = "golang"
// 	commonVar.fileName = "grepS1.txt"
// 	commonVar.caseInSensitivityflag = false
// 	commonVar.wordMatchFlag = false
// 	commonVar.content, commonVar.err = readFile(commonVar.fileName)
// 	expected := "golang golang going golang allows Multithreading"
// 	if commonVar.err != nil {
// 		t.Errorf("Cannot read the file %v, %v", commonVar.fileName, commonVar.err)
// 	}
// 	commonVar.output = searchString(commonVar.searchStr, commonVar.content, commonVar.caseInSensitivityflag, commonVar.wordMatchFlag)
// 	if expected != strings.Join(commonVar.output, " ") {
// 		t.Errorf("expected no error, but did not got the string %v in the file %v", commonVar.searchStr, commonVar.fileName)
// 	}
// }

// func TestSearchStringFromFileWithCaseInsensitivity(t *testing.T) {
// 	commonVar.searchStr = "go"
// 	commonVar.fileName = "grepS1.txt"
// 	commonVar.caseInSensitivityflag = true
// 	commonVar.wordMatchFlag = false
// 	commonVar.content, commonVar.err = readFile(commonVar.fileName)
// 	expected := "golang Golang is a wonderful language golang going Love Golang language golang allows Multithreading go"
// 	if commonVar.err != nil {
// 		t.Errorf("Cannot read the file %v, %v", commonVar.fileName, commonVar.err)
// 	}
// 	commonVar.output = searchString(commonVar.searchStr, commonVar.content, commonVar.caseInSensitivityflag, commonVar.wordMatchFlag)
// 	if expected != strings.Join(commonVar.output, " ") {
// 		t.Errorf("expected no error, but did not got the string %v in the file %v", commonVar.searchStr, commonVar.fileName)
// 	}
// }
// func TestSearchStringFromFileWithCaseInsensitivityAndWordMatch(t *testing.T) {
// 	commonVar.searchStr = "Golang"
// 	commonVar.fileName = "grepS1.txt"
// 	commonVar.caseInSensitivityflag = true
// 	commonVar.wordMatchFlag = true
// 	commonVar.content, commonVar.err = readFile(commonVar.fileName)
// 	expected := "golang Golang is a wonderful language golang going Love Golang language golang allows Multithreading"
// 	if commonVar.err != nil {
// 		t.Errorf("Cannot read the file %v, %v", commonVar.fileName, commonVar.err)
// 	}
// 	commonVar.output = searchString(commonVar.searchStr, commonVar.content, commonVar.caseInSensitivityflag, commonVar.wordMatchFlag)
// 	if expected != strings.Join(commonVar.output, " ") {
// 		t.Errorf("expected no error, but got the string %v in the file %v", commonVar.searchStr, commonVar.fileName)
// 	}
// }

// func TestSearchStringFromFileWithWordMatch(t *testing.T) {
// 	commonVar.searchStr = "go"
// 	commonVar.fileName = "grepS1.txt"
// 	commonVar.caseInSensitivityflag = false
// 	commonVar.wordMatchFlag = true
// 	commonVar.content, commonVar.err = readFile(commonVar.fileName)
// 	expected := "go"
// 	if commonVar.err != nil {
// 		t.Errorf("Cannot read the file %v, %v", commonVar.fileName, commonVar.err)
// 	}
// 	commonVar.output = searchString(commonVar.searchStr, commonVar.content, commonVar.caseInSensitivityflag, commonVar.wordMatchFlag)
// 	if expected != strings.Join(commonVar.output, " ") {
// 		t.Errorf("expected no error, but got the string %v in the file %v", commonVar.searchStr, commonVar.fileName)
// 	}
// }
