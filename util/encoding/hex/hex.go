package hex

import (
	"errors"
	"math"
)

const (
	InvalidByteValue = math.MaxUint8
	Alphabet         = "0123456789abcdef"
)

var LengthError = errors.New("hex string must be even in length")

type InvalidByteError struct {
	Offset int
	Msg    string
}

func (e *InvalidByteError) Error() string {
	return e.Msg
}

func EncodeHexString(b []byte) string {
	var res []byte
	for _, v := range b {
		res = append(res, Alphabet[v>>4], Alphabet[v&0xf])
	}
	return string(res)
}

func decodeHexNibble(c byte) byte {
	switch {
	case c >= '0' && c <= '9':
		return c - '0'
	case c >= 'a' && c <= 'f':
		return c - 'a' + 10
	case c >= 'A' && c <= 'F':
		return c - 'A' + 10
	default:
		return InvalidByteValue
	}
}

func DecodeHexString(hex string) ([]byte, error) {
	var res []byte

	if len(hex)%2 != 0 {
		return nil, LengthError
	}

	for i := 0; i < len(hex); i += 2 {
		var nibbles [2]byte
		for j := 0; j < 2; j++ {
			n := decodeHexNibble(hex[i+j])
			if n == InvalidByteValue {
				return nil, &InvalidByteError{
					Offset: i + j,
					Msg:    "invalid byte in hex string",
				}
			}
			nibbles[j] = n
		}
		res = append(res, nibbles[0]<<4|nibbles[1])
	}

	return res, nil
}
