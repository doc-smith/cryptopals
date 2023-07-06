package main

import "testing"

func TestSolve(t *testing.T) {
	const want = "746865206b696420646f6e277420706c6179"
	got := solve(
		"1c0111001f010100061a024b53535009181c",
		"686974207468652062756c6c277320657965")
	if got != want {
		t.Errorf("want %s, got %s", want, got)
	}
}
