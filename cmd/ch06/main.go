package main

import (
	"fmt"
	"math"
	"math/bits"
	"os"
	"sort"

	"github.com/doc-smith/cryptopals/util/inp"
	"github.com/doc-smith/cryptopals/util/lang"
)

const (
	minKeySize = 2
	maxKeySize = 32
	numKeys    = 4
)

func hammingDistance(x, y []byte) int {
	distance := 0
	for i := range x {
		distance += bits.OnesCount8(x[i] ^ y[i])
	}
	return distance
}

func calcKeySizeDistance(ct []byte, sz int) float64 {
	s := 0.0
	n := 0
	for i := 2 * sz; i < len(ct); i += sz {
		cur := ct[i-sz : i]
		prv := ct[i-2*sz : i-sz]
		s += float64(hammingDistance(cur, prv)) / float64(sz)
		n++
	}
	return s / float64(n)
}

func findKeySizes(ct []byte, min, max, n int) []int {
	type Candidate struct {
		Size     int
		Distance float64
	}

	candidates := make([]Candidate, 0)
	for sz := min; sz <= max; sz++ {
		distance := calcKeySizeDistance(ct, sz)
		candidates = append(candidates, Candidate{sz, distance})
	}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].Distance < candidates[j].Distance
	})

	res := make([]int, 0, n)
	for _, c := range candidates[:n] {
		res = append(res, c.Size)
	}
	return res
}

func recoverKeyByte(ct []byte, k int, keySize int) byte {
	var bestScore float64
	var bestKeyByte byte
	for b := 0; b <= math.MaxUint8; b++ {
		freq := make(map[byte]float64)
		n := 0
		for i := k; i < len(ct); i += keySize {
			freq[uint8(b)^ct[i]]++
			n++
		}
		for b := range freq {
			freq[b] /= float64(n)
		}
		score := lang.ScoreEnglishTextByteFrequencies(freq)
		if score > bestScore {
			bestScore = score
			bestKeyByte = uint8(b)
		}
	}
	return bestKeyByte
}

func recoverKey(ct []byte, keySize int) []byte {
	k := make([]byte, keySize)
	for i := range k {
		k[i] = recoverKeyByte(ct, i, keySize)
	}
	return k
}

func decrypt(dst, ct, key []byte) {
	for i := range dst {
		dst[i] = ct[i] ^ key[i%len(key)]
	}
}

func cloneBytes(src []byte) []byte {
	dst := make([]byte, len(src))
	copy(dst, src)
	return dst
}

func decipher(ct []byte) []byte {
	var bestScore float64
	var bestPt []byte

	pt := make([]byte, len(ct))
	for _, keySize := range findKeySizes(ct, minKeySize, maxKeySize, numKeys) {
		key := recoverKey(ct, keySize)
		decrypt(pt, ct, key)
		score := lang.ScoreEnglishText(pt)
		if score > bestScore {
			bestScore = score
			bestPt = cloneBytes(pt)
		}
	}

	return bestPt
}

func solve(ct []byte) []byte {
	return decipher(ct)
}

func main() {
	inp := inp.ReadBase64(os.Stdin)
	ans := solve(inp)
	fmt.Println(string(ans))
}
