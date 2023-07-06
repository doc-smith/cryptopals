package main

import (
	"bytes"
	"crypto/aes"
	"fmt"
	"math"
	"strings"

	"github.com/doc-smith/cryptopals/util/bts"
	"github.com/doc-smith/cryptopals/util/crypto/rand"
	"github.com/doc-smith/cryptopals/util/crypto/sym"
	"github.com/doc-smith/cryptopals/util/inp"
)

type encryptor struct {
	key []byte
}

func newEncryptor() *encryptor {
	return &encryptor{
		rand.RandBytes(16),
	}
}

func (enc *encryptor) encrypt(attackerControlled []byte) []byte {
	const b64Secret = `Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg
aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq
dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg
YnkK`
	secret := inp.ReadBase64(strings.NewReader(b64Secret))
	pt := bts.Concat(
		attackerControlled,
		secret)
	return sym.EncryptAesEcb(
		sym.PadPKCS7(pt, aes.BlockSize),
		enc.key)
}

type attackParams struct {
	blen, msgLen int
}

func detectAttackParams(enc *encryptor) attackParams {
	var pt []byte
	ct := enc.encrypt(pt)
	origCtLen := len(ct)

	for len(ct) == origCtLen {
		pt = append(pt, 'x')
		ct = enc.encrypt(pt)
	}

	return attackParams{
		blen:   len(ct) - origCtLen,
		msgLen: origCtLen - len(pt),
	}
}

func shiftLeft(b []byte) {
	for i := 1; i < len(b); i++ {
		b[i-1] = b[i]
	}
}

func decipherMessage(enc *encryptor, params attackParams) []byte {
	var prefix []byte
	for i := 0; i < params.msgLen; i++ {
		if i%params.blen == 0 {
			prefix = append(make([]byte, params.blen), prefix...)
		}
		ct := enc.encrypt(prefix[:len(prefix)-i-1])

		found := false
		for b := 0; b <= math.MaxUint8; b++ {
			prefix[len(prefix)-1] = byte(b)
			candidateCt := enc.encrypt(prefix)
			s := len(prefix) - params.blen
			if bytes.Equal(ct[s:s+params.blen], candidateCt[s:s+params.blen]) {
				found = true
				break
			}
		}
		if !found {
			panic(fmt.Sprintf("failed to find byte %d", i))
		}

		shiftLeft(prefix)
	}
	return prefix
}

func solve() []byte {
	enc := newEncryptor()
	return decipherMessage(enc, detectAttackParams(enc))
}

func main() {
	fmt.Println(string(solve()))
}
