package day19

import "fmt"

type Rule struct {
	destination string
	category    string
	operator    string
	threshold   int
	condition   func(value int) bool
}

func (r Rule) String() string {
	return fmt.Sprintf("{%s%s%4d ? %-3s}", r.category, r.operator, r.threshold, r.destination)
}

type Rules []Rule

func (r Rules) Match(part Part) string {
	var destination string

	for _, rule := range r {
		var comparitor int

		switch rule.category {
		case "x":
			comparitor = part.x
		case "m":
			comparitor = part.m
		case "a":
			comparitor = part.a
		case "s":
			comparitor = part.s
		case "z":
			// This is duplicative but it doesn't matter; the z-case conditions are always true; this is just for clarity
			comparitor = 0
		}

		if rule.condition(comparitor) {
			destination = rule.destination
			break
		}
	}

	return destination
}
