package main

import (
	"crypto/aes"
	"fmt"
	"os"

	"github.com/doc-smith/cryptopals/util/crypto/sym"
	"github.com/doc-smith/cryptopals/util/inp"
)

func solve(ct []byte) []byte {
	const key = "YELLOW SUBMARINE"
	iv := make([]byte, aes.BlockSize)
	pt := sym.DecryptAesCbc(ct, []byte(key), iv)
	unpadded, err := sym.UnpadPKCS7(pt)
	if err != nil {
		panic(err)
	}
	return unpadded

}

func main() {
	inp := inp.ReadBase64(os.Stdin)
	ans := solve(inp)
	fmt.Println(string(ans))
}
