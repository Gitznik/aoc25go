package main

import "testing"

func TestKnownInputd2(t *testing.T) {
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
	r := d2(inp)
	if r != 6 {
		t.Fatalf("Expected 6, got %d", r)
	}
}
