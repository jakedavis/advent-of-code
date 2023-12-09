package day7

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

const debug = false

var useJokers = false

const (
	//go:generate stringer -type=Card
	Joker Card = iota + 1 // 1
	Two                   // 2
	Three                 // 3
	Four                  // 4
	Five                  // 5
	Six                   // 6
	Seven                 // 7
	Eight                 // 8
	Nine                  // 9
	Ten                   // 10
	Jack                  // 11
	Queen                 // 12
	King                  // 13
	Ace                   // 14
)

const (
	//go:generate stringer -type=Type
	HighCard     Type = iota + 1 // 1
	OnePair                      // 2
	TwoPair                      // 3
	ThreeOfAKind                 // 4
	FullHouse                    // 5
	FourOfAKind                  // 6
	FiveOfAKind                  // 7
)

type Type int
type Card int
type Hand []Card
type Play struct {
	Hand Hand
	Bid  int
}

func (h Hand) Type() Type {
	cards := map[Card]int{}

	for _, card := range h {
		cards[card] += 1
	}

	if useJokers {
		jokers := cards[Joker]
		for k := range cards {
			cards[k] += jokers
		}

		delete(cards, Joker)
	}

	// Since keys are unique card values, we can use the unique number plus some minor logic
	// to determine the value of the hand.
	switch len(cards) {
	case 5:
		return HighCard
	case 4:
		return OnePair
	case 3:
		for _, amt := range cards {
			if amt == 3 {
				return ThreeOfAKind
			}
		}

		return TwoPair
	case 2:
		for _, amt := range cards {
			if amt == 4 {
				return FourOfAKind
			}
		}

		return FullHouse
	case 1:
		return FiveOfAKind
	case 0:
		// All Jokers
		return FiveOfAKind
	default:
		panic(errors.New("impossile number of cards??"))
	}
}

func parseInput(path string, jValue Card) ([]Play, error) {
	var plays []Play

	contents, err := os.ReadFile(path)
	if err != nil {
		return plays, err
	}

	for _, line := range strings.Split(string(contents), "\n") {
		var hand Hand

		parts := strings.Split(line, " ")
		handParts := strings.Split(parts[0], "")
		for _, p := range handParts {
			var value Card
			switch p {
			case "2":
				value = Two
			case "3":
				value = Three
			case "4":
				value = Four
			case "5":
				value = Five
			case "6":
				value = Six
			case "7":
				value = Seven
			case "8":
				value = Eight
			case "9":
				value = Nine
			case "T":
				value = Ten
			case "J":
				value = jValue
			case "Q":
				value = Queen
			case "K":
				value = King
			case "A":
				value = Ace
			default:
				return plays, fmt.Errorf("invalid value %v for card", p)
			}

			hand = append(hand, value)
		}

		bid, err := strconv.Atoi(parts[1])
		if err != nil {
			return plays, err
		}

		plays = append(plays, Play{Bid: bid, Hand: hand})
	}

	return plays, nil
}

// Given two hands, determine which is stronger by iterating over each value in the first hand
// and comparing it to the second. This assumes two hands of equal strength.
func (a Hand) Compare(b Hand) int {
	for i, c := range a {
		if c == b[i] {
			continue
		}

		// Not equal, so just return whoever is greater
		return int(c - b[i])
	}

	// If all cards are equal.
	return 0
}

func rankPlays(plays []Play) []Play {
	slices.SortFunc(plays, func(a, b Play) int {
		comp := a.Hand.Type() - b.Hand.Type()
		if comp == 0 {
			return a.Hand.Compare(b.Hand)
		}

		return int(comp)
	})

	return plays
}

func calculateWinnings(plays []Play) int {
	var winnings int
	for rank, play := range plays {
		if debug {
			fmt.Printf("[%d] %v => %v\n", rank+1, play.Hand, play.Hand.Type())
		}

		winnings += play.Bid * (rank + 1)
	}
	return winnings
}

func Part1() (int, error) {
	plays, err := parseInput("../inputs/day7", Jack)
	if err != nil {
		return 0, err
	}

	useJokers = false
	plays = rankPlays(plays)
	winnings := calculateWinnings(plays)

	return winnings, nil
}

func Part2() (int, error) {
	plays, err := parseInput("../inputs/day7", Joker)
	if err != nil {
		return 0, err
	}

	useJokers = true
	plays = rankPlays(plays)
	winnings := calculateWinnings(plays)

	return winnings, nil
}
