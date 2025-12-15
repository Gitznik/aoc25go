package main

import (
	"bytes"
	"fmt"

	util "github.com/gitznik/aoc/pkg"
)

func p2(inp []byte) int {
	res := 0

	for r := range bytes.SplitSeq(inp, []byte(",")) {
		invalid := InvalidInRange(r, IsValidP2Repeating)
		if len(invalid) > 0 {
			res += util.Sum(invalid)
		}
	}
	fmt.Printf("Part 2 result: %d\n", res)
	return res

}
