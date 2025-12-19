package main

import (
	"testing"

	util "github.com/gitznik/aoc/pkg"
)

const inp = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

func TestGrid_FreeNeighbors(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		inp []byte
		// Named input parameters for target function.
		x    int
		y    int
		want int
	}{
		{
			name: "neighbors from lower edge",
			inp:  []byte(inp),
			x:    0,
			y:    0,
			want: 7,
		},
		{
			name: "neighbors from upper edge",
			inp:  []byte(inp),
			x:    9,
			y:    9,
			want: 5,
		},
		{
			name: "neighbors from middle",
			inp:  []byte(inp),
			x:    5,
			y:    2,
			want: 2,
		},
		{
			name: "all alone",
			inp: []byte(`...
.@.
...`),
			x:    1,
			y:    1,
			want: 8,
		},
		{
			name: "upper row",
			inp: []byte(`@@@
.@.
...`),
			x:    1,
			y:    1,
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := util.NewGrid(tt.inp, func(r rune) rune { return r })
			got := FreeNeighbors(g, tt.x, tt.y)
			if got != tt.want {
				t.Errorf("FreeNeighbors() = %v, want %v", got, tt.want)
			}
		})
	}
}
