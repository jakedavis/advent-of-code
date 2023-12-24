package day19

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const debug = false

var (
	errInvalidCategory = errors.New("invalid category")
	errInvalidOperator = errors.New("invalid operator")
)

func parseInput(path string) (Workflows, Parts, error) {
	workflows := make(map[string]Rules)

	var parts Parts

	contents, err := os.ReadFile(path)
	if err != nil {
		return workflows, parts, err
	}

	splits := strings.Split(string(contents), "\n\n")

	var scanner *bufio.Scanner

	// Process workflows
	scanner = bufio.NewScanner(strings.NewReader(splits[0]))
	for scanner.Scan() {
		identifier, rules, err := parseWorkflow(scanner.Text())
		if err != nil {
			return workflows, parts, err
		}
		workflows[identifier] = rules
	}

	// Process parts
	scanner = bufio.NewScanner(strings.NewReader(splits[1]))
	for scanner.Scan() {
		part, err := parsePart(scanner.Text())
		if err != nil {
			return workflows, parts, err
		}
		parts = append(parts, part)
	}

	return workflows, parts, nil
}

func Part1() (int, error) {
	workflows, parts, err := parseInput("../inputs/day10")
	if err != nil {
		return 0, err
	}

	if debug {
		for name, w := range workflows {
			fmt.Printf("%-3s %v\n", name, w)
		}
		for _, p := range parts {
			fmt.Println(p)
		}
	}

	var sum int
	for _, part := range parts {
		if workflows.Accept(part) {
			sum += part.Value()
		}
	}

	return sum, nil
}

func Part2() (int, error) {
	workflows, _, err := parseInput("../inputs/day10")
	if err != nil {
		return 0, err
	}

	if debug {
		for name, w := range workflows {
			fmt.Printf("%-3s %v\n", name, w)
		}
	}

	var sum int

	// I don't understand this prompt

	return sum, nil
}
