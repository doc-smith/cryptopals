package main

import (
	"bufio"
	"os"
	"testing"
)

func readTestData(t *testing.T) []string {
	const testDataPath = "testdata/input.txt"
	testDataFile, err := os.Open(testDataPath)
	if err != nil {
		t.Fatalf("failed to open the test data file (%s): %v", testDataPath, err)
	}
	defer func(testDataFile *os.File) {
		_ = testDataFile.Close()
	}(testDataFile)

	scanner := bufio.NewScanner(testDataFile)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		t.Fatalf("failed to read the test data file (%s): %v", testDataPath, err)
	}

	return lines
}

func TestSolve(t *testing.T) {
	lines := readTestData(t)
	want := "Now that the party is jumping\n"
	if got := Solve(lines); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
