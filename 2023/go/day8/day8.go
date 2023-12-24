package day8

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"golang.org/x/exp/slices"
)

// For each start:
// 		Get to Z
// 		Loop once and calculate steps -> this is the period
//
// LCM on all periods

var debug = false

type Instructions []string

type Network struct {
	Entrypoint string
	Nodes      NodeMap
}

type NodeMap map[string]Node

type Node struct {
	Left  string
	Right string
}

func processLine(line string) (string, Node) {
	var node Node

	parts := strings.Split(line, " = ")

	key := parts[0]

	nodeParts := strings.Split(parts[1], ", ")

	node.Left = strings.TrimLeft(nodeParts[0], "(")
	node.Right = strings.TrimRight(nodeParts[1], ")")

	return key, node
}

func parseInput(path string) (Instructions, Network, error) {
	var instructions Instructions
	var network Network
	network.Nodes = make(NodeMap)

	contents, err := os.ReadFile(path)
	if err != nil {
		return instructions, network, err
	}

	splits := strings.Split(string(contents), "\n\n")
	instructions = strings.Split(splits[0], "")

	nodesRaw := strings.Split(splits[1], "\n")
	network.Entrypoint = "AAA"

	for _, nodeRaw := range nodesRaw {
		key, node := processLine(nodeRaw)
		network.Nodes[key] = node
	}

	return instructions, network, nil
}

func (n *Network) Starts() []string {
	var starts []string

	for key := range n.Nodes {
		if strings.Split(key, "")[2] == "A" {
			starts = append(starts, key)
		}
	}
	fmt.Println()

	return starts
}

func (n *Network) Destinations() []string {
	var destinations []string

	for key := range n.Nodes {
		if strings.Split(key, "")[2] == "Z" {
			destinations = append(destinations, key)
		}
	}

	return destinations
}

func (n *Network) Traverse(starts []string, destinations []string, instructions Instructions) int {
	var steps int

	currentNodes := starts

	for {
		for _, i := range instructions {
			var nextNodes []string
			var matches int

			for _, currentNode := range currentNodes {
				switch i {
				case "L":
					nextNodes = append(nextNodes, n.Nodes[currentNode].Left)
				case "R":
					nextNodes = append(nextNodes, n.Nodes[currentNode].Right)
				default:
					panic(errors.New("Invalid instruction"))
				}

				if slices.Contains(destinations, currentNode) {
					matches++
				}
			}

			if debug {
				fmt.Printf("[%.6d] %v\n", steps, currentNodes)
			}

			if matches == len(destinations) {
				return steps
			}

			steps += 1

			if steps%10_000_000 == 0 {
				fmt.Println(steps)
			}

			currentNodes = nextNodes
		}
	}
}

func (n *Network) Period(start string, destinations []string, instructions Instructions) int {
	var steps int

	currentNode := start
	for {
		for _, i := range instructions {
			steps += 1

			switch i {
			case "L":
				currentNode = n.Nodes[currentNode].Left
			case "R":
				currentNode = n.Nodes[currentNode].Right
			}

			for _, destination := range destinations {
				if currentNode == destination {
					return steps
				}
			}
		}
	}
}

func (n *Network) Offset(start string, destinations []string, instructions Instructions) (int, string) {
	var steps int

	currentNode := start
	for {
		for _, i := range instructions {
			steps += 1

			switch i {
			case "L":
				currentNode = n.Nodes[currentNode].Left
			case "R":
				currentNode = n.Nodes[currentNode].Right
			}

			for _, destination := range destinations {
				if currentNode == destination {
					return steps, destination
				}
			}
		}
	}
}

func Part1() (int, error) {
	var steps int

	instructions, network, err := parseInput("../inputs/day8")
	if err != nil {
		return steps, err
	}

	starts := []string{"AAA"}
	destinations := []string{"ZZZ"}

	steps = network.Traverse(starts, destinations, instructions)

	return steps, nil
}

func Part2() (int, error) {
	var steps int

	instructions, network, err := parseInput("../inputs/day8_test_p2")
	if err != nil {
		return steps, err
	}

	starts := network.Starts()
	destinations := network.Destinations()

	debug = false
	if debug {
		fmt.Println()
		fmt.Printf("Starts: %v\n", starts)
		fmt.Printf("Destinations: %v\n", destinations)
		fmt.Println()
		time.Sleep(5 * time.Second)
	}

	for _, s := range starts {
		offset, _ := network.Offset(s, destinations, instructions)
		fmt.Printf("[%s] %d\n", s, offset)
	}

	for i := 0; i < 10; i++ {
		var lcm int

		for _, d := range destinations {
			period := network.Period(d, destinations, instructions)

			if i%period == 0 {
				lcm = i
			} else {
				lcm = 0
				break
			}

			fmt.Printf("[%s] %d\n", d, period)
		}

		if lcm > 0 {
			fmt.Println(lcm)
			break
		}
	}

	// steps = network.Traverse(starts, destinations, instructions)

	return steps, nil
}
