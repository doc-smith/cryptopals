package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/doc-smith/cryptopals/util/inp"
)

func solve(inp string) string {
	bs, err := hex.DecodeString(inp)
	if err != nil {
		panic("input is not a valid hex string")
	}
	return base64.StdEncoding.EncodeToString(bs)
}

func main() {
	inp := inp.ReadLine(os.Stdin)
	fmt.Println(solve(inp))
}
