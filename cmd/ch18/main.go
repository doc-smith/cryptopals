package main

import (
	"fmt"
	"strings"

	"github.com/doc-smith/cryptopals/util/crypto/sym"
	"github.com/doc-smith/cryptopals/util/inp"
)

func getSecret() []byte {
	const secret = "L77na/nrFsKvynd6HzOoG7GHTLXsTVu9qvY/2syLXzhPweyyMTJULu/6/kXX0KSvoOLSFQ=="
	return inp.ReadBase64(strings.NewReader(secret))
}

func main() {
	secret := getSecret()
	pt := sym.DecryptAesCtr(secret, []byte("YELLOW SUBMARINE"), 0)
	fmt.Println(string(pt))
}
