// 1_launch.go

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("../1_input")
	if err != nil {
		log.Fatal("Couldn't read input")
	}

	totalFuel := 0
	splits := strings.Split(string(input), "\r\n")
	for _, mass := range splits {
		mass, err := strconv.Atoi(mass)
		if err != nil {
			log.Fatalf("Error converting mass to int: %d, %s\n", mass, err)
		}

		fuel := CalculateFuel(mass)
		totalFuel += fuel
	}

	fmt.Printf("Total fuel required: %d units\n", totalFuel)
}

// CalculateFuel is a recursive function that calculates the amount of fuel required for the given
// mass, plus any "sub-masses", in other words the fuel required for the fuel.
func CalculateFuel(mass int) int {
	fuel := int(math.Floor(float64(mass/3)) - 2)
	if fuel <= 0 {
		return 0
	}

	fuel += CalculateFuel(fuel)
	return fuel
}
