package main

import "testing"

func TestSolve(t *testing.T) {
	const expected = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	actual := Solve(
		"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	if actual != expected {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}
