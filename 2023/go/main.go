package main

import (
	"fmt"

	"github.com/jakedavis/aoc/2023/day5"
)

func main() {
	var result int
	var err error

	result, err = day5.Part1()
	if err != nil {
		panic(err)
	}

	fmt.Printf("[Day5][Part1] %d\n", result)
}
