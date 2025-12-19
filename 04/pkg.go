package main

import (
	util "github.com/gitznik/aoc/pkg"
)

func FreeNeighbors(g util.Grid[rune], x, y int) int {
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

func BlockedNeighbors(g util.Grid[rune], x, y int) int {
	return 8 - FreeNeighbors(g, x, y)
}
