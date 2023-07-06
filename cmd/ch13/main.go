package main

import (
	"bytes"
	"crypto/aes"
	"fmt"
	"net/url"
	"strings"

	"github.com/doc-smith/cryptopals/util/crypto/rand"
	"github.com/doc-smith/cryptopals/util/crypto/sym"
)

func parseCookie(c string) (map[string]string, error) {
	m, err := url.ParseQuery(c)
	if err != nil {
		return nil, fmt.Errorf("cannot parse cookie")
	}
	res := make(map[string]string)
	for k, vs := range m {
		for _, v := range vs {
			res[k] = v
		}
	}
	return res, nil
}

func profileFor(email string) string {
	safeEmail := email
	for _, c := range []string{"&", "="} {
		safeEmail = strings.ReplaceAll(safeEmail, c, "")
	}
	return fmt.Sprintf("email=%s&uid=10&role=user", safeEmail)
}

func getCookie(email string, key []byte) []byte {
	profile := profileFor(email)
	padded := sym.PadPKCS7([]byte(profile), aes.BlockSize)
	return sym.EncryptAesEcb(padded, key)
}

func isAdmin(encryptedCookie, key []byte) bool {
	pt := sym.DecryptAesEcb(encryptedCookie, key)
	unpadded, err := sym.UnpadPKCS7(pt)
	if err != nil {
		panic(err)
	}
	cookie, err := parseCookie(string(unpadded))
	if err != nil {
		panic(err)
	}
	return cookie["role"] == "admin"
}

func ecbCutAndPaste() {
	// 0123456789abcdef
	//
	// email=xxxxxxxxxx
	// admin........... (11)
	// xxx&uid=10&role=
	// user............
	//
	// email=xxxxxxxxxx (0)
	// xxx&uid=10&role= (32)
	// admin........... (16)

	key := rand.RandBytes(16)
	email := "xxxxxxxxxxadmin" + strings.Repeat("\x0b", 11) + "xxx"

	cookie := getCookie(email, key)

	var attack bytes.Buffer
	attack.Write(cookie[:16])
	attack.Write(cookie[32:48])
	attack.Write(cookie[16:32])

	if !isAdmin(attack.Bytes(), key) {
		panic("ECB cut-and-paste attack failed")
	}
}

func main() {
	ecbCutAndPaste()
}
