package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const Answer = 19690720

func parseInput(file string) []int {
	contents, err := os.ReadFile(file)
	if err != nil {
		panic("file read error")
	}

	parsed := strings.TrimSpace(string(contents))
	splits := strings.Split(parsed, ",")

	ints := make([]int, 0)
	for _, s := range splits {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic("int parse error")
		}

		ints = append(ints, i)
	}

	return ints
}

func part1(input []int) int {
	for i := 0; i < len(input); i += 4 {
		switch input[i] {
		case 1:
			a := input[i+1]
			b := input[i+2]
			r := input[i+3]
			input[r] = input[a] + input[b]
		case 2:
			a := input[i+1]
			b := input[i+2]
			r := input[i+3]
			input[r] = input[a] * input[b]
		case 99:
			return input[0]
		default:
			fmt.Println("Bad opcode:", input[i])
			os.Exit(1)
		}
	}

	return 0
}

func part2(input []int) (int, error) {
	mem := make([]int, len(input), len(input))

	for idx, _ := range make([]int, 99, 99) {
		for idy, _ := range make([]int, 99, 99) {
			copy(mem, input)
			mem[1] = idx
			mem[2] = idy

			if part1(mem) == Answer {
				return idx*100 + idy, nil
			}
		}
	}

	return 0, errors.New("No possibilities")
}

func main() {
	input := parseInput("../2_input")

	alarm := part1(input)
	fmt.Println("[2.1] alarm state =", alarm)

	input = parseInput("../2_input")
	m, err := part2(input)
	if err != nil {
		fmt.Printf("[2.2] err = %s", err)
		os.Exit(1)
	}

	fmt.Printf("[2.2] value = %d", m)
}
