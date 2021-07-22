package main

import (
	"io/ioutil"
	"log"
	"math"
	"strings"
)

func parseInput(file string) ([]string, []string) {
	var lines, wire1, wire2 []string

	barray, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Unable to read file at %s", file)
	}

	lines = strings.Split(string(barray), "\n")
	wire1 = strings.Split(lines[0], ",")
	wire2 = strings.Split(lines[1], ",")

	return wire1, wire2
}

// From an origin (0, 0), calculates the position by walking the directions given by the input
//
// Returns the *position* after each direction
func CalculatePath(input [][]int) [][]int {
}

// Finds the coordinates where the two result arrays cross
func FindIntersections(result1 [][]int, result2 [][]int) (inters [][]int) {
	for coord1 := range result1 {
		for coord2 := range result2 {
			if coord1[0] == coord2[0] && coord1[1] == coord2[1] {
				result = append(result, coord1)
			}
		}
	}

	return inters
}

// Calculates the shortest distance to a specific point given a variety of intersections
func ShortestDistance(intersections [][]int, poi []int) int {
	var minimum int

	for inter := range intersections {
		dist := math.Abs(inter[0]-poi[0]) + math.Abs(inter[1]-poi[1])

		if dist == 0 {
			continue
		}

		if minimum == 0 {
			minimum = dist // Base case
		} else {
			minimum = math.Min(minimum, dist)
		}
	}

	return minimum
}

func main() {
	input1, input2 := parseInput("../3_input")

	// We want to end up with two arrays of coordinates, one for each wire
	// We want to look for entries that are equivalent
	// Then, we want to look for the one with the shortest Manhattan distance
	result1 := CalculatePath(input1)
	result2 := CalculatePath(input2)

	pointOfInterest := []int{0, 0}
	intersections := FindIntersections(result1, result2)
	shortest := ShortestDistance(intersections, pointOfInterest)

	log.Printf("Shortest distance to %d is %d", pointOfInterest, shortest)

	return
}
