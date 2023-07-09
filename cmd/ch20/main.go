package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"

	"github.com/doc-smith/cryptopals/util/crypto/rand"
	"github.com/doc-smith/cryptopals/util/crypto/sym"
	"github.com/doc-smith/cryptopals/util/inp"
	"github.com/doc-smith/cryptopals/util/lang"
)

func readPlaintexts(r io.Reader) [][]byte {
	pts := make([][]byte, 0)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		pt := inp.ReadBase64(strings.NewReader(scanner.Text()))
		pts = append(pts, pt)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return pts
}

func encrypt(pts [][]byte) [][]byte {
	key := rand.RandBytes(16)
	cts := make([][]byte, 0, len(pts))
	for _, pt := range pts {
		const nonce = 0
		cts = append(cts, sym.EncryptAesCtr(pt, key, nonce))
	}
	return cts
}

func recoverKeyByte(cts [][]byte, i int) byte {
	var bestScore float64
	var bestCandidate byte

	for candidate := 0; candidate <= math.MaxUint8; candidate++ {
		freq := make(map[byte]float64)
		n := 0
		for _, ct := range cts {
			freq[uint8(candidate)^ct[i]]++
			n++
		}
		for b := range freq {
			freq[b] /= float64(n)
		}
		score := lang.ScoreEnglishTextByteFrequencies(freq)
		if score > bestScore {
			bestScore = score
			bestCandidate = uint8(candidate)
		}
	}

	return bestCandidate
}

func solve(cts [][]byte) {
	if len(cts) < 2 {
		panic("need at least two ciphertexts")
	}

	minLen := len(cts[0])
	for _, ct := range cts {
		if len(ct) < minLen {
			minLen = len(ct)
		}
	}

	keystream := make([]byte, minLen)
	for i := range keystream {
		keystream[i] = recoverKeyByte(cts, i)
	}

	for _, ct := range cts {
		pt := make([]byte, minLen)
		for i := range pt {
			pt[i] = ct[i] ^ keystream[i]
		}
		fmt.Println(string(pt))
	}
}

func main() {
	pts := readPlaintexts(os.Stdin)
	cts := encrypt(pts)
	solve(cts)
}
