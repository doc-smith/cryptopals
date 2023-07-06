package main

import (
	"fmt"
	"os"

	"github.com/doc-smith/cryptopals/util/crypto/sym"
	"github.com/doc-smith/cryptopals/util/inp"
)

func solve(ct []byte) []byte {
	const key = "YELLOW SUBMARINE"
	return sym.DecryptAesEcb(ct, []byte(key))
}

func main() {
	inp := inp.ReadBase64(os.Stdin)
	ans := solve(inp)
	fmt.Println(string(ans))
}
