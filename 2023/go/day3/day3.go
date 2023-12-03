package day3

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Grid [][]string
type Coordinate []int

// Numeric represents multiple array elements comprising a single number
type Numeric struct {
	// Start coordinate
	Coordinate Coordinate
	// Length of the number (in x positions)
	Length int
}

const digits = "0123456789"
const symbols = "!@#$%^&*()-+/="

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
			y := number.Coordinate[0] + coordinate[0]
			x := number.Coordinate[1] + coordinate[1]

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
	grid, err := parseInput("../inputs/day3_p1_test")
	if err != nil {
		return 0, err
	}

	valueChan := make(chan int, len(grid)*len(grid[0]))

	for y, row := range grid {
		for x, n := range row {
			if !strings.Contains(digits, n) {
				continue
			}

			// 1. once isDigit is true, we want to traverse the row until it's done or we hit a non-number
			//   -> that creates a Numeric
			coord := Coordinate{y, x}

			var num Numeric
			num.Coordinate = coord
			for i := x; i < len(row)-1; i++ {
				if !strings.Contains(digits, grid[y][i]) {
					break
				}
				num.Length += 1
			}

			for i := x; i > 0; i-- {
				if !strings.Contains(digits, grid[y][i]) {
					break
				}
				num.Length -= 1
			}
			if num.Length < 0 {
				num.Coordinate = Coordinate{y, num.Coordinate[1] + num.Length}
				num.Length = int(math.Abs(float64(num.Length)))
			}

			// 2. from the numeric, we can determine all the adjacent spots
			// 	 technically we can just look anywhere but we could ignore the other numbers in theory
			// 3. look for any symbol in an adjacent spot
			if !grid.adjacentToSymbol(num) {
				// 4. if one does not exist, we throw it away
				continue
			}

			// 5. if one exists, strconv the numeric and pass off to channel
			// TODO strconv.Atoi(coord)
			var digits []string
			for i := 0; i < num.Length; i++ {
				digit := grid[num.Coordinate[0]][num.Coordinate[1]+i]
				digits = append(digits, digit)
			}
			value, err := strconv.Atoi(strings.Join(digits, ""))
			if err != nil {
				return 0, err
			}

			fmt.Printf("candidate: %v -> %d\n", num, value)

			valueChan <- value
			// 6. optimally, jump the x to the end of the number+1 index
		}
	}

	close(valueChan)

	var sum int
	for msg := range valueChan {
		sum += msg
	}

	return sum, nil
}

func Part2() (int, error) {
	return 0, nil
}
