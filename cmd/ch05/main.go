package main

import (
	"encoding/hex"
	"fmt"
)

func encrypt(pt, key []byte) []byte {
	ct := make([]byte, len(pt))
	for i := range pt {
		ct[i] = pt[i] ^ key[i%len(key)]
	}
	return ct
}

func solve(pt, key string) string {
	ct := encrypt([]byte(pt), []byte(key))
	return hex.EncodeToString(ct)
}

func main() {
	const plaintext = `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`
	const key = "ICE"
	fmt.Println(solve(plaintext, key))
}
