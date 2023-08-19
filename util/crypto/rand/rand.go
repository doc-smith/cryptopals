package rand

import (
	"crypto/rand"
	"math"
	"math/big"
)

func RandBytes(n int) []byte {
	res := make([]byte, n)
	if _, err := rand.Read(res); err != nil {
		panic(err)
	}
	return res
}

func RandInt64(n int64) int64 {
	max := big.NewInt(n)
	num, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}
	return num.Int64()
}

func RandUint64(n uint64) uint64 {
	if n <= math.MaxInt64 {
		return uint64(RandInt64(int64(n)))
	}

	hMax := int64(n >> 63)
	h := uint64(RandInt64(hMax))
	return h<<63 | uint64(RandInt64(math.MaxInt64))
}

func RandInt(n int) int {
	return int(RandInt64(int64(n)))
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
