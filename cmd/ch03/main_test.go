package main

import "testing"

func TestSolve(t *testing.T) {
	const expected = "Cooking MC's like a pound of bacon"
	actual := Solve(
		"1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	if actual != expected {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}
