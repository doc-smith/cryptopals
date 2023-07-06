package rand

import (
	"crypto/rand"
	"math/big"
)

func RandBytes(n int) []byte {
	res := make([]byte, n)
	if _, err := rand.Read(res); err != nil {
		panic(err)
	}
	return res
}

func RandInt(n int) int {
	max := big.NewInt(int64(n))
	num, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}
	return int(num.Int64())
}

type Face = uint8

const (
	Head Face = iota
	Tail
)

func CoinFlip() Face {
	r := RandInt(1)
	if r == 0 {
		return Head
	}
	return Tail
}
