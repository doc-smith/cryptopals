package main

import (
	"os"
	"testing"

	"github.com/doc-smith/cryptopals/util/inp"
)

func readTestData(t *testing.T) []string {
	const path = "testdata/ciphertext.txt"
	f, err := os.Open(path)
	if err != nil {
		t.Fatalf("failed to open the test data file (%s): %v", path, err)
	}
	defer f.Close()
	return inp.ReadLines(f)
}

func TestSolve(t *testing.T) {
	lines := readTestData(t)
	want := "Now that the party is jumping\n"
	if got := solve(lines); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
