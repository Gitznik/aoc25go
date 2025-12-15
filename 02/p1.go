package main

import (
	"bytes"
	"fmt"

	util "github.com/gitznik/aoc/pkg"
)

func p1(inp []byte) int {
	res := 0

	for r := range bytes.SplitSeq(inp, []byte(",")) {
		invalid := InvalidInRange(r, IsValidID)
		if len(invalid) > 0 {
			res += util.Sum(invalid)
		}
	}
	fmt.Printf("Part 1 result: %d\n", res)
	return res
}
