package main

import "testing"

func TestIsValidID(t *testing.T) {
	tests := []struct {
		name     string
		inp      []byte
		expected bool
	}{
		{
			name:     "Simple Valid",
			inp:      []byte("95"),
			expected: true,
		},
		{
			name:     "Uneven must be true",
			inp:      []byte("100"),
			expected: true,
		},
		{
			name:     "Starts with 0",
			inp:      []byte("0101"),
			expected: false,
		},
		{
			name:     "Simple Invalid",
			inp:      []byte("99"),
			expected: false,
		},
		{
			name:     "Long Invalid",
			inp:      []byte("446446"),
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := IsValidID(tt.inp)
			if res != tt.expected {
				t.Fatalf("got wrong result, expected %v", tt.expected)
			}
		})
	}
}

func TestIsValidIDP2(t *testing.T) {
	tests := []struct {
		name     string
		inp      []byte
		expected bool
	}{
		{
			name:     "Simple Valid",
			inp:      []byte("95"),
			expected: true,
		},
		{
			name:     "Uneven different",
			inp:      []byte("100"),
			expected: true,
		},
		{
			name:     "Uneven mirrored",
			inp:      []byte("101"),
			expected: true,
		},
		{
			name:     "Uneven identical",
			inp:      []byte("111"),
			expected: false,
		},
		{
			name:     "Starts with 0",
			inp:      []byte("0101"),
			expected: false,
		},
		{
			name:     "Simple Invalid",
			inp:      []byte("99"),
			expected: false,
		},
		{
			name:     "Long Invalid",
			inp:      []byte("446446"),
			expected: false,
		},
		{
			name:     "Repeating pattern",
			inp:      []byte("123123123"),
			expected: false,
		},
		{
			name:     "Repeating pattern 2",
			inp:      []byte("12341234"),
			expected: false,
		},
		{
			name:     "Repeating pattern 3",
			inp:      []byte("1212121212"),
			expected: false,
		},
		{
			name:     "Repeating pattern 4",
			inp:      []byte("1111111"),
			expected: false,
		},
		{
			name:     "Repeating pattern 5",
			inp:      []byte("11121112"),
			expected: false,
		},
		{
			name:     "Repeating pattern 6",
			inp:      []byte("101101"),
			expected: false,
		},
		{
			name:     "Repeating pattern 7",
			inp:      []byte("10101"),
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := IsValidP2Repeating(tt.inp)
			if res != tt.expected {
				t.Fatalf("got wrong result, expected %v", tt.expected)
			}
		})
	}
}

func TestInvalidInRange(t *testing.T) {
	tests := []struct {
		name     string
		inp      []byte
		expected []int
	}{
		{
			name:     "multiple invalid",
			inp:      []byte("11-22"),
			expected: []int{11, 22},
		},
		{
			name:     "one invalid",
			inp:      []byte("99-115"),
			expected: []int{99},
		},
		{
			name:     "No invalid",
			inp:      []byte("1698522-1698528"),
			expected: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := InvalidInRange(tt.inp, IsValidID)
			if len(res) != len(tt.expected) {
				t.Fatalf("got wrong result, expected %v values, got %v", len(tt.expected), res)
			}
			for i := range len(res) {
				if res[i] != tt.expected[i] {
					t.Fatalf("got wrong result in index %d, expected %d, got %d", i, res[i], tt.expected[i])
				}
			}
		})
	}
}
