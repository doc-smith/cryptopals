package main

import "testing"

func TestSolve(t *testing.T) {
	const expected = "746865206b696420646f6e277420706c6179"
	actual := Solve(
		"1c0111001f010100061a024b53535009181c",
		"686974207468652062756c6c277320657965")
	if actual != expected {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}
