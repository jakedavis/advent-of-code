// Code generated by "stringer -type=Card"; DO NOT EDIT.

package day7

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Joker-1]
	_ = x[Two-2]
	_ = x[Three-3]
	_ = x[Four-4]
	_ = x[Five-5]
	_ = x[Six-6]
	_ = x[Seven-7]
	_ = x[Eight-8]
	_ = x[Nine-9]
	_ = x[Ten-10]
	_ = x[Jack-11]
	_ = x[Queen-12]
	_ = x[King-13]
	_ = x[Ace-14]
}

const _Card_name = "JokerTwoThreeFourFiveSixSevenEightNineTenJackQueenKingAce"

var _Card_index = [...]uint8{0, 5, 8, 13, 17, 21, 24, 29, 34, 38, 41, 45, 50, 54, 57}

func (i Card) String() string {
	i -= 1
	if i < 0 || i >= Card(len(_Card_index)-1) {
		return "Card(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Card_name[_Card_index[i]:_Card_index[i+1]]
}
