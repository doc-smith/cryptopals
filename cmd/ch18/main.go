package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/doc-smith/cryptopals/util/crypto/sym"
	"github.com/doc-smith/cryptopals/util/inp"
)

func getSecret() []byte {
	const secret = "L77na/nrFsKvynd6HzOoG7GHTLXsTVu9qvY/2syLXzhPweyyMTJULu/6/kXX0KSvoOLSFQ=="
	return inp.ReadBase64(strings.NewReader(secret))
}

func xor(a, b []byte) []byte {
	c := make([]byte, len(a))
	for i := range a {
		c[i] = a[i] ^ b[i]
	}
	return c
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func encryptAesCtr(pt, key []byte, nonce uint64) []byte {
	var ct bytes.Buffer

	ctr := make([]byte, 16)
	binary.LittleEndian.PutUint64(ctr, nonce)

	for i := 0; i < len(pt); i += 16 {
		binary.LittleEndian.PutUint64(ctr[8:], uint64(i/16))

		keystream := sym.EncryptAesEcb(ctr, key)

		s := i
		e := min(i+16, len(pt))
		ct.Write(xor(pt[s:e], keystream))
	}
	return ct.Bytes()
}

func decryptAesCtr(ct, key []byte, nonce uint64) []byte {
	return encryptAesCtr(ct, key, nonce)
}

func main() {
	secret := getSecret()
	pt := decryptAesCtr(secret, []byte("YELLOW SUBMARINE"), 0)
	fmt.Println(string(pt))
}
