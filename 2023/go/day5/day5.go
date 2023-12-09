package day5

import (
	"errors"
	"math"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type Almanac struct {
	Seeds          []SeedMap
	SeedSoilMap    []ResourceMap
	SoilFertMap    []ResourceMap
	FertWaterMap   []ResourceMap
	WaterLightMap  []ResourceMap
	LightTempMap   []ResourceMap
	TempHumMap     []ResourceMap
	HumLocationMap []ResourceMap
}

type ResourceMap struct {
	DestinationStart int
	SourceStart      int
	Range            int
}

type SeedMap struct {
	Start int
	Range int
}

type SeedMapParser func([]string) ([]SeedMap, error)

func parsePart1(seedsAsStr []string) ([]SeedMap, error) {
	var seedMaps []SeedMap

	for _, seed := range seedsAsStr {
		s, err := strconv.Atoi(seed)
		if err != nil {
			return seedMaps, err
		}

		seedMaps = append(seedMaps, SeedMap{Start: s, Range: 1})
	}

	return seedMaps, nil
}

func parsePart2(seedsAsStr []string) ([]SeedMap, error) {
	var seedMaps []SeedMap

	for i := 0; i < len(seedsAsStr); i += 2 {
		start := seedsAsStr[i]
		s, err := strconv.Atoi(start)
		if err != nil {
			return seedMaps, err
		}

		rangeIdx := seedsAsStr[i+1]
		r, err := strconv.Atoi(rangeIdx)
		if err != nil {
			return seedMaps, err
		}

		seedMaps = append(seedMaps, SeedMap{Start: s, Range: r})
	}

	return seedMaps, nil
}

func parseInput(path string, parser SeedMapParser) (Almanac, error) {
	var almanac Almanac

	contents, err := os.ReadFile(path)
	if err != nil {
		return almanac, err
	}

	sections := strings.Split(string(contents), "\n\n")
	seedsAsStr := strings.Split(strings.Split(sections[0], "seeds: ")[1], " ")

	almanac.Seeds, err = parser(seedsAsStr)
	if err != nil {
		return almanac, err
	}

	for idx, section := range sections {
		if idx == 0 {
			continue
		}

		resourceMaps, err := processMap(section)
		if err != nil {
			return almanac, err
		}

		switch idx {
		case 1:
			almanac.SeedSoilMap = resourceMaps
		case 2:
			almanac.SoilFertMap = resourceMaps
		case 3:
			almanac.FertWaterMap = resourceMaps
		case 4:
			almanac.WaterLightMap = resourceMaps
		case 5:
			almanac.LightTempMap = resourceMaps
		case 6:
			almanac.TempHumMap = resourceMaps
		case 7:
			almanac.HumLocationMap = resourceMaps
		default:
			return almanac, errors.New("invalid index for resource map")
		}
	}

	return almanac, nil
}

func processMap(rawMap string) ([]ResourceMap, error) {
	var resourceMaps []ResourceMap

	for idx, raw := range strings.Split(rawMap, "\n") {
		if idx == 0 {
			continue
		}

		var splits []int

		for _, n := range strings.Split(raw, " ") {
			val, err := strconv.Atoi(n)
			if err != nil {
				return resourceMaps, err
			}
			splits = append(splits, val)
		}

		resourceMap := ResourceMap{
			DestinationStart: splits[0],
			SourceStart:      splits[1],
			Range:            splits[2],
		}
		resourceMaps = append(resourceMaps, resourceMap)
	}

	return resourceMaps, nil
}

// Given an integer value which is a destination number and a map from which it was derived,
// calculate what the source number must have been.
func calculateSource(value int, resourceMaps []ResourceMap) int {
	for _, m := range resourceMaps {
		if value >= m.DestinationStart && value < m.DestinationStart+m.Range {
			return m.SourceStart + value - m.DestinationStart
		}
	}

	// Return the value itself if we didn't find anything
	return value
}

func (a *Almanac) LowestLocationSeed() int {
	locations := a.HumLocationMap

	// In these maps, locations are the destination, so we want to sort by that
	slices.SortFunc(locations, func(a, b ResourceMap) int {
		return a.DestinationStart - b.DestinationStart
	})

	for i := 0; i < math.MaxUint32; i++ {
		humidity := calculateSource(i, a.HumLocationMap)
		temp := calculateSource(humidity, a.TempHumMap)
		light := calculateSource(temp, a.LightTempMap)
		water := calculateSource(light, a.WaterLightMap)
		fert := calculateSource(water, a.FertWaterMap)
		soil := calculateSource(fert, a.SoilFertMap)
		seed := calculateSource(soil, a.SeedSoilMap)

		//fmt.Printf("%d <- %d <- %d <- %d <- %d <- %d <- %d <- %d\n", seed, soil, fert, water, light, temp, humidity, i)

		// Check if that seed matches one of the originals - if so, we're done
		for _, m := range a.Seeds {
			if seed >= m.Start && seed < m.Start+m.Range {
				return i
			}
		}
	}

	return -1
}

func Part1() (int, error) {
	almanac, err := parseInput("../inputs/day5", parsePart1)
	if err != nil {
		return -1, err
	}

	return almanac.LowestLocationSeed(), nil
}

func Part2() (int, error) {
	almanac, err := parseInput("../inputs/day5", parsePart2)
	if err != nil {
		return -1, err
	}

	return almanac.LowestLocationSeed(), nil
}
