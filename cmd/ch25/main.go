package main

import (
	"fmt"
	"math"

	"github.com/doc-smith/cryptopals/util/crypto/rand"
	"github.com/doc-smith/cryptopals/util/crypto/sym"
)

type task struct {
	nonce uint64
	key   []byte
	ct    []byte
}

func newTask(secretPt []byte) *task {
	key := rand.RandBytes(16)
	nonce := rand.RandUint64(math.MaxUint64)
	ct := sym.EncryptAesCtr(secretPt, key, nonce)
	return &task{
		nonce,
		key,
		ct,
	}
}

func (s *task) edit(offset int, newPt []byte) {
	if offset+len(newPt) > len(s.ct) {
		panic("edit out of bounds")
	}
	newCt := sym.EncryptAesCtr(newPt, s.key, s.nonce)
	copy(s.ct[offset:], newCt)
}

func decipherPt(task *task) []byte {
	ct := make([]byte, len(task.ct))
	copy(ct, task.ct)

	zeros := make([]byte, len(ct))
	task.edit(0, zeros)

	ks := task.ct
	for i := 0; i < len(ct); i++ {
		ct[i] ^= ks[i]
	}

	return ct
}

func solve() []byte {
	const secretPt = `Rollin' in my 5.0
With my rag-top down so my hair can blow
The girlies on standby waving just to say hi
Did you stop? No, I just drove by`
	task := newTask([]byte(secretPt))
	return decipherPt(task)
}

func main() {
	pt := solve()
	fmt.Println(string(pt))
}
