package sym

import (
	"crypto/aes"
)

func xorBytes(dst, x, y []byte) {
	for i := range x {
		dst[i] = x[i] ^ y[i]
	}
}

func EncryptAesCbc(pt, key, iv []byte) []byte {
	validateInpLen(pt)
	if len(iv) != aes.BlockSize {
		panic("IV must have the same length as one AES block")
	}
	cipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	ct := make([]byte, len(pt))
	prevBlockCt := iv

	for i := 0; i < len(pt); i += aes.BlockSize {
		blockPt := pt[i : i+aes.BlockSize]
		blockCt := ct[i : i+aes.BlockSize]
		xorBytes(blockCt, blockPt, prevBlockCt)
		cipher.Encrypt(blockCt, blockCt)
		prevBlockCt = blockCt
	}

	return ct
}

func DecryptAesCbc(ct, key, iv []byte) []byte {
	validateInpLen(ct)
	if len(iv) != aes.BlockSize {
		panic("IV must have the same length as one AES block")
	}
	cipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	pt := make([]byte, len(ct))

	for i := len(ct) - aes.BlockSize; i > 0; i -= aes.BlockSize {
		blockPt := pt[i : i+aes.BlockSize]
		blockCt := ct[i : i+aes.BlockSize]
		cipher.Decrypt(blockPt, blockCt)
		prevBlockCt := ct[i-aes.BlockSize : i]
		xorBytes(blockPt, blockPt, prevBlockCt)
	}

	cipher.Decrypt(pt[:aes.BlockSize], ct[:aes.BlockSize])
	xorBytes(pt[:aes.BlockSize], pt[:aes.BlockSize], iv)

	return pt
}
