package main

import (
	"encoding/hex"
	"fmt"
	"math"
	"os"

	"github.com/doc-smith/cryptopals/util/inp"
	"github.com/doc-smith/cryptopals/util/lang"
)

func decrypt(dst, ct []byte, key byte) {
	for i := range ct {
		dst[i] = ct[i] ^ key
	}
}

func cloneBytes(src []byte) []byte {
	dst := make([]byte, len(src))
	copy(dst, src)
	return dst
}

func findBestPtCandidate(ct []byte) ([]byte, float64) {
	var bestScore float64
	var bestPt []byte

	pt := make([]byte, len(ct))
	for key := 0; key <= math.MaxUint8; key++ {
		decrypt(pt, ct, byte(key))
		score := lang.ScoreEnglishText(pt)
		if score > bestScore {
			bestScore = score
			bestPt = cloneBytes(pt)
		}
	}

	return bestPt, bestScore
}

func solve(lines []string) string {
	var bestScore float64
	var bestPt []byte

	for _, line := range lines {
		ct, err := hex.DecodeString(line)
		if err != nil {
			panic("input is not a valid hex string")
		}
		pt, score := findBestPtCandidate(ct)
		if score > bestScore {
			bestScore = score
			bestPt = pt
		}
	}

	return string(bestPt)
}

func main() {
	inp := inp.ReadLines(os.Stdin)
	fmt.Println(solve(inp))
}
