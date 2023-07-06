package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/doc-smith/cryptopals/util/inp"
)

func xor(dst, x, y []byte) {
	for i := range x {
		dst[i] = x[i] ^ y[i]
	}
}

func solve(x, y string) string {
	if len(x) != len(y) {
		panic("inputs must be the same length")
	}

	xb, err := hex.DecodeString(x)
	if err != nil {
		panic("input is not a valid hex string")
	}

	yb, err := hex.DecodeString(y)
	if err != nil {
		panic("input is not a valid hex string")
	}

	xor(xb, xb, yb)
	return hex.EncodeToString(xb)
}

func readInput() (string, string) {
	lines := inp.ReadLines(os.Stdin)
	if len(lines) != 2 {
		panic("input must be exactly two lines")
	}
	return lines[0], lines[1]
}

func main() {
	x, y := readInput()
	fmt.Println(solve(x, y))
}
