package main

import (
	"os"
	"testing"

	"github.com/doc-smith/cryptopals/util/inp"
)

func readTestData(t *testing.T) []string {
	path := "testdata/input.txt"
	f, err := os.Open(path)
	if err != nil {
		t.Fatalf("failed to read the test data file (%s): %v", path, err)
	}
	defer f.Close()
	return inp.ReadLines(f)
}

func TestSolve(t *testing.T) {
	const want = 1
	got := solve(readTestData(t))
	if got != want {
		t.Errorf("Solve(...) = %d; want %d", got, want)
	}
}
