package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

const numbers = "0123456789"

var (
	numberStrings = map[string]int{
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"0":     0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"zero":  0,
	}
)

func input(path string) ([]string, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return []string{}, err
	}

	splits := strings.Split(string(contents), "\n")

	return splits, nil
}

func Part1() (int, error) {
	lines, err := input("../inputs/day1")
	if err != nil {
		return -1, err
	}

	var valueChan chan int = make(chan int, len(lines))
	var wg sync.WaitGroup
	var sum int

	for _, line := range lines {
		wg.Add(1)

		go func(line string) {
			defer wg.Done()
			var first, last string

			for _, i := range line {
				if strings.Contains(numbers, string(i)) {
					if first == "" {
						first = string(i)
					}
					last = string(i)
				}
			}

			value, _ := strconv.Atoi(first + last)

			valueChan <- value
		}(line)
	}

	wg.Wait()
	close(valueChan)

	for msg := range valueChan {
		sum += msg
	}

	return sum, nil
}

func Part2() (int, error) {
	lines, err := input("../inputs/day1")
	if err != nil {
		return -1, err
	}

	sum := 0

	for _, line := range lines {
		// Ensure all valid indices compare favorably to these base numbers
		firstIdx := len(line)
		lastIdx := -1

		var first, last int

		for k, v := range numberStrings {
			var idx int

			idx = strings.Index(line, k)
			if idx != -1 {
				if idx < firstIdx {
					firstIdx = idx
					first = v
				}
			}

			idx = strings.LastIndex(line, k)
			if idx != -1 {
				if idx > lastIdx {
					lastIdx = idx
					last = v
				}
			}
		}

		value, _ := strconv.Atoi(fmt.Sprint(first) + fmt.Sprint(last))
		sum += value
	}

	return sum, nil
}
