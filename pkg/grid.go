package util

import (
	"bytes"
	"errors"
	"slices"
)

var ErrOutOfRange = errors.New("index out of range")

type Row[T any] []T

type Grid[T any] []Row[T]

func (g Grid[T]) GetItem(x, y int) (T, error) {
	if err := g.checkBounds(x, y); err != nil {
		var zero T
		return zero, err
	}
	return g[y][x], nil
}

func NewGrid[T any](inp []byte, converter func(rune) T) Grid[T] {
	y := make(Grid[T], 0)
	for l := range bytes.SplitSeq(inp, []byte("\n")) {
		x := make(Row[T], len(l))
		for i, r := range string(l) {
			x[i] = converter(r)
		}
		y = append(y, x)
	}
	slices.Reverse(y)
	return y
}

func (g Grid[T]) SetItem(x, y int, v T) error {
	if err := g.checkBounds(x, y); err != nil {
		return err
	}
	g[y][x] = v
	return nil
}

func (g Grid[T]) checkBounds(x, y int) error {
	if x < 0 || y < 0 {
		return ErrOutOfRange
	}
	if y >= len(g) {
		return ErrOutOfRange
	}
	if x >= len(g[0]) {
		return ErrOutOfRange
	}
	return nil
}
