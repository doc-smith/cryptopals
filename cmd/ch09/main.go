package main

import (
	"fmt"
	"github.com/doc-smith/cryptopals/util/crypto/sym"
)

func main() {
	const s = "YELLOW SUBMARINE"
	padded := sym.PadPKCS7([]byte(s), 20)
	fmt.Printf("%#v\n", string(padded))
}
