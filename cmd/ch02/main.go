package main

import (
	"bufio"
	"fmt"
	"github.com/doc-smith/cryptopals/util/encoding/hex"
	"os"
)

func xor(xs []byte, ys []byte) []byte {
	var res []byte
	for i := 0; i < len(xs); i++ {
		res = append(res, xs[i]^ys[i])
	}
	return res
}

func Solve(x string, y string) string {
	if len(x) != len(y) {
		panic("inputs must be the same length")
	}

	xb, err := hex.DecodeHexString(x)
	if err != nil {
		panic("input is not a valid hex string")
	}

	yb, err := hex.DecodeHexString(y)
	if err != nil {
		panic("input is not a valid hex string")
	}

	return hex.EncodeHexString(xor(xb, yb))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		x := scanner.Text()
		if scanner.Scan() {
			y := scanner.Text()
			fmt.Println(Solve(x, y))
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
