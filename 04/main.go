package main

import (
	"fmt"

	util "github.com/gitznik/aoc/pkg"
)

func main() {
	inpd1 := util.ReadInput(util.POne)
	fmt.Println("Running part 1")
	p1(inpd1)

	fmt.Println("Running part 2")
	p2(inpd1)
}
