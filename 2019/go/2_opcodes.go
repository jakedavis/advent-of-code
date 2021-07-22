package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func parseInput(file string) []int {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("Couldn't read input file")
	}

	strArray := strings.Split(strings.TrimSuffix(string(f), "\n"), ",")
	intArray := []int{}
	for k, i := range strArray {
		j, err := strconv.Atoi(i)
		if err != nil {
			log.Printf("Unable to parse int at index %d: %d", k, j)
		}
		intArray = append(intArray, j)
	}

	return intArray
}

func Add(args []int) int {
	return args[0] + args[1]
}

func Multiply(args []int) int {
	return args[0] * args[1]
}

func Execute(opcode int, args []int) int {
	// log.Printf("Opcode: %d, Arguments: %d", opcode, args)

	if opcode == 1 {
		return Add(args)
	} else if opcode == 2 {
		return Multiply(args)
	} else if opcode == 99 {
		return 0
	} else {
		return -1
	}
}

func GenerateResult(opcodes []int) int {
	var instruction, opcode, result int
	var args []int

	for {
		opcode = opcodes[instruction]
		args = []int{
			opcodes[opcodes[instruction+1]],
			opcodes[opcodes[instruction+2]],
		}
		result = Execute(opcode, args)
		opcodes[opcodes[instruction+3]] = result

		switch result {
		case -1:
			log.Fatalf("Bad opcode or result: %s %s", opcode, args)
		case 0:
			// log.Printf("Received exit condition at index %d", instruction)
			return opcodes[0]
		default:
			instruction += 4
		}
	}
}

func main() {
	opcodes := parseInput("../2_input")

	// opcodes[1] = 12
	// opcodes[2] = 2

	for i := 1; i < 128; i++ {
		for j := 1; j < 128; j++ {
			opcodes[1] = i
			opcodes[2] = j
			result := GenerateResult(opcodes)

			// log.Printf("i=%d j=%d result=%d", i, j, result)

			if result == 19690720 {
				// log.Printf("%d", opcodes)
				log.Printf("Received expected result at i=%d j=%d", i, j)
				log.Printf("100*i*j=%d", 100*i+j)
				return
			}

			opcodes = parseInput("../2_input")
		}
	}
}
