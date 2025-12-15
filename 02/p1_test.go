package main

import "testing"

func TestP1(t *testing.T) {
	t.Run("Known input", func(t *testing.T) {
		inp := []byte("11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124")
		res := p1(inp)
		exp := 1227775554
		if res != exp {
			t.Fatalf("got wrong result, expected %d, got %d", exp, res)
		}
	})
}
