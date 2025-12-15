package util

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type Part int

const (
	POne Part = iota
	PTwo
)

func ReadInput(part Part) []byte {
	var fp string
	switch part {
	default:
		panic(fmt.Errorf("unkown part: %v", part))
	case POne:
		fp = "part1"
	case PTwo:
		fp = "part2"
	}
	f, err := os.ReadFile(fp)
	if err != nil {
		panic(fmt.Errorf("could not read input: %w", err))
	}
	return bytes.Trim(f, "\n")
}

func ToInts(inp []byte) ([]int, error) {
	sinp := string(inp)
	res := make([]int, len(sinp))
	for i, c := range sinp {
		v, err := strconv.Atoi(string(c))
		if err != nil {
			return nil, err
		}
		res[i] = v
	}
	return res, nil
}
