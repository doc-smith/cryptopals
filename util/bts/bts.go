package bts

import "bytes"

func Concat(bs ...[]byte) []byte {
	var buf bytes.Buffer
	for _, b := range bs {
		buf.Write(b)
	}
	return buf.Bytes()
}
