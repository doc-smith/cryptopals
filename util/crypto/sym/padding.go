package sym

import (
	"fmt"
	"math"
)

func PadPKCS7(b []byte, blen int) []byte {
	plen := blen - len(b)%blen
	if plen > math.MaxUint8 {
		panic(fmt.Sprintf("cannot pad more than %d bytes", math.MaxUint8))
	}
	for i := 0; i < plen; i++ {
		b = append(b, byte(plen))
	}
	return b
}

func UnpadPKCS7(b []byte) ([]byte, error) {
	if len(b) == 0 {
		return make([]byte, 0), nil
	}
	padLen := int(b[len(b)-1])
	for _, x := range b[len(b)-padLen:] {
		if x != byte(padLen) {
			return nil, fmt.Errorf("invalid padding")
		}
	}
	res := make([]byte, len(b)-padLen)
	copy(res, b[:len(res)])
	return res, nil
}
