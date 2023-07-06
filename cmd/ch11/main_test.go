package main

import "testing"

func TestHasTwoEqualBlocks(t *testing.T) {
	tests := []struct {
		name string
		ct   []byte
		blen int
		want bool
	}{
		{
			name: "empty",
			ct:   []byte{},
			blen: 1,
			want: false,
		},
		{
			name: "single block",
			ct:   []byte{0x01, 0x02, 0x03, 0x04},
			blen: 4,
			want: false,
		},
		{
			name: "two blocks",
			ct:   []byte{0x01, 0x02, 0x01, 0x02},
			blen: 2,
			want: true,
		},
		{
			name: "two and a half blocks",
			ct:   []byte{0x01, 0x02, 0x01, 0x02, 0x01},
			blen: 2,
			want: true,
		},
		{
			name: "two blocks (different)",
			ct:   []byte{0x01, 0x02, 0x03, 0x04},
			blen: 2,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasTwoEqualBlocks(tt.ct, tt.blen); got != tt.want {
				t.Errorf("hasTwoEqualBlocks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSmoke(t *testing.T) {
	const iterCnt = 100
	testDetectionOracle(iterCnt)
}
