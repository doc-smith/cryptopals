package main

import (
	"testing"

	"github.com/doc-smith/cryptopals/util/crypto/sym"
)

func TestPadPKCS7(t *testing.T) {
	const want = "YELLOW SUBMARINE\x04\x04\x04\x04"
	got := sym.PadPKCS7([]byte("YELLOW SUBMARINE"), 20)
	if string(got) != want {
		t.Fatalf("padding is incorrect")
	}
}
