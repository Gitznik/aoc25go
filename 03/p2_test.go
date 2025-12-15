package main

import "testing"

func TestP2(t *testing.T) {
	t.Run("Known input", func(t *testing.T) {
		inp := []byte(`987654321111111
811111111111119
234234234234278
818181911112111`)
		res := p2(inp)
		exp := 3121910778619
		if res != exp {
			t.Fatalf("got wrong result, expected %d, got %d", exp, res)
		}
	})
}
