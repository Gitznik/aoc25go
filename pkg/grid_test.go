package util_test

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

func TestGrid_GetItem(t *testing.T) {
	tests := []struct {
		name    string
		inp     []byte
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
			name:    "get out of bounds xl",
			inp:     []byte(inp),
			x:       100,
			y:       0,
			want:    '.',
			wantErr: true,
		},
		{
			name:    "get out of bounds xn",
			inp:     []byte(inp),
			x:       -1,
			y:       0,
			want:    '.',
			wantErr: true,
		},
		{
			name:    "get out of bounds yl",
			inp:     []byte(inp),
			x:       0,
			y:       100,
			want:    '.',
			wantErr: true,
		},
		{
			name:    "get out of bounds yn",
			inp:     []byte(inp),
			x:       0,
			y:       -1,
			want:    '.',
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := util.NewGrid(tt.inp, func(r rune) rune { return r })
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
			if tt.want != got {
				t.Errorf("GetItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_SetItem(t *testing.T) {
	tests := []struct {
		name    string
		inp     []byte
		x       int
		y       int
		v       rune
		wantErr bool
	}{
		{
			name:    "set item in bounds",
			inp:     []byte(inp),
			x:       0,
			y:       0,
			v:       0,
			wantErr: false,
		},
		{
			name:    "set item out of bounds",
			inp:     []byte(inp),
			x:       -1,
			y:       0,
			v:       0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := util.NewGrid(tt.inp, func(r rune) rune { return r })
			gotErr := g.SetItem(tt.x, tt.y, tt.v)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("SetItem() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("SetItem() succeeded unexpectedly")
			}
			item, err := g.GetItem(tt.x, tt.y)
			if err != nil {
				t.Error("Could not get item after setting")
			}
			if item != tt.v {
				t.Errorf("GetItem() = %v, want %v", item, tt.v)
			}
		})
	}
}
