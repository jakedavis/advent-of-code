package day19

import (
	"fmt"
	"strconv"
	"strings"
)

type Part struct {
	x int
	m int
	a int
	s int
}

type Parts []Part

func (p Part) String() string {
	return fmt.Sprintf("x=%4d m=%4d a=%4d s=%4d", p.x, p.m, p.a, p.s)
}

func (p Part) Value() int {
	return p.x + p.m + p.a + p.s
}

func parsePart(line string) (Part, error) {
	var part Part

	// We have to parse this manually because it isn't actually JSON
	line = strings.TrimSuffix(strings.TrimPrefix(line, "{"), "}")

	for _, element := range strings.Split(line, ",") {
		e := strings.Split(element, "=")

		conv, err := strconv.Atoi(e[1])
		if err != nil {
			return part, err
		}

		switch e[0] {
		case "x":
			part.x = conv
		case "m":
			part.m = conv
		case "a":
			part.a = conv
		case "s":
			part.s = conv
		default:
			return part, errInvalidCategory
		}
	}

	return part, nil
}
