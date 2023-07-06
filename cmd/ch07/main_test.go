package main

import (
	"os"
	"strings"
	"testing"

	"github.com/doc-smith/cryptopals/util/inp"
)

func readTestData(t *testing.T) []byte {
	path := "testdata/ciphertext.txt"
	f, err := os.Open(path)
	if err != nil {
		t.Fatalf("failed to read the test data file (%s): %v", path, err)
	}
	defer f.Close()
	return inp.ReadBase64(f)
}

func TestSolve(t *testing.T) {
	ct := readTestData(t)
	pt := string(solve(ct))
	if !strings.HasPrefix(pt, "I'm back and I'm ringin' the bell") {
		t.Errorf("decoded plaintext doesn't look right")
	}
}
