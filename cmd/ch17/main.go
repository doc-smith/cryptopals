package main

import (
	"bytes"
	"crypto/aes"
	"fmt"
	"math"
	"strings"

	"github.com/doc-smith/cryptopals/util/crypto/rand"
	"github.com/doc-smith/cryptopals/util/crypto/sym"
	"github.com/doc-smith/cryptopals/util/inp"
)

const blockSize = aes.BlockSize

func getRandPt() []byte {
	pts := []string{
		"MDAwMDAwTm93IHRoYXQgdGhlIHBhcnR5IGlzIGp1bXBpbmc=",
		"MDAwMDAxV2l0aCB0aGUgYmFzcyBraWNrZWQgaW4gYW5kIHRoZSBWZWdhJ3MgYXJlIHB1bXBpbic=",
		"MDAwMDAyUXVpY2sgdG8gdGhlIHBvaW50LCB0byB0aGUgcG9pbnQsIG5vIGZha2luZw==",
		"MDAwMDAzQ29va2luZyBNQydzIGxpa2UgYSBwb3VuZCBvZiBiYWNvbg==",
		"MDAwMDA0QnVybmluZyAnZW0sIGlmIHlvdSBhaW4ndCBxdWljayBhbmQgbmltYmxl",
		"MDAwMDA1SSBnbyBjcmF6eSB3aGVuIEkgaGVhciBhIGN5bWJhbA==",
		"MDAwMDA2QW5kIGEgaGlnaCBoYXQgd2l0aCBhIHNvdXBlZCB1cCB0ZW1wbw==",
		"MDAwMDA3SSdtIG9uIGEgcm9sbCwgaXQncyB0aW1lIHRvIGdvIHNvbG8=",
		"MDAwMDA4b2xsaW4nIGluIG15IGZpdmUgcG9pbnQgb2g=",
		"MDAwMDA5aXRoIG15IHJhZy10b3AgZG93biBzbyBteSBoYWlyIGNhbiBibG93",
	}
	pt := pts[rand.RandInt(len(pts))]
	return inp.ReadBase64(strings.NewReader(pt))
}

type service struct {
	key []byte
}

func newService() *service {
	return &service{
		rand.RandBytes(16),
	}
}

func (s *service) encrypt(pt []byte) ([]byte, []byte) {
	iv := rand.RandBytes(blockSize)
	return iv, sym.EncryptAesCbc(sym.PadPKCS7(pt, blockSize), s.key, iv)
}

func (s *service) hasValidPadding(iv, ct []byte) bool {
	padded := sym.DecryptAesCbc(ct, s.key, iv)
	_, err := sym.UnpadPKCS7(padded)
	return err == nil
}

func decipherSingleBlock(iv, block []byte, s *service) []byte {
	pt := make([]byte, blockSize)
	for pad := 1; pad <= blockSize; pad++ {
		iv := make([]byte, blockSize)
		for i := 1; i <= pad; i++ {
			iv[len(iv)-i] = byte(pad) ^ pt[len(pt)-i]
		}
		found := false
		for candidate := 0; candidate <= math.MaxUint8; candidate++ {
			iv[len(iv)-pad] = byte(candidate)
			if s.hasValidPadding(iv, block) {
				if pad == 1 {
					iv[len(iv)-2] ^= 1
					if !s.hasValidPadding(iv, block) {
						continue
					}
				}
				pt[len(pt)-pad] = byte(candidate) ^ byte(pad)
				found = true
				break
			}
		}
		if !found {
			panic("couldn't find padding")
		}
	}
	for i := range pt {
		pt[i] ^= iv[i]
	}
	return pt
}

func decipher(iv, ct []byte, s *service) []byte {
	var pt bytes.Buffer
	for i := 0; i < len(ct); i += blockSize {
		pt.Write(decipherSingleBlock(iv, ct[i:i+blockSize], s))
		iv = ct[i : i+16]
	}
	return pt.Bytes()
}

func main() {
	srv := newService()

	pt := getRandPt()
	iv, ct := srv.encrypt(pt)

	deciphered, err := sym.UnpadPKCS7(decipher(iv, ct, srv))
	if err != nil {
		panic("deciphered is not padded correctly")
	}
	if !bytes.Equal(deciphered, pt) {
		panic("deciphered != pt, padding oracle attack failed")
	}
	fmt.Println(string(deciphered))
}
