package day4

import (
	"math"
	"os"
	"strings"
)

type Card struct {
	ResultNumbers []string
	MyNumbers     []string
	Matches       int
	Instances     int
}

func (c *Card) determineMatches() {
	for _, num := range c.MyNumbers {
		if num == "" {
			continue
		}
		for _, result := range c.ResultNumbers {
			if result == num {
				c.Matches++
			}
		}
	}
}

func parseInput(path string) ([]*Card, error) {
	var cards []*Card

	contents, err := os.ReadFile(path)
	if err != nil {
		return cards, err
	}

	for _, line := range strings.Split(string(contents), "\n") {
		var card Card

		parts := strings.Split(line, ": ")
		numSets := strings.Split(parts[1], " | ")
		// This is pretty dodgy because single digit numbers will have a space in front,
		// but this applies equally to results and my numbers, so it's ... kinda fine.
		card.ResultNumbers = strings.Split(numSets[0], " ")
		card.MyNumbers = strings.Split(numSets[1], " ")
		card.determineMatches()

		cards = append(cards, &card)
	}

	return cards, nil
}

func Part1() (int, error) {
	cards, err := parseInput("../inputs/day4")
	if err != nil {
		return -1, err
	}

	var sum int
	for _, card := range cards {
		sum += int(math.Pow(float64(2), float64(card.Matches-1)))
	}

	return sum, nil
}

func calculateInstances(cards []*Card) {
	for id, card := range cards {
		card.Instances += 1
		calculateInstances(cards[id+1 : id+card.Matches+1])
	}
}

func Part2() (int, error) {
	cards, err := parseInput("../inputs/day4")
	if err != nil {
		return -1, err
	}

	calculateInstances(cards)

	var sum int
	for _, card := range cards {
		sum += card.Instances
	}

	return sum, nil
}
