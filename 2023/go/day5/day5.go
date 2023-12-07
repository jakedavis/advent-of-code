package day5

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Almanac struct {
	Seeds          Seeds
	SeedSoilMap    ResourceMap
	SoilFertMap    ResourceMap
	FertWaterMap   ResourceMap
	WaterLightMap  ResourceMap
	LightTempMap   ResourceMap
	TempHumMap     ResourceMap
	HumLocationMap ResourceMap
}

type Seeds []int
type ResourceMapRaw struct {
	DestinationStart int
	SourceStart      int
	Range            int
}

type ResourceMap map[int]int

func parseInput(path string) (Almanac, error) {
	var almanac Almanac

	contents, err := os.ReadFile(path)
	if err != nil {
		return almanac, err
	}

	sections := strings.Split(string(contents), "\n\n")

	seedsAsStr := strings.Split(strings.Split(sections[0], "seeds: ")[1], " ")
	var seeds Seeds
	for _, seedStr := range seedsAsStr {
		seed, err := strconv.Atoi(seedStr)
		if err != nil {
			return almanac, err
		}

		seeds = append(seeds, seed)
	}
	almanac.Seeds = seeds

	for idx, section := range sections {
		if idx == 0 {
			continue
		}

		resourceMap, err := processRawMap(section)
		if err != nil {
			return almanac, err
		}
		switch idx {
		case 1:
			almanac.SeedSoilMap = resourceMap
		case 2:
			almanac.SoilFertMap = resourceMap
		case 3:
			almanac.FertWaterMap = resourceMap
		case 4:
			almanac.WaterLightMap = resourceMap
		case 5:
			almanac.LightTempMap = resourceMap
		case 6:
			almanac.TempHumMap = resourceMap
		case 7:
			almanac.HumLocationMap = resourceMap
		default:
			return almanac, errors.New("invalid index for resource map")
		}
	}

	return almanac, nil
}

func processRawMap(rawMap string) (ResourceMap, error) {
	resourceMap := make(ResourceMap)

	for i := 0; i < math.MaxInt16; i++ {
		resourceMap[i] = i
	}

	for idx, raw := range strings.Split(rawMap, "\n") {
		if idx == 0 {
			continue
		}

		var splits []int

		for _, n := range strings.Split(raw, " ") {
			val, err := strconv.Atoi(n)
			if err != nil {
				return resourceMap, err
			}
			splits = append(splits, val)
		}

		for i := 0; i < splits[2]; i++ {
			resourceMap[splits[1]+i] = splits[0] + i
		}
	}
	fmt.Println(resourceMap)
	fmt.Println()

	return resourceMap, nil
}

func (a *Almanac) LowestLocation() int {
	var location int = math.MaxInt64

	for _, seed := range a.Seeds {
		soil := a.SeedSoilMap[seed]
		fert := a.SoilFertMap[soil]
		water := a.FertWaterMap[fert]
		light := a.WaterLightMap[water]
		temp := a.LightTempMap[light]
		humidity := a.TempHumMap[temp]
		loc := a.HumLocationMap[humidity]

		fmt.Printf("%d -> %d -> %d -> %d -> %d -> %d -> %d -> %d\n", seed, soil, fert, water, light, temp, humidity, loc)

		if loc < location {
			location = loc
		}
	}

	return location
}

func Part1() (int, error) {
	almanac, err := parseInput("../inputs/day5_test")
	if err != nil {
		return -1, err
	}

	return almanac.LowestLocation(), nil
}

func Part2() (int, error) {
	return 0, nil
}
