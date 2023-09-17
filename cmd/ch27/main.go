package main

import (
	"bytes"
	"crypto/aes"
	"log"

	"github.com/doc-smith/cryptopals/util/crypto/rand"
	"github.com/doc-smith/cryptopals/util/crypto/sym"
)

// CBC mode with IV = key
//
// https://en.wikipedia.org/wiki/Block_cipher_mode_of_operation
//
// Decryption:
//
// C1 | 0 | C1
//
// C1 = E(P1 ^ key)
//
// P1 = D(C1) ^ key = D(E(P1 ^ key)) ^ key = P1 ^ key ^ key = P1
// P2 = ...
// P3 = D(C1) ^ 0 = D(E(P1 ^ key)) = P1 ^ key
// key = P1 ^ P3

type service struct {
	key []byte
}

func newService() *service {
	return &service{
		key: rand.RandBytes(16),
	}
}

func (s *service) encrypt(pt []byte) []byte {
	iv := s.key
	return sym.EncryptAesCbc(pt, s.key, iv)
}

func (s *service) decrypt(ct []byte) []byte {
	iv := s.key
	return sym.DecryptAesCbc(ct, s.key, iv)
}

func (s *service) testKey(candidate []byte) bool {
	return bytes.Equal(s.key, candidate)
}

func recoverKey(s *service) []byte {
	const pt = "hello world! thi"
	if len(pt) != aes.BlockSize {
		panic("plaintext must be exactly one block")
	}
	ct := s.encrypt([]byte(pt))

	var evil bytes.Buffer
	evil.Write(ct)
	zeros := make([]byte, aes.BlockSize)
	evil.Write(zeros)
	evil.Write(ct)

	evilPt := s.decrypt(evil.Bytes())
	p1 := evilPt[:aes.BlockSize]
	p3 := evilPt[2*aes.BlockSize : 3*aes.BlockSize]

	key := make([]byte, aes.BlockSize)
	for i := range key {
		key[i] = p1[i] ^ p3[i]
	}

	return key
}

func solve() {
	s := newService()
	key := recoverKey(s)
	if !s.testKey(key) {
		log.Fatalf("attack failed")
	}
}

func main() {
	solve()
}
