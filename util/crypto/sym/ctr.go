package sym

import (
	"bytes"
	"crypto/aes"
	"encoding/binary"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func EncryptAesCtr(pt, key []byte, nonce uint64) []byte {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	var ct bytes.Buffer
	const blockSize = 16

	ctr := make([]byte, blockSize)
	binary.LittleEndian.PutUint64(ctr, nonce)

	for i := 0; i < len(pt); i += blockSize {
		keystream := make([]byte, blockSize)
		binary.LittleEndian.PutUint64(ctr[8:], uint64(i/blockSize))
		cipher.Encrypt(keystream, ctr)

		ctBlock := make([]byte, min(blockSize, len(pt[i:])))
		for j := range ctBlock {
			ctBlock[j] = pt[i+j] ^ keystream[j]
		}
		ct.Write(ctBlock)
	}
	return ct.Bytes()
}

func DecryptAesCtr(ct, key []byte, nonce uint64) []byte {
	return EncryptAesCtr(ct, key, nonce)
}
