package main

import (
	"bytes"
	"math"

	"github.com/doc-smith/cryptopals/util/crypto/rand"
	"github.com/doc-smith/cryptopals/util/crypto/sym"
)

func getCookie(userdata, key []byte, nonce uint64) []byte {
	var b bytes.Buffer
	b.Write([]byte("comment1=cooking%20MCs;userdata="))
	b.Write(userdata)
	b.Write([]byte(";comment2=%20like%20a%20pound%20of%20bacon"))
	pt := b.Bytes()
	return sym.EncryptAesCtr(pt, key, nonce)
}

func isAdmin(ct, key []byte, nonce uint64) bool {
	pt := sym.DecryptAesCtr(ct, key, nonce)
	for _, pair := range bytes.Split(pt, []byte(";")) {
		k, v, found := bytes.Cut(pair, []byte("="))
		if !found {
			panic("invalud key value pair")
		}
		if bytes.Equal(k, []byte("admin")) && bytes.Equal(v, []byte("true")) {
			return true
		}
	}
	return false
}

func ctrBitFlipping() {
	// 0123456789abcdef
	// comment1=cooking
	// %20MCs;userdata=
	// x;comment2=%20li
	// ke%20a%20pound%2
	// 0of%20bacon
	// ;admin=true

	key := rand.RandBytes(16)
	nonce := rand.RandUint64(math.MaxUint64)

	cookie := getCookie([]byte("x"), key, nonce)

	target := []byte(";admin=true")
	wtf := []byte("0of%20bacon")
	for i := range target {
		b := wtf[i] ^ target[i]
		cookie[len(cookie)-len(target)+i] ^= b
	}

	if !isAdmin(cookie, key, nonce) {
		panic("CTR bitflipping attack failed")
	}
}

func main() {
	ctrBitFlipping()
}
