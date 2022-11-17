package main

import (
	"bufio"
	"fmt"
	"github.com/doc-smith/cryptopals/util/ciphers/xor"
	"github.com/doc-smith/cryptopals/util/encoding/hex"
	"github.com/doc-smith/cryptopals/util/lang"
	"math"
	"os"
)

func crackSingleByteXor(ciphertext []byte) (byte, []byte) {
	var bestKey byte
	var bestPlaintext []byte
	var bestScore float64

	for key := 0; key < math.MaxUint8; key++ {
		plaintext := xor.DecryptSingleByteXor(ciphertext, byte(key))
		score := lang.ScoreEnglishText(plaintext)
		if score > bestScore {
			bestScore = score
			bestKey = byte(key)
			bestPlaintext = plaintext
		}
	}

	return bestKey, bestPlaintext
}

func Solve(hexCiphertext string) string {
	cipertext, err := hex.DecodeHexString(hexCiphertext)
	if err != nil {
		panic("input is not a valid hex string")
	}
	_, plaintext := crackSingleByteXor(cipertext)
	return string(plaintext)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		fmt.Println(Solve(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
