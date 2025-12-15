package main

import "testing"

func TestLimitedBank(t *testing.T) {
	tests := []struct {
		name     string
		inp      []byte
		expected int
	}{
		{
			name:     "987654321111111",
			inp:      []byte("987654321111111"),
			expected: 987654321111,
		},
		{
			name:     "811111111111119",
			inp:      []byte("811111111111119"),
			expected: 811111111119,
		},
		{
			name:     "234234234234278",
			inp:      []byte("234234234234278"),
			expected: 434234234278,
		},
		{
			name:     "818181911112111",
			inp:      []byte("818181911112111"),
			expected: 888911112111,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := parseLimitedBank(tt.inp, 12)
			if res != tt.expected {
				t.Fatalf("got wrong result %v, expected %v", res, tt.expected)
			}
		})
	}
}
