package main

import (
	"bytes"
	"fmt"
)

func p1(inp []byte) int {
	total := 0
	for l := range bytes.SplitSeq(inp, []byte("\n")) {
		total += parseLimitedBank(l, 2)
	}
	fmt.Printf("Part 1 result: %d\n", total)
	return total
}
