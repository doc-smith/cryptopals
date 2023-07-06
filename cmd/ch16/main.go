package main

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"

	"github.com/doc-smith/cryptopals/util/crypto/sym"
)

func randBytes(n int) []byte {
	res := make([]byte, n)
	if _, err := rand.Read(res); err != nil {
		panic(err)
	}
	return res
}

func getCookie(userdata, key, iv []byte) []byte {
	var b bytes.Buffer
	b.Write([]byte("comment1=cooking%20MCs;userdata="))
	b.Write(userdata)
	b.Write([]byte(";comment2=%20like%20a%20pound%20of%20bacon"))
	pt := b.Bytes()
	return sym.EncryptAesCbc(
		sym.PadPKCS7(pt, aes.BlockSize),
		key,
		iv)
}

func isAdmin(ct, key, iv []byte) bool {
	pt, err := sym.UnpadPKCS7(sym.DecryptAesCbc(ct, key, iv))
	if err != nil {
		panic(err)
	}
	for _, pair := range bytes.Split(pt, []byte(";")) {
		k, v, found := bytes.Cut(pair, []byte("="))
		if !found {
			panic("invalud key value pair")
		}
		if bytes.Compare(k, []byte("admin")) == 0 && bytes.Compare(v, []byte("true")) == 0 {
			return true
		}
	}
	return false
}

func cbcBitFlipping() {
	// 0123456789abcdef
	// comment1=cooking
	// %20MCs;userdata=
	// x;comment2=%20li
	// ke%20a%20pound%2
	// 0of%20bacon
	// ;admin=true

	key := randBytes(16)
	iv := randBytes(aes.BlockSize)

	cookie := getCookie([]byte("x"), key, iv)

	target := []byte(";admin=true")
	wtf := []byte("0of%20bacon")
	for i := range target {
		b := wtf[i] ^ target[i]
		cookie[len(cookie)-2*aes.BlockSize+i] ^= b
	}

	if !isAdmin(cookie, key, iv) {
		panic("CBC bitflipping ttack failed")
	}
}

func main() {
	cbcBitFlipping()
}
