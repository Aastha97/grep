package main

import (
	"strings"
	"testing"
)

func TestSearchSubstring(t *testing.T) {
	commonVar.searchStr = "go"
	commonVar.option = 2
	commonVar.inputs = []string{"golang", "Golang is a wonderful language", "java", "going", "golang supports in multithearding"}
	commonVar.output, _ = searchString(commonVar.searchStr, strings.Join(commonVar.inputs, "\r\n"), commonVar.option)
	if strings.Contains(strings.Join(commonVar.output, " "), "java") {
		t.Errorf("expected no error, but also printed String: java")
	}
}

func TestSearchSubstringWithCaseSensitivity(t *testing.T) {
	commonVar.searchStr = "go"
	commonVar.option = 2
	commonVar.inputs = []string{"golang", "Golang is a wonderful language", "going", "golang supports in multithearding"}
	commonVar.output, _ = searchString(commonVar.searchStr, strings.Join(commonVar.inputs, "\r\n"), commonVar.option)
	if strings.Contains(strings.Join(commonVar.output, " "), "Golang is a wonderful language") {
		t.Errorf("expected no error, but also printed String: Golang is a wonderful language")
	}
}
