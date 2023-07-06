package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/doc-smith/cryptopals/util/inp"
)

func readFile(name string) []byte {
	content, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return content
}

func readBase64File(name string) []byte {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	return inp.ReadBase64(f)
}

func TestSolve(t *testing.T) {
	want := readFile("testdata/plaintext.txt")
	got := solve(readBase64File("testdata/ciphertext.txt"))
	if bytes.Compare(got, want) != 0 {
		t.Error("decrypted message is wrong")
	}
}
