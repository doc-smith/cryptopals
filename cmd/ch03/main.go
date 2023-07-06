package main

import (
	"encoding/hex"
	"fmt"
	"math"
	"os"

	"github.com/doc-smith/cryptopals/util/inp"
	"github.com/doc-smith/cryptopals/util/lang"
)

func decrypt(dst, src []byte, key byte) {
	for i := range src {
		dst[i] = src[i] ^ key
	}
}

func cloneBytes(src []byte) []byte {
	dst := make([]byte, len(src))
	copy(dst, src)
	return dst
}

func crackSingleByteXor(ct []byte) []byte {
	var bestPt []byte
	var bestScore float64

	pt := make([]byte, len(ct))
	for key := 0; key <= math.MaxUint8; key++ {
		decrypt(pt, ct, byte(key))
		score := lang.ScoreEnglishText(pt)
		if score > bestScore {
			bestScore = score
			bestPt = cloneBytes(pt)
		}
	}

	return bestPt
}

func solve(hexCiphertext string) string {
	ct, err := hex.DecodeString(hexCiphertext)
	if err != nil {
		panic("input is not a valid hex string")
	}
	return string(crackSingleByteXor(ct))
}

func main() {
	inp := inp.ReadLine(os.Stdin)
	fmt.Println(solve(inp))
}
