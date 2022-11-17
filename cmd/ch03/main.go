package main

import (
	"bufio"
	"fmt"
	"github.com/doc-smith/cryptopals/util/ciphers/xor"
	"github.com/doc-smith/cryptopals/util/encoding/hex"
	"math"
	"os"
)

var englishLetterFrequencies = map[byte]float64{
	'a': 0.08167,
	'b': 0.01492,
	'c': 0.02782,
	'd': 0.04253,
	'e': 0.12702,
	'f': 0.02228,
	'g': 0.02015,
	'h': 0.06094,
	'i': 0.06966,
	'j': 0.00153,
	'k': 0.00772,
	'l': 0.04025,
	'm': 0.02406,
	'n': 0.06749,
	'o': 0.07507,
	'p': 0.01929,
	'q': 0.00095,
	'r': 0.05987,
	's': 0.06327,
	't': 0.09056,
	'u': 0.02758,
	'v': 0.00978,
	'w': 0.02360,
	'x': 0.00150,
	'y': 0.01974,
	'z': 0.00074,
}

func countByteFrequency(bs []byte) map[byte]float64 {
	counts := make(map[byte]float64)
	n := 0
	for _, c := range bs {
		counts[c]++
		n++
	}

	for b := range counts {
		counts[b] /= float64(n)
	}

	return counts
}

func bhattacharyyaCoefficient(f, g map[byte]float64) float64 {
	sum := 0.0
	for b, freq := range f {
		sum += math.Sqrt(freq * g[b])
	}
	return sum
}

func scoreEnglishText(text []byte) float64 {
	return bhattacharyyaCoefficient(
		englishLetterFrequencies,
		countByteFrequency(text))
}

func crackSingleByteXor(ciphertext []byte) (byte, []byte) {
	var bestKey byte
	var bestPlaintext []byte
	var bestScore float64

	for key := 0; key < math.MaxUint8; key++ {
		plaintext := xor.DecryptSingleByteXor(ciphertext, byte(key))
		score := scoreEnglishText(plaintext)
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
