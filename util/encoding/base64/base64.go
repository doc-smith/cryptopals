package base64

const (
	Padding  byte = '='
	Alphabet      = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

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
