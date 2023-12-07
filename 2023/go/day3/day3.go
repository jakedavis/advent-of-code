package day3

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Grid [][]string
type Coordinate struct {
	X int
	Y int
}

// Numeric represents multiple array elements comprising a single number
type Numeric struct {
	// Start coordinate
	Coordinate Coordinate
	// Length of the number (in x positions)
	Length int
}

const digits = "0123456789"
const symbols = "/#@=$%-*&+"

var (
	adjacents = []Coordinate{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
)

func parseInput(path string) (Grid, error) {
	var grid Grid

	contents, err := os.ReadFile(path)
	if err != nil {
		return grid, err
	}

	ySplits := strings.Split(string(contents), "\n")
	grid = make([][]string, len(ySplits))
	for y, line := range ySplits {
		grid[y] = make([]string, len([]rune(line)))
		for x, char := range []rune(line) {
			grid[y][x] = string(char)
		}
	}

	return grid, nil
}

func (g Grid) adjacentToSymbol(number Numeric) bool {
	for i := 0; i < number.Length; i++ {
		for _, coordinate := range adjacents {
			y := number.Coordinate.Y + coordinate.Y
			x := number.Coordinate.X + coordinate.X + i

			if y < 0 || x < 0 {
				continue
			}

			if y > len(g)-1 || x > len(g[0])-1 {
				continue
			}

			if strings.Contains(symbols, g[y][x]) {
				return true
			}
		}
	}

	return false
}

func Part1() (int, error) {
	grid, err := parseInput("../inputs/day3_p1")
	if err != nil {
		return 0, err
	}

	valueChan := make(chan int, len(grid)*len(grid[0]))

	for y, row := range grid {
		var skip int

		selections := []int{}
		for x, n := range row {
			if skip > 0 {
				skip--
				continue
			}

			if !strings.Contains(digits, n) {
				continue
			}

			// Traverse the row until it's done or we hit a non-number
			var num Numeric
			num.Coordinate = Coordinate{X: x, Y: y}
			for i := x; i < len(row); i++ {
				if !strings.Contains(digits, grid[y][i]) {
					break
				}
				num.Length += 1
			}

			// Look for any symbol in an adjacent spot to any of the relevant cells
			if !grid.adjacentToSymbol(num) {
				// If we don't find one, throw it away
				continue
			}

			// If there was a symbol, create a number out of the entry and add it to the sum
			var digits []string
			for i := 0; i < num.Length; i++ {
				digit := grid[num.Coordinate.Y][num.Coordinate.X+i]
				digits = append(digits, digit)
			}
			value, err := strconv.Atoi(strings.Join(digits, ""))
			if err != nil {
				return 0, err
			}

			//fmt.Printf("%.3v %d -> %.3d\n", num.Coordinate, num.Length, value)
			valueChan <- value
			selections = append(selections, value)

			// This allows us to skip the numbers we've already seen
			skip = num.Length
		}
		fmt.Printf("[%.3d] %v\n", y, selections)
	}

	close(valueChan)

	var sum int
	for msg := range valueChan {
		sum += msg
	}

	return sum, nil
}

func Part2() (int, error) {
	grid, err := parseInput("../inputs/day3_p1")
	if err != nil {
		return 0, err
	}

	valueChan := make(chan int, len(grid)*len(grid[0]))

	for y, row := range grid {
		for x, n := range row {
			if n != "*" {
				continue
			}

			ratio := 1
			adjacentNumbers := func(coord Coordinate) []int {
				var nums []int
				var numeric Numeric

				var digits []string
				for i := 0; i < numeric.Length; i++ {
					digit := grid[numeric.Coordinate.Y][numeric.Coordinate.X+i]
					digits = append(digits, digit)
				}
				num, err := strconv.Atoi(strings.Join(digits, ""))
				if err != nil {
					panic(err)
				}
				nums = append(nums, num)

				return nums
			}

			for _, adjacent := range adjacentNumbers(Coordinate{X: x, Y: y}) {
				ratio *= adjacent
			}

			valueChan <- ratio
		}
	}

	close(valueChan)

	var sum int
	for msg := range valueChan {
		sum += msg
	}

	return sum, nil
}
