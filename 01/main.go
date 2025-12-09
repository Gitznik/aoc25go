package main

import (
	"fmt"

	util "github.com/gitznik/aoc/pkg"
)

func main() {
	inpd1 := util.ReadInput(util.POne)
	fmt.Println("Running day 1")
	d1(inpd1)

	fmt.Println("Running day 2")
	d2(inpd1)
}
