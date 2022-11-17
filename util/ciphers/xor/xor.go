package xor

func EncryptSingleByteXor(plaintext []byte, key byte) []byte {
	ciphertext := make([]byte, len(plaintext))
	for i, b := range plaintext {
		ciphertext[i] = b ^ key
	}
	return ciphertext
}

func DecryptSingleByteXor(ciphertext []byte, key byte) []byte {
	return EncryptSingleByteXor(ciphertext, key)
}
