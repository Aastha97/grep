package main

import (
	"testing"
)

// test function
func TestGrep(t *testing.T) {
	main()
	searchStr := "golang"
	content := "golang"
	option := 1
	actual, err := searchString(searchStr, content, option)
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	if actual != nil {
		t.Errorf("got output %s", actual)
	}
}
