package main

import (
	"bytes"
	"crypto/aes"

	"github.com/doc-smith/cryptopals/util/crypto/rand"
	"github.com/doc-smith/cryptopals/util/crypto/sym"
)

type mode = uint8

const (
	ECB mode = iota
	CBC
)

func addGarbage(pt []byte) []byte {
	pfx := rand.RandBytes(rand.RandInt(10))
	sfx := rand.RandBytes(rand.RandInt(10))
	return append(append(pfx, pt...), sfx...)
}

func encryptionOracle(pt []byte) ([]byte, mode) {
	key := rand.RandBytes(aes.BlockSize)
	padded := sym.PadPKCS7(addGarbage(pt), aes.BlockSize)

	if rand.CoinFlip() == rand.Head {
		ct := sym.EncryptAesEcb(padded, key)
		return ct, ECB
	}

	iv := rand.RandBytes(aes.BlockSize)
	ct := sym.EncryptAesCbc(padded, key, iv)
	return ct, CBC
}

func hasTwoEqualBlocks(ct []byte, blen int) bool {
	for i := 0; i+2*blen <= len(ct); i += blen {
		f := ct[i : i+blen]
		s := ct[i+blen : i+2*blen]
		if bytes.Equal(f, s) {
			return true
		}
	}
	return false
}

func detectionOracle(ct []byte) mode {
	if hasTwoEqualBlocks(ct, aes.BlockSize) {
		return ECB
	}
	return CBC
}

func testDetectionOracle(iterCnt int) {
	for i := 0; i < iterCnt; i++ {
		inp := make([]byte, 3*aes.BlockSize)
		ct, mode := encryptionOracle(inp)
		guess := detectionOracle(ct)
		if guess != mode {
			panic("detection oracle failed")
		}
	}
}

func main() {
	const iterCnt = 100
	testDetectionOracle(iterCnt)
}
