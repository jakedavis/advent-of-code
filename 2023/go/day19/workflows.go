package day19

import (
	"strconv"
	"strings"
)

type Workflows map[string]Rules

func (w Workflows) Accept(part Part) bool {
	// Initial case
	destination := "in"
	ruleset := w[destination]

	for {
		destination = ruleset.Match(part)
		switch destination {
		case "A":
			return true
		case "R":
			return false
		default:
			// In this case, go to the next step
			ruleset = w[destination]
		}
	}
}

func parseWorkflow(line string) (string, []Rule, error) {
	var rules []Rule

	splits := strings.Split(line, "{")
	identifier := splits[0]

	// Kind of gross but we just want the juicy bits in the curly braces
	rulesRaw := strings.Split(strings.TrimSuffix(splits[1], "}"), ",")

	finalDestination := string(rulesRaw[len(rulesRaw)-1])
	finalRule := Rule{
		category:    "z", // Special code I'm inventing for formatting purposes
		condition:   func(_ int) bool { return true },
		destination: finalDestination,
	}

	for _, element := range rulesRaw[0 : len(rulesRaw)-1] {
		splits := strings.Split(element, ":")

		asRune := []rune(splits[0])
		destination := splits[1]

		category := string(asRune[0])
		operator := string(asRune[1])
		threshold, err := strconv.Atoi(string(asRune[2:]))
		if err != nil {
			return identifier, rules, err
		}

		if !strings.Contains("xmas", string(category)) {
			return identifier, rules, errInvalidCategory
		}

		var fn func(value int) bool
		switch operator {
		case ">":
			fn = func(value int) bool {
				return value > threshold
			}
		case "<":
			fn = func(value int) bool {
				return value < threshold
			}
		default:
			return identifier, rules, errInvalidOperator
		}

		rule := Rule{
			category:    category,
			operator:    operator,
			threshold:   threshold,
			condition:   fn,
			destination: destination,
		}

		rules = append(rules, rule)
	}

	// The last rule is a special case and so we'll handle it here.
	rules = append(rules, finalRule)

	return identifier, rules, nil
}
