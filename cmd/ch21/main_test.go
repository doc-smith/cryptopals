package main

import (
	"testing"

	"github.com/doc-smith/cryptopals/util/crypto/rand"
)

func TestMersenneTwister(t *testing.T) {
	// The test has the same requirement as the C++ implementation:
	//   10000-th consecutive invocation is required to produce 4123659995
	// See https://en.cppreference.com/w/cpp/numeric/random/mersenne_twister_engine

	const (
		n    = 10000
		seed = 5489
	)

	twister := rand.NewMersenneTwister(seed)
	var val uint32
	for i := 0; i < n; i++ {
		val = twister.Next()
	}

	const want = 4123659995
	if val != want {
		t.Fatalf("%d-th valuer must be %d, got %d", n, want, val)
	}
}
