package main

import "testing"

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

func TestGrid_GetItem(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		inp []byte
		// Named input parameters for target function.
		x       int
		y       int
		want    rune
		wantErr bool
	}{
		{
			name:    "get @ from edge",
			inp:     []byte(inp),
			x:       0,
			y:       0,
			want:    '@',
			wantErr: false,
		},
		{
			name:    "get . from edge",
			inp:     []byte(inp),
			x:       1,
			y:       0,
			want:    '.',
			wantErr: false,
		},
		{
			name:    "get . from upper edge",
			inp:     []byte(inp),
			x:       9,
			y:       9,
			want:    '.',
			wantErr: false,
		},
		{
			name:    "get out of bounds",
			inp:     []byte(inp),
			x:       100,
			y:       0,
			want:    '.',
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGrid(tt.inp)
			got, gotErr := g.GetItem(tt.x, tt.y)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetItem() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetItem() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if tt.want != got {
				t.Errorf("GetItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			g := NewGrid(tt.inp)
			got := g.FreeNeighbors(tt.x, tt.y)
			if got != tt.want {
				t.Errorf("FreeNeighbors() = %v, want %v", got, tt.want)
			}
		})
	}
}
