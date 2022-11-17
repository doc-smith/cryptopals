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

func findBestPlaintextCandidate(ciphertext []byte) ([]byte, float64) {
	var bestPlaintext []byte
	var bestScore float64

	for key := 0; key < math.MaxUint8; key++ {
		plaintext := xor.DecryptSingleByteXor(ciphertext, byte(key))

		score := lang.ScoreEnglishText(plaintext)
		if score > bestScore {
			bestScore = score
			bestPlaintext = plaintext
		}
	}

	return bestPlaintext, bestScore
}

func Solve(lines []string) string {
	var bestCandidate []byte
	var bestScore float64

	for _, line := range lines {
		ciphertext, err := hex.DecodeHexString(line)
		if err != nil {
			panic("input is not a valid hex string")
		}

		plaintextCandidate, candidateScore := findBestPlaintextCandidate(ciphertext)
		if candidateScore > bestScore {
			bestScore = candidateScore
			bestCandidate = plaintextCandidate
		}
	}

	return string(bestCandidate)
}

func readInput() []string {
	scanner := bufio.NewScanner(os.Stdin)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return lines
}

func main() {
	lines := readInput()
	result := Solve(lines)
	fmt.Println(result)
}
