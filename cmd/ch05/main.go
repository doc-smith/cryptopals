package main

import (
	"encoding/hex"
	"fmt"
)

func encrypt(plaintext, key []byte) []byte {
	ciphertext := make([]byte, len(plaintext))
	for i, b := range plaintext {
		ciphertext[i] = b ^ key[i%len(key)]
	}
	return ciphertext
}

func Solve(plaintext, key string) string {
	return hex.EncodeToString(encrypt([]byte(plaintext), []byte(key)))
}

func main() {
	const plaintext = `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`
	const key = "ICE"
	fmt.Println(Solve(plaintext, key))
}
