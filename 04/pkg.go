package main

import (
	"bytes"
	"errors"
	"slices"
)

type Row []rune

type Grid []Row

var ErrOutOfRange = errors.New("index out of range")

func (g Grid) GetItem(x, y int) (rune, error) {
	var zero rune
	if x < 0 || y < 0 {
		return zero, ErrOutOfRange
	}
	if y >= len(g) {
		return zero, ErrOutOfRange
	}
	if x >= len(g[0]) {
		return zero, ErrOutOfRange
	}
	return g[y][x], nil
}

func NewGrid(inp []byte) Grid {
	y := make(Grid, 0)
	for l := range bytes.SplitSeq(inp, []byte("\n")) {
		x := make(Row, len(l))
		for i, r := range string(l) {
			x[i] = r
		}
		y = append(y, x)
	}
	slices.Reverse(y)
	return y
}

func (g Grid) SetEmpty(x, y int) error {
	if x < 0 || y < 0 {
		return ErrOutOfRange
	}
	if y >= len(g) {
		return ErrOutOfRange
	}
	if x >= len(g[0]) {
		return ErrOutOfRange
	}
	g[y][x] = '.'
	return nil
}

func (g Grid) FreeNeighbors(x, y int) int {
	if _, err := g.GetItem(x, y); err != nil {
		panic(err)
	}
	free := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			r, err := g.GetItem(x+i, y+j)
			if err != nil || r == '.' {
				free++
			}
		}
	}
	return free
}

func (g Grid) BlockedNeighbors(x, y int) int {
	return 8 - g.FreeNeighbors(x, y)
}
