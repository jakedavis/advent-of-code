package day6

import (
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/jakedavis/aoc/2023/util"
)

type Race struct {
	Time     int
	Distance int
}

func (r *Race) DetermineNumberOfWinners() int {
	var winners int

	for i := 0; i < r.Time; i++ {
		distance := i * (r.Time - i)
		if distance > r.Distance {
			winners++
		}
	}

	return winners
}

func preprocessInput(path string) ([]string, []string, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return []string{}, []string{}, err
	}

	splits := strings.Split(string(contents), util.WindowsNewline)
	timesRaw := splits[0]

	timesRaw = strings.TrimPrefix(timesRaw, "Time:")
	timesRaw = strings.TrimSpace(timesRaw)
	times := regexp.MustCompile(`\s+`).Split(timesRaw, -1)

	distancesRaw := splits[1]
	distancesRaw = strings.TrimPrefix(distancesRaw, "Distance:")
	distancesRaw = strings.TrimSpace(distancesRaw)
	distances := regexp.MustCompile(`\s+`).Split(distancesRaw, -1)

	if len(times) != len(distances) {
		return times, distances, errors.New("times and distances are of uneven length")
	}

	return times, distances, nil
}

func parseIndividually(times []string, distances []string) ([]Race, error) {
	var races []Race

	for i := 0; i < len(times); i++ {
		t, err := strconv.Atoi(times[i])
		if err != nil {
			return races, err
		}

		d, err := strconv.Atoi(distances[i])
		if err != nil {
			return races, err
		}

		race := Race{
			Time:     t,
			Distance: d,
		}

		races = append(races, race)
	}

	return races, nil
}

func parseAsOne(times []string, distances []string) ([]Race, error) {
	var races []Race

	var timeStr string
	var distanceStr string
	for i := 0; i < len(times); i++ {
		timeStr += times[i]
		distanceStr += distances[i]
	}

	time, err := strconv.Atoi(timeStr)
	if err != nil {
		return races, err
	}
	distance, err := strconv.Atoi(distanceStr)
	if err != nil {
		return races, err
	}

	races = append(races, Race{Time: time, Distance: distance})

	return races, nil
}

type ParseFunc func([]string, []string) ([]Race, error)

func parseInput(path string, fn ParseFunc) ([]Race, error) {
	var races []Race

	times, distances, err := preprocessInput(path)
	if err != nil {
		return races, err
	}

	races, err = fn(times, distances)
	if err != nil {
		return races, err
	}

	return races, nil
}

func Part1() (int, error) {
	races, err := parseInput("../inputs/day6", parseIndividually)
	if err != nil {
		return 0, err
	}

	ways := 1
	for _, race := range races {
		ways *= race.DetermineNumberOfWinners()
	}

	return ways, nil
}

func Part2() (int, error) {
	races, err := parseInput("../inputs/day6", parseAsOne)
	if err != nil {
		return 0, err
	}

	ways := 1
	for _, race := range races {
		ways *= race.DetermineNumberOfWinners()
	}

	return ways, nil
}
