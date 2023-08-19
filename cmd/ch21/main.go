package main

import "github.com/doc-smith/cryptopals/util/crypto/rand"

func main() {
	const seed = 123
	twister := rand.NewMersenneTwister(seed)
	for i := 0; i < 10; i++ {
		println(twister.Next())
	}
}
