package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func d1(inp []byte) int {
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
			dp, _ = dp.Sub(v)
		} else {
			dp, _ = dp.Add(v)
		}
		if int(dp) == 0 {
			c++
		}
	}
	fmt.Printf("Day1 result: %d\n", c)
	return c
}
