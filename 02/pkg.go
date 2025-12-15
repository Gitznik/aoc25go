package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func IsValidID(id []byte) bool {
	if bytes.HasPrefix(id, []byte("0")) {
		return false
	}
	if len(id)%2 != 0 {
		return true
	}
	half := len(id) / 2
	h1 := id[0:half]
	h2 := id[half:]

	for i := range half {
		if h1[i] != h2[i] {
			return true
		}
	}

	return false
}

func RepeatsSequence(inp, seq []byte) bool {
	if len(inp)%len(seq) != 0 {
		return false
	}

	for i := 0; i < len(inp); i += len(seq) {
		if !bytes.Equal(inp[i:i+len(seq)], seq) {
			return false
		}
	}
	return true
}

func IsValidP2Repeating(id []byte) bool {
	if bytes.HasPrefix(id, []byte("0")) {
		return false
	}

	half := len(id) / 2
	for i := range half {
		if RepeatsSequence(id, id[:i+1]) {
			return false
		}
	}
	return true
}

func InvalidInRange(ran []byte, validator func([]byte) bool) []int {
	l, r, ok := bytes.Cut(ran, []byte("-"))
	if !ok {
		panic(fmt.Errorf("input did not have expected separator: %s", ran))
	}
	li, err := strconv.Atoi(string(l))
	if err != nil {
		panic(fmt.Errorf("input was not a valid number: %s - %w", l, err))
	}
	ri, err := strconv.Atoi(string(r))
	if err != nil {
		panic(fmt.Errorf("input was not a valid number: %s - %w", r, err))
	}
	results := make([]int, 0)
	for i := li; i <= ri; i++ {
		if !validator(fmt.Appendf([]byte{}, "%d", i)) {
			results = append(results, i)
		}
	}
	return results
}
