package main

import (
	"bytes"
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/doc-smith/cryptopals/util/inp"
)

func hasEqualBlocks(ct []byte, bLen int) bool {
	const bs = aes.BlockSize
	for i := bs; i < len(ct); i += bs {
		for j := 0; j < i; j += bs {
			bi := ct[i : i+bs]
			bj := ct[j : j+bs]
			if bytes.Equal(bi, bj) {
				return true
			}
		}
	}
	return false
}

func solve(inp []string) int {
	equalBlockCnt := 0
	for _, line := range inp {
		ct, err := hex.DecodeString(line)
		if err != nil {
			panic(err)
		}
		if hasEqualBlocks(ct, aes.BlockSize) {
			equalBlockCnt++
		}
	}
	return equalBlockCnt
}

func main() {
	inp := inp.ReadLines(os.Stdin)
	ans := solve(inp)
	fmt.Printf("The number of lines with at least one pair of equal blocks: %d\n", ans)
}
