package main

import (
	"fmt"

	util "github.com/gitznik/aoc/pkg"
)

func p1(inp []byte) int {
	total := 0
	g := util.NewGrid(inp, func(r rune) rune { return r })
	for y, r := range g {
		for x := range r {
			if v, err := g.GetItem(x, y); err != nil || v != '@' {
				continue
			}
			if BlockedNeighbors(g, x, y) < 4 {
				total++
			}
		}
	}
	fmt.Printf("Part 1 result: %d\n", total)
	return total
}
