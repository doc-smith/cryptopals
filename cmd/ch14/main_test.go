package main

import (
	"bytes"
	"testing"
)

func TestSolve(t *testing.T) {
	want := []byte(`Rollin' in my 5.0
With my rag-top down so my hair can blow
The girlies on standby waving just to say hi
Did you stop? No, I just drove by`)
	got := solve()
	if bytes.Equal(got, want) {
		t.Errorf("Solve() = %q; want %q", got, want)
	}
}
