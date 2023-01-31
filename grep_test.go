package main

import (
	"strings"
	"testing"
)

type GrepData struct {
	searchStr           string
	fileName            string
	dirName             string
	newFileName         string
	isCaseInSensitivity bool
	isWordMatch         bool
	recursive           bool
	output              string
}

func TestGrepFileSearch(t *testing.T) {
	testData := []GrepData{
		{"result", "grepS1.txt", "", "", false, false, false, ""},
		//go grep
		{"go", "grepS1.txt", "", "", false, false, false, "golang golang going golang allows Multithreading go"},
		//-i go grepS1.txt
		{"go", "grepS1.txt", "", "", true, false, false, "golang Golang is a wonderful language golang going Love Golang language golang allows Multithreading go"},
		//-i -w go grepS1.txt
		{"go", "grepS1.txt", "", "", true, true, false, "go"},
		//-w go grepS1.txt
		{"go", "grepS1.txt", "", "", false, true, false, "go"},
		//file not found
		{"go", "grepS2.txt", "", "", false, true, false, ""},
		//no arguments passed
		{"", "", "", "", false, false, false, ""},
	}

	for _, val := range testData {
		result, err := readFileLineByLine(val.fileName, val.searchStr, val.isCaseInSensitivity, val.isWordMatch, val.recursive)

		if strings.Join(result, " ") != val.output {
			t.Errorf("got error %v, but got the string %v in the file %v", err, val.searchStr, val.fileName)
		} else {
			t.Logf("correct output with flags -i(%v) -w(%v) -r(%v) : %v", val.isCaseInSensitivity, val.isWordMatch, val.recursive, val.output)
		}
	}
}

func TestGrepForDir(t *testing.T) {
	testData := []GrepData{
		{"result", "", "", "", false, false, false, ""},
		// -r result
		{"result", "", "", "", false, false, true, ""},
		{"result", "", "tests", "", false, false, true, ""},
		// -r result tests
		{"golang", "", "tests", "", false, false, true, "tests\\grepS2.txt: golang is a wonderful language\ntests\\grepS2.txt: Love golang language"},
		//-r -i golang tests
		{"golang", "", "tests", "", true, false, true, "tests\\grepS2.txt: Golang\ntests\\grepS2.txt: golang is a wonderful language\ntests\\grepS2.txt: Golang going\ntests\\grepS2.txt: Love golang language\ntests\\grepS2.txt: Golang allows Multithreading"},
	}

	for _, val := range testData {
		result, err := recursiveCallFromDir(val.dirName, val.searchStr, val.isCaseInSensitivity, val.isWordMatch, val.recursive)

		if strings.Join(result, " ") != val.output {
			t.Errorf("got error %v, but got the string %v in the file %v", err, val.searchStr, val.fileName)
		} else {
			t.Logf("correct output with flags -i(%v) -w(%v) -r(%v) : %v", val.isCaseInSensitivity, val.isWordMatch, val.recursive, val.output)
		}
	}
}

func TestGrepForOutputFile(t *testing.T) {
	testData := []GrepData{
		{"result", "", "tests", "out1.txt", false, false, true, ""},
	}

	for _, val := range testData {
		result, err := writeFile(strings.Fields(val.output), val.fileName)
		warning := "\n" + val.fileName + " file already exists, cannot write in file"
		if result == warning {
			t.Errorf("got error %v, but got the string %v in the file %v", err, val.searchStr, val.fileName)
		} else {
			t.Logf("correct output with flags -i(%v) -w(%v) -r(%v) : %v", val.isCaseInSensitivity, val.isWordMatch, val.recursive, val.output)
		}
	}
}
