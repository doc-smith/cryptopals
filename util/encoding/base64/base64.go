package base64

import (
	"errors"
	"fmt"
)

var LengthError = errors.New("input length must be a multiple of 4")

type InvalidByteError int

func (e InvalidByteError) Error() string {
	return fmt.Sprintf("invalid input byte at offset %d", int(e))
}

const (
	InvalidByteValue byte = 0xff
	Padding          byte = '='
	Alphabet              = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

var reverseAlphabet [256]byte

func init() {
	for i := 0; i < len(reverseAlphabet); i++ {
		reverseAlphabet[i] = InvalidByteValue
	}

	for i := 0; i < len(Alphabet); i++ {
		reverseAlphabet[Alphabet[i]] = byte(i)
	}
}

func Encode(b []byte) string {
	var res []byte

	sextetOffsets := [...]int{18, 12, 6, 0}

	si := 0
	for ; si < len(b); si += 3 {
		chunk := uint(b[si])<<16 | uint(b[si+1])<<8 | uint(b[si+2])
		for _, shift := range sextetOffsets {
			res = append(res, Alphabet[(chunk>>shift)&0b111111])
		}
	}

	remaining := len(b) - si

	switch remaining {
	case 1:
		chunk := uint(b[si]) << 16
		for _, shift := range sextetOffsets[:2] {
			res = append(res, Alphabet[(chunk>>shift)&0b111111])
		}
		res = append(res, Padding, Padding)
	case 2:
		chunk := uint(b[si])<<16 | uint(b[si+1])<<8
		for _, shift := range sextetOffsets[:3] {
			res = append(res, Alphabet[(chunk>>shift)&0b111111])
		}
		res = append(res, Padding)
	}

	return string(res)
}

func decodeSegment(s []byte) ([]byte, error) {
	sextetOffsets := [...]int{18, 12, 6, 0}

	si := 0
	var chunk uint

	for ; si < 4 && s[si] != Padding; si++ {
		x := reverseAlphabet[s[si]]
		if x == InvalidByteValue {
			return nil, InvalidByteError(si)
		}
		chunk |= uint(x) << sextetOffsets[si]
	}

	for i := si; i < 4; i++ {
		if s[i] != Padding {
			return nil, InvalidByteError(i)
		}
	}

	res := make([]byte, si-1)
	for i := 0; i < si-1; i++ {
		// 16, 8, 0
		offset := 16 - 8*i
		res[i] = byte((chunk >> offset) & 0xff)
	}
	return res[:], nil
}

func DecodeString(s string) ([]byte, error) {
	var res []byte
	si := 0
	for ; si < len(s); si += 4 {
		segment := []byte(s[si : si+4])
		decoded, err := decodeSegment(segment)
		if err != nil {
			return nil, err
		}
		res = append(res, decoded...)
	}
	if len(s)%4 != 0 {
		return nil, LengthError
	}
	return res, nil
}
