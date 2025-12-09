package main

import "testing"

func TestKnownInput(t *testing.T) {
	inp := []byte(`
L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
	`)
	r := d1(inp)
	if r != 3 {
		t.Fatalf("Expected 3, got %d", r)
	}
}
