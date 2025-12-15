package main

import (
	"strconv"
	"strings"

	util "github.com/gitznik/aoc/pkg"
)

func parseLimitedBank(bank []byte, limit int) int {
	values, err := util.ToInts(bank)
	if err != nil {
		panic(err)
	}
	res := make([]int, limit)
	ptr := 0
	for i := range limit {
		maxIdx := findIdxMax(values[ptr : len(values)-(limit-i-1)])
		res[i] = maxIdx + ptr
		ptr += maxIdx + 1
	}
	var sb strings.Builder
	for _, n := range res {
		sb.WriteString(strconv.Itoa(values[n]))
	}
	result := sb.String()
	j, err := strconv.Atoi(result)
	if err != nil {
		panic(err)
	}
	return j
}

func findIdxMax(inp []int) int {
	maxIdx := -1
	for i, v := range inp {
		if maxIdx == -1 || v > inp[maxIdx] {
			maxIdx = i
		}
	}
	return maxIdx
}
