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

type maxByKey[V any] struct {
	MaxKey float64
	Value  V
}

func newMaxByKey[V any]() *maxByKey[V] {
	r := &maxByKey[V]{}
	r.MaxKey = math.Inf(-1)
	return r
}

func (m *maxByKey[V]) Update(key float64, value V) {
	if key > m.MaxKey {
		m.MaxKey = key
		m.Value = value
	}
}

func findBestPlaintextCandidate(ciphertext []byte) ([]byte, float64) {
	bestPlaintext := newMaxByKey[[]byte]()
	for key := 0; key < math.MaxUint8; key++ {
		plaintext := xor.DecryptSingleByteXor(ciphertext, byte(key))
		score := lang.ScoreEnglishText(plaintext)
		bestPlaintext.Update(score, plaintext)
	}
	return bestPlaintext.Value, bestPlaintext.MaxKey
}

func Solve(lines []string) string {
	bestPlaintextCandidate := newMaxByKey[[]byte]()
	for _, line := range lines {
		ciphertext, err := hex.DecodeHexString(line)
		if err != nil {
			panic("input is not a valid hex string")
		}
		plaintextCandidate, candidateScore := findBestPlaintextCandidate(ciphertext)
		bestPlaintextCandidate.Update(candidateScore, plaintextCandidate)
	}
	return string(bestPlaintextCandidate.Value)
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
