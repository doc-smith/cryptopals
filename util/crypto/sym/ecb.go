package sym

import (
	"crypto/aes"
	"crypto/cipher"
)

func validateInpLen(inp []byte) {
	if len(inp)%aes.BlockSize != 0 {
		panic("input length is not the multiple of the block size")
	}
}

func transform(
	src, key []byte,
	f func(cipher.Block, []byte, []byte)) []byte {

	validateInpLen(src)
	cipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	dst := make([]byte, len(src))
	for i := 0; i < len(src); i += aes.BlockSize {
		f(cipher, dst[i:], src[i:])
	}
	return dst
}

func EncryptAesEcb(pt, key []byte) []byte {
	return transform(pt, key, func(cp cipher.Block, dst, src []byte) {
		cp.Encrypt(dst, src)
	})
}

func DecryptAesEcb(ct, key []byte) []byte {
	return transform(ct, key, func(cp cipher.Block, dst, src []byte) {
		cp.Decrypt(dst, src)
	})
}
