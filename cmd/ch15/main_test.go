package main

import (
	"testing"

	"github.com/doc-smith/cryptopals/util/crypto/sym"
)

func TestPadding(t *testing.T) {
	type test struct {
		inp   []byte
		valid bool
	}

	tests := []test{
		{[]byte("ICE ICE BABY\x04\x04\x04\x04"), true},
		{[]byte("ICE ICE BABY\x05\x05\x05\x05"), false},
		{[]byte("ICE ICE BABY\x01\x02\x03\x04"), false},
	}

	for _, tc := range tests {
		_, err := sym.UnpadPKCS7(tc.inp)
		if !tc.valid && err == nil {
			t.Fatalf("UnpadPKCS7 didn't return an error on the input with invalid padding")
		} else if tc.valid && err != nil {
			t.Fatalf("UnpadPKCS7 returned an error on the input with valid padding")
		}
	}
}
