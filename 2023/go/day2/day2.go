package day2

import (
	"os"
	"strconv"
	"strings"
)

var (
	limits = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
)

type Game struct {
	Id    int
	Hands []Hand
	Valid bool
}

type Hand struct {
	Dice map[string]int
}

func (g *Game) Power() int {
	var red, green, blue int

	for _, hand := range g.Hands {
		if hand.Dice["red"] > red {
			red = hand.Dice["red"]
		}
		if hand.Dice["green"] > green {
			green = hand.Dice["green"]
		}
		if hand.Dice["blue"] > blue {
			blue = hand.Dice["blue"]
		}
	}

	return red * green * blue
}

func parseInput(path string) ([]Game, error) {
	games := []Game{}

	contents, err := os.ReadFile(path)
	if err != nil {
		return games, err
	}

	for _, line := range strings.Split(string(contents), "\n") {
		var game Game

		components := strings.Split(line, ": ")
		handsAsStr := strings.Split(components[1], "; ")

		gameIdStr := strings.Split(components[0], " ")
		id, _ := strconv.Atoi(gameIdStr[1])
		game.Id = id
		game.Valid = true // default to true because we'll negate it later

		var hands []Hand
		for _, handStr := range handsAsStr {
			hand := Hand{
				Dice: map[string]int{
					"red":   0,
					"blue":  0,
					"green": 0,
				},
			}

			diceAsStr := strings.Split(handStr, ", ")
			for _, dieAsStr := range diceAsStr {
				pieces := strings.Split(dieAsStr, " ")

				quantity, _ := strconv.Atoi(pieces[0])
				color := pieces[1]

				hand.Dice[color] = quantity
			}

			hands = append(hands, hand)
		}

		game.Hands = hands

		games = append(games, game)
	}

	return games, nil
}

func Part1() (int, error) {
	games, err := parseInput("../inputs/day2")
	if err != nil {
		return 0, err
	}

	var sum int
	for _, game := range games {
		for _, hand := range game.Hands {
			if hand.Dice["red"] > limits["red"] {
				game.Valid = false
			}

			if hand.Dice["green"] > limits["green"] {
				game.Valid = false
			}

			if hand.Dice["blue"] > limits["blue"] {
				game.Valid = false
			}
		}

		if game.Valid {
			sum += game.Id
		}
	}

	return sum, nil
}

func Part2() (int, error) {
	games, err := parseInput("../inputs/day2")
	if err != nil {
		return 0, err
	}

	var sum int
	for _, game := range games {
		sum += game.Power()
	}

	return sum, nil
}
