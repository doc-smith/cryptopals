package hex

import (
	"errors"
	"fmt"
)

const (
	invalidByteValue byte = 0xff
	alphabet              = "0123456789abcdef"
)

var LengthError = errors.New("hex string must be even in length")

type InvalidByteError int

func (e InvalidByteError) Error() string {
	return fmt.Sprintf("invalid input byte at offset %d", int(e))
}

func EncodeHexString(b []byte) string {
	res := make([]byte, 0, len(b)*2)
	for _, v := range b {
		res = append(res, alphabet[v>>4], alphabet[v&0xf])
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
		return invalidByteValue
	}
}

func DecodeHexString(hex string) ([]byte, error) {
	if len(hex)%2 != 0 {
		return nil, LengthError
	}

	res := make([]byte, 0)

	for i := 0; i < len(hex); i += 2 {
		var nibbles [2]byte
		for j := 0; j < 2; j++ {
			offset := i + j
			n := decodeHexNibble(hex[offset])
			if n == invalidByteValue {
				return nil, InvalidByteError(offset)
			}
			nibbles[j] = n
		}
		res = append(res, nibbles[0]<<4|nibbles[1])
	}

	return res, nil
}
