package main

import (
	"fmt"

	"github.com/jakedavis/aoc/2023/day2"
)

func main() {
	var result int
	var err error

	result, err = day2.Part1()
	if err != nil {
		panic(err)
	}

	fmt.Printf("[Day2][Part1] %d\n", result)

	result, err = day2.Part2()
	if err != nil {
		panic(err)
	}

	fmt.Printf("[Day2][Part2] %d\n", result)
}
