package day9

import (
	"os"
	"strconv"
	"strings"
)

type Sequence []int

func (s Sequence) zeros() bool {
	for _, num := range s {
		if num != 0 {
			return false
		}
	}

	return true
}

func (s Sequence) lineage() []Sequence {
	var lineage []Sequence

	lineage = append(lineage, s)

	for {
		next := s.next()
		lineage = append(lineage, next)

		if next.zeros() {
			break
		}

		s = next
	}

	return lineage
}

func (s Sequence) extrapolateLeft() int {
	lineage := s.lineage()
	lineage[len(lineage)-1] = Sequence{0, 0}

	for i := len(lineage) - 1; i > 0; i-- {
		previous := lineage[i-1]
		current := lineage[i]

		// We're generating values to the left. In the base case, we added an additional 0.
		// Afterward, we just use the 0th element to create the new 0th element. You never
		// actually need more values than that.
		previous = Sequence{previous[0] - current[0], previous[0]}

		// Finally, overwrite the previous value
		lineage[i-1] = previous
	}

	// We only end up caring about the first element of the first sequence
	first := lineage[0]

	// The new first element of the first sequence is the backwards-in-time element
	return first[0]
}

func (s Sequence) extrapolateRight() int {
	// Put the first sequence onto the lineage
	lineage := s.lineage()

	for i := len(lineage) - 1; i > 0; i-- {
		previous := lineage[i-1]
		current := lineage[i]

		// The "new" value is the addition of the last values of the current and previous sequences
		previous = append(previous, previous[len(previous)-1]+current[len(current)-1])

		// Overwrite the array
		lineage[i-1] = previous
	}

	// We only end up caring about the first sequence
	first := lineage[0]

	// The last element of the first sequence is the next element in the history
	return first[len(first)-1]
}

func (s Sequence) next() Sequence {
	var nextSequence []int

	for _, n := range s {
		if n != 0 {
			break
		}
	}

	for i := 0; i < len(s)-1; i++ {
		diff := float64(s[i+1]) - float64(s[i])
		nextSequence = append(nextSequence, int(diff))
	}

	return nextSequence
}

func parseInput(path string) ([]Sequence, error) {
	var sequences []Sequence

	contents, err := os.ReadFile(path)
	if err != nil {
		return sequences, err
	}

	for _, line := range strings.Split(string(contents), "\n") {
		var sequence Sequence
		for _, value := range strings.Split(line, " ") {
			num, err := strconv.Atoi(value)
			if err != nil {
				return sequences, err
			}

			sequence = append(sequence, num)
		}
		sequences = append(sequences, sequence)
	}

	return sequences, nil
}

func Part1() (int, error) {
	var sum int

	sequences, err := parseInput("../inputs/day9")
	if err != nil {
		return sum, err
	}

	for _, sequence := range sequences {
		sum += sequence.extrapolateRight()
	}

	return sum, nil
}

func Part2() (int, error) {
	var sum int

	sequences, err := parseInput("../inputs/day9")
	if err != nil {
		return sum, err
	}

	for _, sequence := range sequences {
		sum += sequence.extrapolateLeft()
	}

	return sum, nil
}
