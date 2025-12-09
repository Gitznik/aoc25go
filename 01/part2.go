package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func d2(inp []byte) int {
	clicks := 0
	c := 0
	dp := NewDialValue(50)
	for l := range bytes.SplitSeq(inp, []byte("\n")) {
		if len(l) < 2 {
			continue
		}
		v, err := strconv.Atoi(string(l[1:]))
		if err != nil {
			panic(fmt.Errorf("could not parse integer in line %s: %v", l, err))
		}
		if bytes.HasPrefix(l, []byte("L")) {
			dp, clicks = dp.Sub(v)
		} else {
			dp, clicks = dp.Add(v)
		}
		c += clicks

	}
	fmt.Printf("Day2 result: %d\n", c)
	return c
}
