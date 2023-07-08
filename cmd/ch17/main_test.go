package main

import (
	"bytes"
	"testing"

	"github.com/doc-smith/cryptopals/util/crypto/sym"
)

func TestDecipher(t *testing.T) {
	srv := newService()

	tests := []struct {
		name string
		pt   []byte
	}{
		{
			name: "single block",
			pt:   []byte("hello"),
		},
		{
			name: "two blocks",
			pt:   []byte("hello brave new world"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iv, ct := srv.encrypt(tt.pt)
			got, err := sym.UnpadPKCS7(decipher(iv, ct, srv))
			if err != nil {
				t.Fatalf("deciphered is not padded correctly: %v", err)
			}
			want := tt.pt
			if !bytes.Equal(got, want) {
				t.Fatalf("deciphered != pt, padding oracle attack failed")
			}
		})
	}
}
