package main

import (
	"bytes"
	"fmt"
)

func p2(inp []byte) int {
	total := 0
	for l := range bytes.SplitSeq(inp, []byte("\n")) {
		total += parseLimitedBank(l, 12)
	}
	fmt.Printf("Part 2 result: %d\n", total)
	return total
}
