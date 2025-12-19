package main

import (
	"fmt"
)

func p2(inp []byte) int {
	total := 0
	g := NewGrid(inp)
	for {
		rm := findRemovable(g)
		if len(rm) == 0 {
			break
		}
		total += len(rm)
		for _, v := range rm {
			if err := g.SetEmpty(v.X, v.Y); err != nil {
				panic(err)
			}
		}
	}
	fmt.Printf("Part 2 result: %d\n", total)
	return total
}

type Pos struct {
	X int
	Y int
}

func findRemovable(g Grid) []Pos {
	rm := make([]Pos, 0)
	for y, r := range g {
		for x := range r {
			if v, err := g.GetItem(x, y); err != nil || v != '@' {
				continue
			}
			if g.BlockedNeighbors(x, y) < 4 {
				rm = append(rm, Pos{x, y})
			}
		}
	}
	return rm
}
