package main

import (
	"bytes"
	"os"
	"testing"
)

const (
	inputFile  = "./testdata/test.md"
	resultFile = "test.md.html"
	goldenFile = "./testdata/test.md.html"
)

func TestParseContent(t *testing.T) {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatal(err)
	}
	result := parseContent(input)

	expected, err := os.ReadFile(goldenFile)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(result, expected) {
		t.Logf("golden:\n%s\n", expected)
		t.Logf("result:\n%s\n", result)
		t.Error("Result content does not match golden file")
	}
	os.Remove(resultFile)
}
