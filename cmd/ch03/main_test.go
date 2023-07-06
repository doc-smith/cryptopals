package main

import "testing"

func TestSolve(t *testing.T) {
	const want = "Cooking MC's like a pound of bacon"
	got := solve(
		"1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}
