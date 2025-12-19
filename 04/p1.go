package main

import (
	"fmt"
)

func p1(inp []byte) int {
	total := 0
	g := NewGrid(inp)
	for y, r := range g {
		for x := range r {
			if v, err := g.GetItem(x, y); err != nil || v != '@' {
				continue
			}
			if g.BlockedNeighbors(x, y) < 4 {
				total++
			}
		}
	}
	fmt.Printf("Part 1 result: %d\n", total)
	return total
}
